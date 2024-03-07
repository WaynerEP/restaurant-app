package inventory

import "github.com/WaynerEP/restaurant-app/server/service"

type ApiGroup struct {
	UnitOfMeasureApi
	SupplyCategoryApi
	SupplyApi
}

var (
	unitOfMeasureService  = service.ServiceGroupApp.InventoryServiceGroup.UnitOfMeasureService
	supplyCategoryService = service.ServiceGroupApp.InventoryServiceGroup.SupplyCategoryService
	supplyService         = service.ServiceGroupApp.InventoryServiceGroup.SupplyService
)
