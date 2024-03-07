package system

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.CSystemApi
	{
		sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig) // Get system configuration
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) // Set system configuration
		sysRouter.POST("getServerInfo", systemApi.GetServerInfo)     // Get server information
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)       // Reload the system
	}
}
