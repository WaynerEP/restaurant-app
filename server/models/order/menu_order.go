package order

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"github.com/WaynerEP/restaurant-app/server/models/reservation"
	"time"
)

type MenuOrder struct {
	common.ModelId
	OrderNumber            string                              `json:"orderNumber"  gorm:"size:20;not null"` // Order identification number
	CustomerID             uint                                `json:"customerId" gorm:"not null"`
	Subtotal               float64                             `json:"subtotal" gorm:"not null"`                                  // Order total
	Total                  float64                             `json:"total" gorm:"not null;check:total > 0"`                     // Order total
	Notes                  string                              `json:"notes" gorm:"type:text"`                                    // Additional notes for the order
	Discount               float64                             `json:"discount"`                                                  // Discounts applied to the order total
	Taxes                  float64                             `json:"taxes"`                                                     // Taxes applied to the order total
	DeliveryMethod         string                              `json:"deliveryMethod"  gorm:"size:50"`                            // Delivery method selected by the contact
	EstimatedDeliveryTime  time.Time                           `json:"estimatedDeliveryTime"`                                     // Estimated delivery time of the order
	Status                 string                              `json:"status"  gorm:"size:20;not null;default:'Pendiente'"`       // Order status
	PaymentStatus          string                              `json:"paymentStatus" gorm:"size:20;not null;default:'Pendiente'"` // Payment status of the order
	OrderDate              time.Time                           `json:"orderDate" gorm:"not null"`                                 // Date and time the order was placed
	ReasonRejection        string                              `json:"reasonRejection"`
	PaymentOrders          []PaymentOrder                      `json:"paymentOrders"`                                                               // Payments related to this order
	MenuOrderItems         []MenuOrderItem                     `json:"menuOrderItems"`                                                              // Payments related to this order
	FloorEnvironmentTables []reservation.FloorEnvironmentTable `json:"floorEnvironmentTables" gorm:"many2many:menu_order_floor_environment_tables"` // Many-to-many relationship with FloorEnvironmentTable
	common.ModelTime
	common.ControlBy
}

type MenuOrderItem struct {
	common.ModelId
	MenuOrderID  uint    `json:"orderId" gorm:"not null"`                          // ID de la orden asociada
	ItemID       uint    `json:"dishId" gorm:"not null"`                           // ID del plato asociado
	Quantity     int     `json:"quantity" gorm:"not null;check:quantity > 0"`      // Cantidad del plato en la orden
	UnitPrice    float64 `json:"unitPrice" gorm:"not null;check:unit_price > 0"`   // Precio unitario del plato en la orden
	TotalPrice   float64 `json:"totalPrice" gorm:"not null;check:total_price > 0"` // Precio total del plato en la orden (Cantidad * Precio unitario)
	SpecialNotes string  `json:"specialNotes" gorm:"type:text"`                    // Notas especiales para el plato en la orden
	common.ModelTime
}

// AssignToDefaultCustomer Función para asignar la orden al cliente por defecto si no se proporcionan datos del cliente
func (order *MenuOrder) AssignToDefaultCustomer() {
	if order.CustomerID == 0 {
		order.CustomerID = 1
	}
}

//Pendiente: La orden de menú se ha recibido pero aún no se ha comenzado a preparar.
//Aprobado:La orden de menú ha sido aprobada
//Rechazado: La orden de menú ha sido rechazada por algún motivo.
//En preparación: La orden de menú se está preparando actualmente en la cocina.
//Listo para servir: La orden de menú está preparada y lista para ser servida al cliente.
//Entregado: La orden de menú ha sido entregada al cliente.
//Cancelado: La orden de menú ha sido cancelada por el cliente o el restaurante.
//En espera: La orden de menú está en espera debido a algún problema o circunstancia especial.
//Completado: La orden de menú se ha completado con éxito y se ha entregado al cliente.
