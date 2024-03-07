package common

import (
	"gorm.io/gorm"
	"time"
)

type ModelId struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}

type ControlBy struct {
	CreatedBy uint `json:"createdBy" gorm:"index"`
	UpdatedBy uint `json:"updatedBy" gorm:"index"`
}
