package reservation

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type EnvRouter struct{}

func (e *EnvRouter) InitEnvRouter(Router *gin.RouterGroup) {
	envRouter := Router.Group("environment").Use(middleware.OperationRecord())
	envRouterWithoutRecord := Router.Group("environment")
	tableApi := v1.ApiGroupApp.ReservationApiGroup.EnvApi
	{
		envRouter.POST("createEnvironment", tableApi.CreateEnvironment)
		envRouter.POST("createFloorEnvironment", tableApi.CreateFloorEnvironment)
		envRouter.PUT("updateEnvironment", tableApi.UpdateEnvironment)
		envRouter.DELETE("deleteEnvironment", tableApi.DeleteEnvironment)
		envRouter.DELETE("deleteFloorEnvironment", tableApi.DeleteFloorEnvironment)
	}
	{
		envRouterWithoutRecord.GET("getEnvironment", tableApi.GetEnvironment)
		envRouterWithoutRecord.GET("getEnvironmentList", tableApi.GetEnvironmentList)
		envRouterWithoutRecord.GET("getOptionsForSelect", tableApi.GetOptionsForSelect)
		envRouterWithoutRecord.GET("getEnvironmentsByFloorId/:id", tableApi.GetEnvironmentsByFloorId)
	}
}
