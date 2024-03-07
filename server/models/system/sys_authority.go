package system

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
)

// SysAuthority represents a system authority entity.
type SysAuthority struct {
	ID              uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	AuthorityName   string          `json:"authorityName" gorm:"not null;size:80;comment:Role name"  validate:"required,unique_db=authority_name"` // Role name
	ParentId        *uint           `json:"parentId" gorm:"comment:Parent role ID"`                                                                // Parent role ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []SysUser       `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:Default menu;default:dashboard"` // Default menu (default: dashboard)
	common.ModelTime
}

// TableName returns the table name for SysAuthority.
func (SysAuthority) TableName() string {
	return "sys_authorities"
}
