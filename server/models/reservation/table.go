package reservation

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"time"
)

type Table struct {
	common.ModelId
	TableNumber string `json:"tableNumber" gorm:"size:30;not null"`
	Description string `json:"description" gorm:"size:255"`
	common.ModelTime
}

type FloorEnvironmentTable struct {
	common.ModelId
	FloorEnvironmentID uint      `json:"floorEnvironmentId" gorm:""` // ID del ambiente en el piso
	TableID            uint      `json:"tableId" gorm:""`            // ID de la mesa
	Table              Table     `json:"table"`
	SpecificCapacity   int       `json:"specificCapacity" gorm:"not null;default:4"`  // Capacidad específica de la mesa en este ambiente
	Status             string    `json:"status" gorm:"not null;default:'Disponible'"` // Estado de la mesa en este ambiente
	Location           string    `json:"location" gorm:"size:255"`                    // Ubicación física de la mesa en el ambiente
	ImageURL           string    `json:"imageURL" `                                   // URL de la imagen del ambiente
	AdditionalInfo     string    `json:"additionalInfo" gorm:"type:text"`             // Información adicional sobre la mesa en este ambiente
	LastCleanedAt      time.Time `json:"lastCleanedAt"`                               // Última vez que se limpió la mesa en este ambiente
	common.ModelTime
}

/*Piso:
Activo: El piso está abierto y disponible para uso.
Inactivo: El piso está cerrado o fuera de servicio por mantenimiento u otras razones.
En limpieza: El piso está siendo limpiado y no está disponible para uso.
Reservado: El piso ha sido reservado para un evento o función específica.
Ocupado: El piso está ocupado por clientes.
Cerrado: El piso está cerrado temporalmente por razones como feriados o cierres temporales.
Ambiente:
Activo: El ambiente está abierto y disponible para uso.
Inactivo: El ambiente está cerrado o fuera de servicio por mantenimiento u otras razones.
Reservado: El ambiente ha sido reservado para un evento o función específica.
Ocupado: El ambiente está ocupado por clientes.
Cerrado: El ambiente está cerrado temporalmente por razones como feriados o cierres temporales.
Mesa:
Disponible: La mesa está disponible para ser reservada o utilizada.
Reservada: La mesa ha sido reservada para un cliente en un momento específico.
Ocupada: La mesa está siendo utilizada por clientes.
Sucia: La mesa necesita ser limpiada antes de ser utilizada nuevamente.
Fuera de servicio: La mesa está fuera de servicio debido a daños u otros problemas.
Cerrada: La mesa está temporalmente fuera de servicio por razones como feriados o cierres temporales.*/
