package reservation

import "github.com/WaynerEP/restaurant-app/server/service"

type ApiGroup struct {
	FloorApi
	EnvApi
	TableApi
}

var (
	floorService = service.ServiceGroupApp.ReservationGroup.FloorService
	envService   = service.ServiceGroupApp.ReservationGroup.EnvService
	tableService = service.ServiceGroupApp.ReservationGroup.TableService
)
