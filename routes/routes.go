package routes

import (
	"lander/controllers"
	"lander/middlewares"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	server := gin.New()
	server.Static("/public", "./public")

	api := server.Group("api")
	{
		v1 := api.Group("v1")
		{
			//Broker
			v1.POST("broker/login", controllers.Login)
			v1.POST("broker/registration", controllers.RegisterBroker)

			broker := v1.Group("broker").Use(middlewares.Authorize())
			{
				broker.PUT("reset-password", controllers.ResetPassword)
				broker.GET("profile-info", controllers.GetBrokerProfileInfo)
				broker.PUT("profile-update", controllers.UpdateBroker)
				broker.POST("send-message", controllers.SendMessage)
				broker.GET("get-all-prospects", controllers.GetAllProspects)
				broker.GET("get-prospect-by-email", controllers.GetProspectByEmail)
				broker.GET("get-prospect-by-id", controllers.GetProspectById)
				broker.POST("set-price", controllers.GetAllProspects)
				broker.GET("get-all-domain-queries", controllers.GetProspectDomainQuery)
				broker.PUT("update-domain-asking-price", controllers.UpdateDomainAskingPrice)

			}

			//Prospect
			prospect := v1.Group("prospect")
			{
				prospect.POST("create-prospect", controllers.RegisterProspect)
				prospect.POST("verify-email", controllers.VerifyProspectEmail)
				prospect.POST("send-message", controllers.SendMessage)
				prospect.POST("create-domain-query", controllers.CreateProspectDomainQuery)
				prospect.POST("create-transaction", controllers.CreateNewTransaction)
			}

			//Common

		}
	}

	return server
}
