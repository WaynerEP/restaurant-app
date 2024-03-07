package system

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddBaseMenu)           // Add menu
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority) // Add menu and role association
		menuRouter.POST("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)     // Delete menu
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     // Update menu
	}
	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)                   // Get menu tree
		menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList)           // Paginate and get the basic menu list
		menuRouterWithoutRecord.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   // Get user dynamic routes
		menuRouterWithoutRecord.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // Get specified role menu
		menuRouterWithoutRecord.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // Get menu by ID
	}
	return menuRouter
}
