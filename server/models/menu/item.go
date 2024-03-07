package menu

import "github.com/WaynerEP/restaurant-app/server/models/common"

// Item .
type Item struct {
	common.ModelId
	InternalCode     string           `json:"internalCode" gorm:"size:100;not null"`
	Name             string           `json:"name" gorm:"size:80;not null"`   // Nombre del plato o item
	Description      string           `json:"description" gorm:"size:255"`    // Descripción del plato
	Price            float64          `json:"price" gorm:"not null"`          // Precio del plato
	ItemCategoryID   string           `json:"itemCategoryId" gorm:"not null"` // Categoría del plato (por ejemplo, entrante, principal, postre, etc.)
	ItemCategory     ItemCategory     `json:"itemCategory"`
	NutritionalValue NutritionalValue `json:"nutritionalValue"`
	ItemSupplies     []ItemSupply     `json:"itemSupplies"`
	ImageURL         string           `json:"imageURL"` // URL de la imagen del plato

	IsActive *bool `json:"isActive" gorm:"not null;default:1"` // Indica si el plato está activo o no
	common.ModelTime
}

type ItemSupply struct {
	ItemID   uint    `json:"itemId" gorm:"primaryKey"`   // ID del plato asociado
	SupplyID uint    `json:"supplyId" gorm:"primaryKey"` // ID del ingrediente asociado
	Quantity float64 `json:"quantity" gorm:"not null"`   // Cantidad del ingrediente necesaria para el plato
}
