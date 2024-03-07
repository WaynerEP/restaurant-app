package models

type UnitMeasure struct {
	Code   string `json:"code" gorm:"size:10;primaryKey;not null"`
	Name   string `json:"name" gorm:"size:100; not null; unique;"`
	Status *bool  `json:"status" gorm:"default:true"`
}

type UnitsMeasure []UnitMeasure
