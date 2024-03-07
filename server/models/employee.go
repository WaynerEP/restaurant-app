package models

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"time"
)

type Employee struct {
	common.ModelId
	PersonID     uint       `json:"person_id"`
	Person       Person     `json:"person"`
	JobEmail     string     `json:"jobEmail" gorm:"size:255; not null; unique;"`
	Idiom        string     `json:"idiom" gorm:"size:20;"`
	OfficeID     uint       `json:"officeID" gorm:"not null"`
	StartWork    time.Time  `json:"startWork" gorm:"default:now();"`
	EndWork      *time.Time `json:"endWork"`
	ChargeID     uint       `json:"chargeID" gorm:"not null"`
	EmployeeCode string     `json:"employeeCode" gorm:"size:30; not null"`
	Avatar       *string    `json:"avatar"`
	Status       *bool      `json:"status" gorm:"not null;default:true"`
	common.ControlBy
	common.ModelTime
}

// Employees Posibles estados de un empleado:
// - 1; Activo: El empleado está trabajando actualmente en la empresa.
// - 0; Inactivo: El empleado no está trabajando en la empresa en este momento.

type Employees []Employee

func (u *Employee) SetId(ID uint) {
	u.ID = ID
}
