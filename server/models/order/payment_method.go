package order

import "github.com/WaynerEP/restaurant-app/server/models/common"

type PaymentMethod struct {
	common.ModelId
	Name        string `json:"name" gorm:"not null;size:50"` // Nombre del método de pago
	Description string `json:"description" gorm:"size:255"`  // Descripción del método de pago
}
