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

	prospect.CookiSecrect = GenerateCookieSecret()
	prospect.IsActive = true
	prospect.IsDelete = false
	prospect.IsEmailVerified = false
	prospect.CreatedDate = time.Now()
	prospect.ModifiedDate = time.Now()

	err := database.DB.Create(&prospect).Error;

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