package router

import (
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/WaynerEP/restaurant-app/server/plugin/email/api"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct{}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	EmailApi := api.ApiGroupApp.EmailApi.EmailTest
	SendEmail := api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)  // Send a test email
		emailRouter.POST("sendEmail", SendEmail) // Send an email
	}
}
