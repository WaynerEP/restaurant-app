package v1

import (
	"github.com/WaynerEP/restaurant-app/server/api/v1/contact"
	"github.com/WaynerEP/restaurant-app/server/api/v1/inventory"
	"github.com/WaynerEP/restaurant-app/server/api/v1/menu"
	"github.com/WaynerEP/restaurant-app/server/api/v1/order"
	"github.com/WaynerEP/restaurant-app/server/api/v1/reservation"
	"github.com/WaynerEP/restaurant-app/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ContactApiGroup     contact.ApiGroup
	InventoryApiGroup   inventory.ApiGroup
	MenuApiGroup        menu.ApiGroup
	ReservationApiGroup reservation.ApiGroup
	OrderApiGroup       order.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
