package response

import "github.com/WaynerEP/restaurant-app/server/models/inventory"

type SupplyCategoryResponse struct {
	SupplyCategory inventory.SupplyCategory `json:"supplyCategory"`
}
