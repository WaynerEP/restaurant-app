package contact

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"github.com/WaynerEP/restaurant-app/server/models/order"
)

type Customer struct {
	common.ModelId
	DocumentType   string               `json:"documentType" gorm:"size:25; not null" validate:"required"`         // Tipo de documento de identidad del cliente (por ejemplo, DNI, pasaporte, etc.)
	DocumentNumber string               `json:"documentNumber" gorm:"size:20;not null;unique" validate:"required"` // Número de documento de identidad del cliente
	Name           string               `json:"name" gorm:"size:150;not null" validate:"required"`                 // Nombre del cliente
	Lastname       string               `json:"lastname" gorm:"size:150;not null" validate:"required"`
	Email          string               `json:"email" gorm:"size:255"` // Correo electrónico del cliente
	Phone          string               `json:"phone" gorm:"size:15"`  // Número de teléfono del cliente
	Address        string               `json:"address"`
	MenuOrders     []order.MenuOrder    `json:"menuOrders"`    // Relación uno a muchos con Order
	PaymentOrders  []order.PaymentOrder `json:"paymentOrders"` // Relación uno a muchos con PaymentOrder
	common.ModelTime
	common.ControlBy
}

// TableName returns the table name for SysUser.
func (Customer) TableName() string {
	return "customers"
}
