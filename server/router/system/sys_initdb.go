package system

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // Initialize the database
		initRouter.POST("checkdb", dbApi.CheckDB) // Check if the database needs initialization
	}
}
