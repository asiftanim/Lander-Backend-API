package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"lander/database"
	"lander/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var broker models.Broker
	c.BindJSON(&broker)

	var brokerPass = broker.Password

	database.DB.Where("Email = ?", broker.Email).First(&broker)

	match := CheckPasswordHash(brokerPass, broker.Password)

	if broker.ID == 0 {
		c.JSON(http.StatusNotFound, "User not found!!!")
	}

	if match {
		token, err := CreateToken(broker.ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"_token": token,
		})
	} else {
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

func RegisterBroker(c *gin.Context) {
	var broker models.Broker
	c.BindJSON(&broker)

	hashPass, _ := HashPassword(broker.Password)
	broker.Password = hashPass
	broker.IsActive = true
	broker.IsDelete = false
	broker.CreatedAt = time.Now()
	broker.ModifiedAt = time.Now()

	err := database.DB.Create(&broker).Error

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

func InsertChat(c *gin.Context) {
	var chat models.Chat
	c.BindJSON(&chat)

	chat.CreatedAt = time.Now()

	err := database.DB.Create(&chat).Error

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, chat)
	}
}

func GetAllProspects(c *gin.Context) {
	var prospects []models.Prospect

	result := database.DB.Find(&prospects)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": prospects,
		})
	}
}

func UpdateBroker(c *gin.Context) {
	var broker models.Broker
	c.BindJSON(&broker)

	myuuid := uuid.NewV4()
	fmt.Println(myuuid)

	c.JSON(http.StatusOK, myuuid)
	GenerateBase64ToImage(broker.ImagePath)
	//save file to local folder and set name as UUID

	// folder_path_from_env := "../public/broker/"
	//broker.ImagePath = folder_path_from_env+myuuid+".jpg"

	//err := database.DB.Model(&broker).Where("id = ?", broker.ID).Update("image_path", broker.ImagePath).Error

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	c.JSON(http.StatusNotFound, err.Error)
	// } else {
	// 	c.JSON(http.StatusOK, broker)
	// }

}

func UpdateDomainAskingPrice(c *gin.Context) {
	var domainQuery models.ProspectDomainQuery
	c.BindJSON(&domainQuery)

	domainQuery.Status = 4
	domainQuery.ModifiedAt = time.Now()

	fmt.Println(domainQuery)

	err := database.DB.Model(&domainQuery).Where("id = ?", domainQuery.ID).Update("asking_price", domainQuery.AskingPrice).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, err.Error)
	} else {
		c.JSON(http.StatusOK, domainQuery)
	}

}

func GenerateBase64ToImage(data string) {

	r := bytes.NewReader([]byte(data))

	reader := base64.NewDecoder(base64.StdEncoding, r)
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	//Encode from image format to writer
	pngFilename := "test.png"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Png file", pngFilename, "created")

}
