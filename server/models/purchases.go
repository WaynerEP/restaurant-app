package models

import (
	"time"

	"github.com/WaynerEP/restaurant-app/server/models/common"
)

type PurchaseOrder struct {
	common.ModelId
	ContactID             uint                   `json:"contactID" validate:"required"`
	CurrencyCode          string                 `json:"currencyCode" validate:"required" gorm:"size:10; not null;"`
	Currency              Currency               `gorm:"foreignKey:CurrencyCode;references:Code"`
	WarehouseID           uint                   `json:"warehouseID" validate:"required"`
	PaymentMethodCode     string                 `json:"paymentMethodCode" gorm:"size:10; not null;" validate:"required"`
	PaymentMethod         PaymentMethod          `gorm:"foreignKey:PaymentMethodCode;references:Code"`
	Subtotal              float64                `json:"subtotal" gorm:"decimal(9,2); not null; comment:'Es el total sin descuento ni igv.'"`
	TotalDiscount         float64                `json:"totalDiscount" gorm:"decimal(9,2) not null; comment:'Es el descuento que se aplica al subtotal'"`
	TotalTax              float64                `json:"totalTax" gorm:"decimal(9,2); not null; comment:'Es el igv aplicable al subtotal con descuento'"`
	TotalExonerated       float64                `json:"totalExonerated" gorm:"decimal(9,2);"`
	Total                 float64                `json:"total" gorm:"decimal(9,2); not null; comment:'Es el total con descuento e igv.'"`
	PurchaseOrderProducts []PurchaseOrderProduct `json:"purchaseOrderProducts"`
	TypeChange            float64                `json:"typeChange"`
	Credit                *bool                  `json:"credit" gorm:"default:0;"`
	Terms                 string                 `json:"terms"`
	Notes                 string                 `json:"notes"`
	DateDelivery          time.Time              `json:"dateDelivery"`
	DueDate               time.Time              `json:"dueDate"`
	Status                *bool                  `json:"status" gorm:"default:1;"` //0: no pagado, 1 pagado
	common.ControlBy
	common.ModelTime
}

func (u *PurchaseOrder) SetId(ID uint) {
	u.ID = ID
}

type PurchaseOrderProduct struct {
	common.ModelId
	PurchaseOrderID uint    `json:"purchaseOrderID"`
	ProductID       uint    `json:"productId"`
	Price           float64 `json:"price" gorm:"decimal(9,2); not null; comment:'Es el precio con el cual se va a vender.'" validate:"required"`
	Unit            string  `json:"unit" gorm:"not null; size:20;"`
	Quantity        uint64  `json:"quantity" gorm:"not null; CHECK:quantity>0;" validate:"required"`
	Discount        float64 `json:"discount" gorm:"decimal(9,2); not null; default:0;"`
	Subtotal        float64 `json:"subtotal" gorm:"decimal(9,2); not null"`
}
