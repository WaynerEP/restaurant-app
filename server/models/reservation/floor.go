package reservation

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
)

type Floor struct {
	common.ModelId
	Name              string             `json:"name" gorm:"size:50;not null"`            // Límite de tamaño para el nombre
	Description       string             `json:"description" gorm:"size:255"`             // Límite de tamaño para la descripción
	TableCapacity     int                `json:"tableCapacity" gorm:"not null;default:1"` // Capacidad total de mesas
	Status            string             `json:"status" gorm:"not null;default:'Activo'"` // Estado del piso
	Location          string             `json:"location" gorm:"size:255"`                // Ubicación
	ImageURL          string             `json:"imageURL"`                                // URL de la imagen del piso
	AdditionalInfo    string             `json:"additionalInfo" gorm:"type:text"`         // Información adicional
	FloorEnvironments []FloorEnvironment `json:"floorEnvironment"`
	common.ModelTime
}
