package response

import "github.com/WaynerEP/restaurant-app/server/models/reservation"

type TableResponse struct {
	Table reservation.Table `json:"table"`
}
