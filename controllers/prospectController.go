package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"lander/database"
	"lander/models"
	"fmt"
	"crypto/rand"
	"encoding/hex"
)

func RegisterProspect(c *gin.Context){
	var prospect models.Prospect
	c.BindJSON(&prospect)

	//check if cookie exist. If not then check if mail exists. if not then register user else retrive and set cookie

	prospect.CookiSecrect = GenerateCookieSecret()
	prospect.IsActive = true
	prospect.IsDelete = false
	prospect.IsEmailVerified = false
	prospect.CreatedDate = time.Now()
	prospect.ModifiedDate = time.Now()

	err := database.DB.Create(&prospect).Error;

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "ProspectCookiSecrect", Value: prospect.CookiSecrect, Expires: expiration}
	http.SetCookie(c.Writer, &cookie)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, "Email Already Exists")
	} else {

		c.JSON(http.StatusOK, prospect)
	}
}



func VerifyProspectEmail(c *gin.Context) {
	var prospect models.Prospect
	c.BindJSON(&prospect)

	err := database.DB.Where("cooki_secrect = ?", prospect.CookiSecrect).First(&prospect).Error

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}else{
		if(prospect.IsEmailVerified){
			c.JSON(http.StatusOK, "Email Already Verified")
		}else{
			database.DB.Model(&prospect).Update("is_email_verified", true)
			c.JSON(http.StatusOK, "Prospect Email Verified")
		}
		
	}
}

func GenerateCookieSecret() string {
	key := make([]byte, 64)
	_ , err := rand.Read(key)
	if err != nil {
		// handle error here
	}
	
	return hex.EncodeToString(key)
}

func GetProspectByEmail(c *gin.Context){
	var prospect models.Prospect
	c.BindJSON(&prospect)

	err := database.DB.Where("Email = ?", prospect.Email).First(&prospect).Error

	if err != nil{
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}else{
		c.JSON(http.StatusOK, prospect)
	}
}

func GetProspectById(c *gin.Context){
	var prospect models.Prospect
	c.BindJSON(&prospect)

	err := database.DB.Where("Id = ?", prospect.Id).First(&prospect).Error

	if err != nil{
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}else{
		c.JSON(http.StatusOK, prospect)
	}
}

func CreateProspectDomainQuery(c *gin.Context){
	var domainQuery models.ProspectDomainQuery
	c.BindJSON(&domainQuery)

	domainQuery.Status = 1 //will change to Enum later
	domainQuery.CreatedDate = time.Now()
	domainQuery.ModifiedDate = time.Now()

	fmt.Println(domainQuery)

	err := database.DB.Create(&domainQuery).Error;

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, err.Error)
	} else {
		c.JSON(http.StatusOK, domainQuery)
	}
}

func GetProspectDomainQuery(c *gin.Context){
	var domainQueries []models.ProspectDomainQuery
	c.BindJSON(&domainQueries)

	result := database.DB.Find(&domainQueries)

	if result.Error != nil{
		c.JSON(http.StatusNotFound, result.Error)
			return
	}else{
		c.JSON(http.StatusOK, gin.H{
			"data": domainQueries,
		})
	}
}

func CreateNewTransaction(c *gin.Context){
	var transaction models.Transaction
	c.BindJSON(&transaction)

	
	transaction.CreatedDate = time.Now()
	transaction.ModifiedDate = time.Now()
	transaction.IsActive = true
	transaction.IsDelete = false
	
	fmt.Println(transaction)

	err := database.DB.Create(&transaction).Error;

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, err.Error)
	} else {
		c.JSON(http.StatusOK, transaction)
	}

}