package models

type Currency struct {
	Code         string  `json:"code" gorm:"primaryKey; size:10; NOT NULL" validate:"required,unique_db=ID"`
	Name         string  `json:"name" gorm:"size:30; unique; NOT NULL;" validate:"required,unique_db=ID"`
	Locale       string  `json:"locale" gorm:"size:10; NOT NULL" validate:"required,unique_db=ID"`
	Symbol       string  `json:"symbol" gorm:"size:4; NOT NULL"`
	ExChangeRate float64 `json:"exChangeRate" gorm:"type:decimal(9,2); default:0"`
	Status       *bool   `json:"status" gorm:"default:true;"`
}

type Currencies []Currency

func (u *Currency) SetId(code string) {
	u.Code = code
}
