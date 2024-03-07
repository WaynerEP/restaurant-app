package menu

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
)

type ItemCatRouter struct{}

func (e *ItemCatRouter) InitItemCatRouter(Router *gin.RouterGroup) {
	itemCatRouter := Router.Group("itemCategory").Use(middleware.OperationRecord())
	itemCatRouterWithoutRecord := Router.Group("itemCategory")
	itemCatApi := v1.ApiGroupApp.MenuApiGroup.ItemCategoryApi
	{
		itemCatRouter.POST("createItemCategory", itemCatApi.CreateItemCategory)
		itemCatRouter.PUT("updateItemCategory", itemCatApi.UpdateItemCategory)
		itemCatRouter.DELETE("deleteItemCategory", itemCatApi.DeleteItemCategory)
	}
	{
		itemCatRouterWithoutRecord.GET("getItemCategory", itemCatApi.GetItemCategory)
		itemCatRouterWithoutRecord.GET("getItemCategoryList", itemCatApi.GetItemCategory)
	}
}
