package menu

import "github.com/WaynerEP/restaurant-app/server/service"

type ApiGroup struct {
	ItemApi
	ItemCategoryApi
}

var (
	itemService         = service.ServiceGroupApp.MenuServiceGroup.ItemService
	itemCategoryService = service.ServiceGroupApp.MenuServiceGroup.ItemCategoryService
)
