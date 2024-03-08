package response

import "github.com/WaynerEP/restaurant-app/server/models/reservation"

type FloorResponse struct {
	Floor reservation.Floor `json:"floor"`
}
