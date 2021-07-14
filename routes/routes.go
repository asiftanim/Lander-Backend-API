package routes

import (
	"lander/controllers"
	
	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine{
	server := gin.New()

	//Code refactor needed
	api_version := server.Group("/v1/api")
	{
		//Broker
		api_version.POST("login", controllers.Login)
		api_version.POST("registration", controllers.RegisterBroker)

		//Prospect
		api_version.POST("create-prospect", controllers.RegisterProspect)
		api_version.POST("verify-email", controllers.VerifyProspectEmail)

		//Common
		api_version.POST("create-chat", controllers.InsertChat)

	}

	return server;
}