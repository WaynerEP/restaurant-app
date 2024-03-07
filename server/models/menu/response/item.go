package response

import "github.com/WaynerEP/restaurant-app/server/models/menu"

type ItemResponse struct {
	Item menu.Item `json:"item"`
}
