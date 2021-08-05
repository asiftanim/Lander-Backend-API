package controllers

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"lander/configs"
	"lander/database"
	"lander/models"
	"lander/services"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var broker models.Broker
	c.BindJSON(&broker)

	var brokerPass = broker.Password

	database.DB.Where("email = ?", broker.Email).First(&broker)

	match := CheckPasswordHash(brokerPass, broker.Password)

	if broker.Id == 0 {
		c.JSON(http.StatusNotFound, "User not found!!!")
	}

	if match {
		token, err := services.CreateToken(broker.Id, broker.Email)
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

func RegisterBroker(c *gin.Context) {
	var broker models.Broker
	c.BindJSON(&broker)

	hashPass, hashErr := HashPassword(broker.Password)

	if hashErr != nil {
		fmt.Println(hashErr.Error())
		c.JSON(http.StatusNotFound, "Unable to encrypt password")
	}

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

func ResetPassword(c *gin.Context) {
	var broker models.Broker
	c.BindJSON(&broker)

	c.JSON(http.StatusOK, "Development in progress")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SendMessage(c *gin.Context) {
	var chat models.Chat
	c.BindJSON(&chat)

	chat.Status = 1
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
	myuuid := uuid.NewV4().String()

	//Delete existing profile image
	var old_broker models.Broker
	old_broker_result := database.DB.First(&old_broker, broker.Id)
	if old_broker_result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, old_broker_result.Error)
	}

	path, err := ConvertAndSaveBase64toPng(myuuid, broker.ImagePath)
	//fmt.Println(path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	broker.ImagePath = path
	broker.ModifiedAt = time.Now()

	db_err := database.DB.Model(&broker).Where("id = ?", broker.Id).Updates(models.Broker{Name: broker.Name, ImagePath: broker.ImagePath, ModifiedAt: broker.ModifiedAt}).Error

	if db_err != nil {
		fmt.Println(db_err.Error())
		c.JSON(http.StatusNotFound, db_err.Error)
	}

	if old_broker.ImagePath != "" {
		delete_error := DeleteBrokerOldImage(old_broker.ImagePath)
		if delete_error != nil {
			fmt.Println(delete_error)
			c.JSON(http.StatusNotFound, delete_error.Error)
		} else {
			c.JSON(http.StatusOK, broker)
			fmt.Println("ok")
		}
	}

}

func UpdateDomainAskingPrice(c *gin.Context) {
	var domainQuery models.ProspectDomainQuery
	c.BindJSON(&domainQuery)

	domainQuery.Status = 4
	domainQuery.ModifiedAt = time.Now()

	err := database.DB.Model(&domainQuery).Where("id = ?", domainQuery.Id).Update("asking_price", domainQuery.AskingPrice).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, err.Error)
	} else {
		c.JSON(http.StatusOK, domainQuery)
	}
}

func ConvertAndSaveBase64toPng(name string, data string) (string, error) {

	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Connot load config", err)
	}

	env_image_path := config.IMAGE_PATH

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	image_path := env_image_path + name + ".jpg"

	f, err := os.OpenFile(image_path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fmt.Println("Image file", image_path, "created")

	return image_path, err
}

func GetBrokerProfileInfo(c *gin.Context) {
	var broker models.Broker
	c.BindJSON(&broker)
	fmt.Println(broker.Id)
	result := database.DB.First(&broker, broker.Id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error)
		return
	} else {
		c.JSON(http.StatusOK, broker)
	}
}

func DeleteBrokerOldImage(broker_old_image_path string) error {
	//delete old file

	err := os.Remove(broker_old_image_path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
}
