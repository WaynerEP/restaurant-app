package inventory

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/gin-gonic/gin"
)

type SupplyRouter struct{}

func (e *SupplyRouter) InitSupplyRouter(Router *gin.RouterGroup) {
	supplyRouter := Router.Group("supply").Use(middleware.OperationRecord())
	supplyRouterWithoutRecord := Router.Group("supply")
	supplyApi := v1.ApiGroupApp.InventoryApiGroup.SupplyApi
	{
		supplyRouter.POST("createSupply", supplyApi.CreateSupply)
		supplyRouter.PUT("updateSupply", supplyApi.UpdateSupply)
		supplyRouter.DELETE("deleteSupply", supplyApi.DeleteSupply)
	}
	{
		supplyRouterWithoutRecord.GET("getSupply", supplyApi.GetSupply)
		supplyRouterWithoutRecord.GET("getSupplyList", supplyApi.GetSupply)
	}
}
