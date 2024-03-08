package order

import "github.com/WaynerEP/restaurant-app/server/service"

type ApiGroup struct {
	MenuOrderApi
}

var (
	menuOrderService = service.ServiceGroupApp.OrderGroup.MenuOrderService
)
