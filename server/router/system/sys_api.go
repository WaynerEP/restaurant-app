package system

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")

	apiPublicRouterWithoutRecord := RouterPub.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // Create API
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // Delete API
		apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // Get a single API by ID
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // Update API
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // Delete selected APIs
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // Get all APIs
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // Get API list
	}
	{
		apiPublicRouterWithoutRecord.GET("freshCasbin", apiRouterApi.FreshCasbin) // Refresh Casbin permissions
	}
}
