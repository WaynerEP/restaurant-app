package inventory

import "github.com/WaynerEP/restaurant-app/server/models/common"

type SupplyCategory struct {
	common.ModelId
	Name        string `json:"name" gorm:"size:50;not null"` // Name of the supply category
	Description string `json:"description" gorm:"size:255"`
}
