package routes

import (
	"lander/controllers/broker"
	"github.com/gin-gonic/gin"
)

// var {

// }

func Init() *gin.Engine{
	server := gin.New()

	api_version := server.Group("/v1/api")
	{
		api_version.POST("login", controllers.Login)
		api_version.POST("registration", controllers.RegisterBroker)
		api_version.POST("create_chat", controllers.InsertChat)
	}

	return server;
}