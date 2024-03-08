package reservation

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type TableRouter struct{}

func (e *TableRouter) InitTableRouter(Router *gin.RouterGroup) {
	tableRouter := Router.Group("table").Use(middleware.OperationRecord())
	tableRouterWithoutRecord := Router.Group("table")
	tableApi := v1.ApiGroupApp.ReservationApiGroup.TableApi
	{
		tableRouter.POST("createTable", tableApi.CreateTable)
		tableRouter.POST("createFloorEnvTable", tableApi.CreateFloorEnvTable)
		tableRouter.PUT("updateTable", tableApi.UpdateTable)
		tableRouter.DELETE("deleteTable", tableApi.DeleteTable)
		tableRouter.DELETE("deleteFloorEnvironmentTable", tableApi.DeleteFloorEnvironmentTable)
	}
	{
		tableRouterWithoutRecord.GET("getTable", tableApi.GetTable)
		tableRouterWithoutRecord.GET("getTableList", tableApi.GetTableList)
		tableRouterWithoutRecord.GET("getOptionsForSelect", tableApi.GetOptionsForSelect)
		tableRouterWithoutRecord.GET("getTablesByFloorEnvironmentId/:id", tableApi.GetTablesByFloorEnvironmentId)

	}
}
