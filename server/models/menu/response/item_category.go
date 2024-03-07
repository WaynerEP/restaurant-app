package response

import "github.com/WaynerEP/restaurant-app/server/models/menu"

type ItemCategoryResponse struct {
	ItemCategory menu.ItemCategory `json:"itemCategory"`
}
