package models

import (
	"time"

	"github.com/WaynerEP/restaurant-app/server/models/common"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" form:"name" gorm:"size:150;not null;unique" validate:"required,min=3,unique_db"`
	Description string         `json:"description" form:"description"`
	Slug        string         `json:"slug" form:"slug" gorm:"size:250; ;not null"`
	ParentId    uint           `json:"parentId" form:"parentId"`
	Image       string         `json:"image" form:"image"`
	Status      *bool          `json:"status" form:"status" gorm:"default:true"`
	Products    []Product      `json:"products,omitempty"`
	CreatedAt   time.Time      `json:"createdAt,omitempty" form:"-"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty" form:"-"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty" form:"-"`
}

type CategoryList struct {
	Category
	Total int64 `json:"total"`
}

type Categories []CategoryList

func (u *Category) SetId(ID uint) {
	u.ID = ID
}

type Subcategory struct {
	common.ModelId
	Name       string `json:"name" gorm:"size:150; not null; unique;" validate:"required,unique_db"`
	Status     *bool  `json:"status" gorm:"default:true;"`
	CategoryID uint   `json:"categoryId"`
	common.ModelTime
}

type Subcategories []Subcategory
