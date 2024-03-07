package order

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"time"
)

type PaymentOrder struct {
	common.ModelId
	MenuOrderID      uint      `json:"menuOrderID" gorm:"not null"` // Order ID
	CustomerID       uint      `json:"customerId" gorm:"not null"`
	Amount           float64   `json:"amount" gorm:"not null;check:amount > 0"`                    // Amount of the payment
	ReceivedAmount   float64   `json:"receivedAmount" gorm:"not null;check:amount > 0"`            // Amount received
	ChangeAmount     float64   `json:"changeAmount" gorm:"not null;default:0"`                     // Change amount
	PendingAmount    float64   `json:"pendingAmount" gorm:"not null;default:0"`                    // Total de pago pendiente del cliente
	TotalOrderAmount float64   `json:"totalOrderAmount" gorm:"not null"`                           //Total de la orden
	PaymentMethodID  uint      `json:"paymentMethod"  gorm:"not null"`                             // Payment method used
	TransactionID    string    `json:"transactionId"  gorm:"size:25"`                              // Transaction ID
	PaymentStatus    string    `json:"paymentStatus"  gorm:"size:20;not null;default:'Pendiente'"` // Payment status
	PaymentDate      time.Time `json:"paymentDate"  gorm:"not null"`                               // Date and time of the payment
	common.ModelTime
	common.ControlBy
}

/*Pendiente: El pago se ha registrado pero aún no se ha procesado.
Aprobado: El pago ha sido aprobado y se ha completado correctamente.
Rechazado: El pago ha sido rechazado debido a un problema con la transacción.
Cancelado: El pago ha sido cancelado antes de su procesamiento.
En proceso: El pago está siendo procesado y aún no se ha completado.
Expirado: El pago ha caducado antes de ser completado.
Fallido: El pago ha fallado debido a un problema técnico u otro motivo.
Reembolsado: El pago ha sido reembolsado al cliente.*/
