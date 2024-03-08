package service

import (
	"github.com/WaynerEP/restaurant-app/server/service/contact"
	"github.com/WaynerEP/restaurant-app/server/service/inventory"
	"github.com/WaynerEP/restaurant-app/server/service/menu"
	"github.com/WaynerEP/restaurant-app/server/service/order"
	"github.com/WaynerEP/restaurant-app/server/service/reservation"
	"github.com/WaynerEP/restaurant-app/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	CustomerServiceGroup  contact.ServiceGroup
	InventoryServiceGroup inventory.ServiceGroup
	MenuServiceGroup      menu.ServiceGroup
	ReservationGroup      reservation.ServiceGroup
	OrderGroup            order.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
