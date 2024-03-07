package inventory

import "github.com/WaynerEP/restaurant-app/server/models/common"

type UnitMeasure struct {
	common.ModelId
	Code         string `json:"code" gorm:"size:20;not null"`
	Name         string `json:"name" gorm:"size:50;not null"` // Name of the unit of measure
	Abbreviation string `json:"abbreviation" gorm:"size:10;not null"`
}
