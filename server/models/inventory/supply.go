package inventory

import "github.com/WaynerEP/restaurant-app/server/models/common"

type Supply struct {
	common.ModelId
	Name             string         `json:"name" gorm:"size:50;not null"`                         // Name of the supply
	Description      string         `json:"description" gorm:"type:text"`                         // Description of the supply
	SupplyCategoryID uint           `json:"categoryId" gorm:"not null"`                           // ID of the supply category
	SupplyCategory   SupplyCategory `json:"supplyCategory" validate:"-"`                          // Relationship with the supply category
	UnitMeasureID    uint           `json:"unitMeasureId" gorm:"not null"`                        // Unit of measure for the supply
	StockQuantity    float64        `json:"stockQuantity" gorm:"not null;check:stock_quantity>0"` // Stock quantity of the supply
	MinStockAlert    float64        `json:"minStockAlert" gorm:"not null;default:10"`             // Minimum stock alert level for the supply
	ImageURL         string         `json:"imageURL"`                                             // URL of the image for the supply
	IsActive         *bool          `json:"isActive" gorm:"not null;default:1"`                   // Indicates if the supply is active
	common.ModelTime
}
