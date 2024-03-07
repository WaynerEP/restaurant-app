package contact

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	customerApi := v1.ApiGroupApp.ContactApiGroup.CustomerApi
	{
		customerRouter.POST("createCustomer", customerApi.CreateCustomer)   // Create customer
		customerRouter.PUT("updateCustomer", customerApi.UpdateCustomer)    // Update customer
		customerRouter.DELETE("deleteCustomer", customerApi.DeleteCustomer) // Delete customer
	}
	{
		customerRouterWithoutRecord.GET("getCustomer", customerApi.GetCustomer)      // Get individual customer information
		customerRouterWithoutRecord.GET("customerList", customerApi.GetCustomerList) // Get customer list
	}
}
