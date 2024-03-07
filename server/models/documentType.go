package models

import "github.com/WaynerEP/restaurant-app/server/models/common"

type DocumentType struct {
	common.ModelId
	Code         string `json:"code" gorm:"size:8; not null;"`
	Description  string `json:"description" gorm:"size:80; not null; unique;"`
	Abbreviation string `json:"abbreviation" gorm:"size:50; not null; unique;"`
	MinChars     int32  `json:"minChars" gorm:"not null;"`
	MaxChars     int32  `json:"maxChars" gorm:"not null;"`
	Status       *bool  `json:"status" gorm:"default:true;"`
}

type DocumentTypes []DocumentType
