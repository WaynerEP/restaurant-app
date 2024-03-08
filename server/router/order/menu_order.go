package order

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type MenuOrderRouter struct{}

func (e *MenuOrderRouter) InitMenuOrderRouter(Router *gin.RouterGroup) {
	menuOrderRouter := Router.Group("menuOrder").Use(middleware.OperationRecord())
	menuOrderRouterWithoutRecord := Router.Group("menuOrder")
	menuOrderApi := v1.ApiGroupApp.OrderApiGroup.MenuOrderApi
	{
		menuOrderRouter.PUT("readyMenuOrder", menuOrderApi.ReadyMenuOrder)
		menuOrderRouter.PUT("approveMenuOrder", menuOrderApi.ApproveMenuOrder)
		menuOrderRouter.PUT("rejectMenuOrder", menuOrderApi.RejectMenuOrder)
		menuOrderRouter.PUT("updateStatusMenuOrder", menuOrderApi.UpdateStatusMenuOrder)
		menuOrderRouter.POST("createMenuOrder", menuOrderApi.CreateMenuOrder)
		menuOrderRouter.PUT("updateMenuOrder", menuOrderApi.UpdateMenuOrder)
		menuOrderRouter.DELETE("deleteMenuOrder", menuOrderApi.DeleteMenuOrder)
	}
	{
		menuOrderRouterWithoutRecord.GET("getMenuOrder", menuOrderApi.GetMenuOrder)
		menuOrderRouterWithoutRecord.GET("getMenuOrderList", menuOrderApi.GetMenuOrder)
	}
}
