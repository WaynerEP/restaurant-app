package menu

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type ItemRouter struct{}

func (e *ItemRouter) InitItemRouter(Router *gin.RouterGroup) {
	itemRouter := Router.Group("item").Use(middleware.OperationRecord())
	itemRouterWithoutRecord := Router.Group("item")
	itemApi := v1.ApiGroupApp.MenuApiGroup.ItemApi
	{
		itemRouter.POST("createItem", itemApi.CreateItem)
		itemRouter.PUT("updateItem", itemApi.UpdateItem)
		itemRouter.DELETE("deleteItem", itemApi.DeleteItem)
	}
	{
		itemRouterWithoutRecord.GET("getItem", itemApi.GetItem)
		itemRouterWithoutRecord.GET("getItemList", itemApi.GetItem)
	}
}
