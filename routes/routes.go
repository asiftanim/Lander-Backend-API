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
				broker.GET("get-all-prospects", controllers.GetAllProspects)
				broker.GET("get-prospect-by-email", controllers.GetProspectByEmail)
				broker.GET("get-prospect-by-id", controllers.GetProspectById)
				broker.POST("set-price", controllers.GetAllProspects)
				broker.GET("get-all-domain-queries", controllers.GetProspectDomainQuery)
				
			}
			
			//Prospect
			prospect := v1.Group("prospect")
			{
				prospect.POST("create-prospect", controllers.RegisterProspect)
				prospect.POST("verify-email", controllers.VerifyProspectEmail)
				prospect.POST("create-message", controllers.InsertChat)
				prospect.POST("create-domain-query", controllers.CreateProspectDomainQuery)
			}
	
			//Common
			
	
		}
	}

	

	return server;
}