package response

import (
	"github.com/WaynerEP/restaurant-app/server/models/inventory"
)

type SupplyResponse struct {
	Supply inventory.Supply `json:"supply"`
}
