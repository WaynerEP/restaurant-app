package reservation

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
)

type Floor struct {
	common.ModelId
	Name              string             `json:"name" gorm:"size:50;not null;unique" validate:"unique_db"` // Límite de tamaño para el nombre
	Description       string             `json:"description" gorm:"size:255"`                              // Límite de tamaño para la descripción
	TableCapacity     int                `json:"tableCapacity" gorm:"not null;default:1"`                  // Capacidad total de mesas
	Status            string             `json:"status" gorm:"not null;default:'Activo'"`                  // Estado del piso
	Location          string             `json:"location" gorm:"size:255"`                                 // Ubicación
	ImageURL          string             `json:"imageURL"`                                                 // URL de la imagen del piso
	AdditionalInfo    string             `json:"additionalInfo" gorm:"type:text"`                          // Información adicional
	FloorEnvironments []FloorEnvironment `json:"floorEnvironments"`
	common.ModelTime
}

type FloorEnvironment struct {
	common.ModelId
	FloorID                uint                    `json:"floorId" gorm:"not null"`       // ID del piso
	EnvironmentID          uint                    `json:"environmentId" gorm:"not null"` // ID del ambiente
	Environment            Environment             `json:"environment"`
	TableCapacity          int                     `json:"tableCapacity" gorm:"not null;default:1"` // Capacidad total de mesas
	Status                 string                  `json:"status" gorm:"not null;default:'Activo'"` // Estado actual del ambiente en este piso
	Location               string                  `json:"location" gorm:"size:255"`                // Ubicación específica del ambiente en el piso
	ImageURL               string                  `json:"imageURL"`                                // URL de la imagen del ambiente
	FloorEnvironmentTables []FloorEnvironmentTable `json:"floorEnvironmentTables"`
}
