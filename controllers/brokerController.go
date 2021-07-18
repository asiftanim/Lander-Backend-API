package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	"lander/database"
	"lander/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
  
  
func Login(c *gin.Context) {
	var broker models.Broker
	c.BindJSON(&broker)
	

	var brokerPass = broker.Password

	database.DB.Where("Email = ?", broker.Email).First(&broker)

	match := CheckPasswordHash(brokerPass, broker.Password)

	if broker.Id == 0{
		c.JSON(http.StatusNotFound, "User not found!!!")
	}

	if(match){
		token, err := CreateToken(broker.Id)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"_token": token,
		})
	}else{
		c.JSON(http.StatusUnauthorized, "Please provide valid login credentials")
	}

}

func CreateToken(userId uint) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func RegisterBroker(c *gin.Context){
	var broker models.Broker
	c.BindJSON(&broker)

	hashPass, _ := HashPassword(broker.Password)
	broker.Password = hashPass
	broker.IsActive = true
	broker.IsDelete = false
	broker.CreatedDate = time.Now()
	broker.ModifiedDate = time.Now()

	err := database.DB.Create(&broker).Error;

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, "Email Already Exists")
	} else {
		c.JSON(http.StatusOK, broker)
	}
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func InsertChat(c *gin.Context){
	var chat models.Chat
	c.BindJSON(&chat)

	chat.CreatedDate = time.Now()

	err := database.DB.Create(&chat).Error;

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, chat)
	}
}

func GetAllProspects(c *gin.Context){
	var prospects []models.Prospect

	result := database.DB.Find(&prospects)

	if result.Error != nil{
		c.JSON(http.StatusNotFound, result.Error)
			return
	}else{
		c.JSON(http.StatusOK, gin.H{
			"data": prospects,
		})
	}
}

func UpdateBroker(c *gin.Context){

}



