package menu

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"gorm.io/datatypes"
)

// Recipe Estructura del modelo Recipe
type Recipe struct {
	common.ModelId
	Name        string         `json:"name" gorm:"size:50;not null"`          // Nombre de la receta
	Description string         `json:"description" gorm:"type:text;not null"` // Descripción de la receta
	ImageURLs   datatypes.JSON `json:"imageURLs" gorm:""`                     // URLs de las imágenes asociadas a la receta
	ItemID      uint           `json:"itemID" `
	IsActive    *bool          `json:"isActive" gorm:"not null;default:1"` // Indica si la receta está activa o no
	common.ModelTime
}
