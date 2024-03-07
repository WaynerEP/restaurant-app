package initialize

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/WaynerEP/restaurant-app/server/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	// Set to release mode
	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) // DebugMode, ReleaseMode, TestMode
	}

	Router := gin.New()
	Router.Use(gin.Recovery())
	if global.GVA_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger())
	}

	Router.Use(middleware.Cors())

	InstallPlugin(Router) // Install plugins
	systemRouter := router.AppRouterGroup.System
	contactRouter := router.AppRouterGroup.Contact
	inventoryRouter := router.AppRouterGroup.Inventory
	menuRouter := router.AppRouterGroup.Menu
	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// Health check
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // Register basic function routes without authentication
		systemRouter.InitInitRouter(PublicGroup) // Automatically initialize related routes
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		//System
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)   // Register API routes
		systemRouter.InitJwtRouter(PrivateGroup)                // JWT-related routes
		systemRouter.InitUserRouter(PrivateGroup)               // Register user routes
		systemRouter.InitMenuRouter(PrivateGroup)               // Register menu routes
		systemRouter.InitSystemRouter(PrivateGroup)             // System-related routes
		systemRouter.InitCasbinRouter(PrivateGroup)             // Authorization-related routes
		systemRouter.InitAuthorityRouter(PrivateGroup)          // Register role routes
		systemRouter.InitSysOperationRecordRouter(PrivateGroup) // Operation record
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup) // Dictionary detail management
		//Contacts(customers, providers, etc)
		contactRouter.InitCustomerRouter(PrivateGroup)
		//Inventory
		inventoryRouter.InitUnitOfMeasureRouter(PrivateGroup)
		inventoryRouter.InitSupplyCategoryRouter(PrivateGroup)
		inventoryRouter.InitSupplyRouter(PrivateGroup)
		//Menu
		menuRouter.InitItemCatRouter(PrivateGroup)
		menuRouter.InitItemRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
