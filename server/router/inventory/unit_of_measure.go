package inventory

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type UnitOfMeasureRouter struct{}

func (e *UnitOfMeasureRouter) InitUnitOfMeasureRouter(Router *gin.RouterGroup) {
	measureRouter := Router.Group("unitOfMeasure").Use(middleware.OperationRecord())
	measureRouterWithoutRecord := Router.Group("unitOfMeasure")
	unitOfMeasureApi := v1.ApiGroupApp.InventoryApiGroup.UnitOfMeasureApi
	{
		measureRouter.POST("createUnitOfMeasure", unitOfMeasureApi.CreateUnitOfMeasure)
		measureRouter.PUT("updateUnitOfMeasure", unitOfMeasureApi.UpdateUnitOfMeasure)
		measureRouter.DELETE("deleteUnitOfMeasure", unitOfMeasureApi.DeleteUnitOfMeasure)
	}
	{
		measureRouterWithoutRecord.GET("getUnitOfMeasure", unitOfMeasureApi.GetUnitOfMeasure)
		measureRouterWithoutRecord.GET("getUnitOfMeasureList", unitOfMeasureApi.GetUnitOfMeasure)
	}
}
