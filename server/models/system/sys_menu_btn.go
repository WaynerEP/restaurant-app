package system

import "github.com/WaynerEP/restaurant-app/server/models/common"

// SysBaseMenuBtn represents a system base menu button entity.
type SysBaseMenuBtn struct {
	common.ModelId
	Name          string `json:"name" gorm:"not null;size:50;comment:Button key" validate:"required,unique_db"`
	Desc          string `json:"desc" gorm:"Button remark"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:Menu ID"`
	common.ModelTime
}
