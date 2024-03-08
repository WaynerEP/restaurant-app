package response

import "github.com/WaynerEP/restaurant-app/server/models/reservation"

type EnvironmentResponse struct {
	Environment reservation.Environment `json:"environment"`
}
