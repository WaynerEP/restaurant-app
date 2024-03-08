package reservation

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"time"
)

type Environment struct {
	common.ModelId
	Name        string `json:"name" gorm:"size:50;unique"`  // Límite de tamaño para el nombre
	Description string `json:"description" gorm:"size:255"` // Límite de tamaño para la descripción
}

type FloorEnvironmentTable struct {
	common.ModelId
	FloorEnvironmentID uint      `json:"floorEnvironmentId" gorm:""`                  // ID del ambiente en el piso
	TableID            uint      `json:"tableId" gorm:"not null"`                     // ID de la mesa
	Table              Table     `json:"table"`                                       //
	SpecificCapacity   int       `json:"specificCapacity" gorm:"not null;default:4"`  // Capacidad específica de la mesa en este ambiente
	Status             string    `json:"status" gorm:"not null;default:'Disponible'"` // Estado de la mesa en este ambiente
	Location           string    `json:"location" gorm:"size:255"`                    // Ubicación física de la mesa en el ambiente
	ImageURL           string    `json:"imageURL" `                                   // URL de la imagen del ambiente
	AdditionalInfo     string    `json:"additionalInfo" gorm:"type:text"`             // Información adicional sobre la mesa en este ambiente
	LastCleanedAt      time.Time `json:"lastCleanedAt"`                               // Última vez que se limpió la mesa en este ambiente
}

type MenuOrderFloorEnvironmentTable struct {
	FloorEnvironmentTableID uint `json:"-"`
	MenuOrderID             uint `json:"-"`
}
