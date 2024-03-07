package response

import "github.com/WaynerEP/restaurant-app/server/models/inventory"

type UnitOfMeasureResponse struct {
	UnitOfMeasure inventory.UnitMeasure `json:"UnitOfMeasure"`
}
