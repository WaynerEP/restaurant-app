package reservation

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type FloorRouter struct{}

func (e *FloorRouter) InitFloorRouter(Router *gin.RouterGroup) {
	floorRouter := Router.Group("floor").Use(middleware.OperationRecord())
	floorRouterWithoutRecord := Router.Group("floor")
	floorApi := v1.ApiGroupApp.ReservationApiGroup.FloorApi
	{
		floorRouter.POST("createFloor", floorApi.CreateFloor)
		floorRouter.PUT("updateFloor", floorApi.UpdateFloor)
		floorRouter.DELETE("deleteFloor", floorApi.DeleteFloor)
	}
	{
		floorRouterWithoutRecord.GET("getFloor", floorApi.GetFloor)
		floorRouterWithoutRecord.GET("getFloorList", floorApi.GetFloorList)
		floorRouterWithoutRecord.GET("getTreeFloor", floorApi.GetTreeFloor)
		floorRouterWithoutRecord.GET("getOptionsForSelect", floorApi.GetOptionsForSelect)
	}
}
