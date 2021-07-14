package routes

import (
	"lander/controllers"
	
	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine{
	server := gin.New()

	api := server.Group("api")
	{
		v1 := api.Group("v1")
		{
			//Broker
			broker := v1.Group("broker")
			{
				broker.POST("login", controllers.Login)
				broker.POST("registration", controllers.RegisterBroker)
				broker.PUT("profile-update", controllers.UpdateBroker)
				broker.POST("create-message", controllers.InsertChat)
			}
			
			//Prospect
			prospect := v1.Group("prospect")
			{
				prospect.POST("create-prospect", controllers.RegisterProspect)
				prospect.POST("verify-email", controllers.VerifyProspectEmail)
				prospect.POST("create-message", controllers.InsertChat)
			}
	
			//Common
			
	
		}
	}

	

	return server;
}