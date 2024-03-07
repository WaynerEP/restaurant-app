package models

import "github.com/WaynerEP/restaurant-app/server/models/common"

type Charge struct {
	common.ModelId
	Name          string      `json:"name" gorm:"size:80; not null; unique;" validate:"required,unique_db"`
	Description   string      `json:"description" gorm:"size:200; not null;"`
	CompanyAreaID uint        `json:"companyAreaId" gorm:"not null;" validate:"required"`
	CompanyArea   CompanyArea `json:"companyArea"`
	Employees     []Employee  `json:"employees"`
	Status        *bool       `json:"status" gorm:"DEFAULT:1;"`
	common.ModelTime
}

type Charges []Charge

func (u *Charge) SetId(ID uint) {
	u.ID = ID
}
