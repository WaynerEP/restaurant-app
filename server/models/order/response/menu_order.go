package response

import "github.com/WaynerEP/restaurant-app/server/models/order"

type MenuOrderResponse struct {
	MenuOrder order.MenuOrder `json:"menuOrder"`
}
