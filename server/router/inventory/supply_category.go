package inventory

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type SupplyCategoryRouter struct{}

func (e *SupplyCategoryRouter) InitSupplyCategoryRouter(Router *gin.RouterGroup) {
	supplyCategoryRouter := Router.Group("supplyCategory").Use(middleware.OperationRecord())
	supplyCategoryRouterWithoutRecord := Router.Group("supplyCategory")
	supplyCategoryApi := v1.ApiGroupApp.InventoryApiGroup.SupplyCategoryApi
	{
		supplyCategoryRouter.POST("createSupplyCategory", supplyCategoryApi.CreateSupplyCategory)
		supplyCategoryRouter.PUT("updateSupplyCategory", supplyCategoryApi.UpdateSupplyCategory)
		supplyCategoryRouter.DELETE("deleteSupplyCategory", supplyCategoryApi.DeleteSupplyCategory)
	}
	{
		supplyCategoryRouterWithoutRecord.GET("getSupplyCategory", supplyCategoryApi.GetSupplyCategory)
		supplyCategoryRouterWithoutRecord.GET("getSupplyCategoryList", supplyCategoryApi.GetSupplyCategory)
	}
}
