package models

type EconomicActivity struct {
	Code        string `json:"code" gorm:"size:10; primaryKey; not null"`
	Description string `json:"description" gorm:"size:250; not null;"`
	Status      *bool  `json:"status,omitempty" gorm:"default:true;"`
}

type EconomicActivities []EconomicActivity
