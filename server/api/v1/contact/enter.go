package contact

import "github.com/WaynerEP/restaurant-app/server/service"

type ApiGroup struct {
	CustomerApi
}

var (
	customerService = service.ServiceGroupApp.CustomerServiceGroup.CustomerService
)
