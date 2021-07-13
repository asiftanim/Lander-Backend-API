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
)

type User struct {
	ID uint64            `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}

var user = User{
	ID:            1,
	Username: "username",
	Password: "password",
	Phone: "49123454322", //this is a random number
}
  
  
func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
	   c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
	   return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
	   c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	   return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
	   c.JSON(http.StatusUnprocessableEntity, err.Error())
	   return
	}
	c.JSON(http.StatusOK, token)
}

func CreateToken(userId uint64) (string, error) {
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

	err := CreateUser(&broker)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, broker)
	}
}

func CreateUser(broker *models.Broker) (err error) {
	if err = database.DB.Create(broker).Error; err != nil {
		return err
	}
	return nil
}

func InsertChat(c *gin.Context){
	var chat models.Chat
	c.BindJSON(&chat)

	err := CreateChat(&chat)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, chat)
	}
}

func CreateChat(chat *models.Chat) (err error) {
	if err = database.DB.Create(chat).Error; err != nil {
		return err
	}
	return nil
}

