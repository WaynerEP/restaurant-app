package system

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"github.com/gofrs/uuid/v5"
)

// SysUser represents a system user entity.
type SysUser struct {
	common.ModelId
	UUID           uuid.UUID      `json:"uuid" gorm:"index;comment:User UUID"`                                                                                             // User UUID
	Username       string         `json:"userName" gorm:"index;not null;size:50;comment:User login name" validate:"required"`                                              // User login name
	Password       string         `json:"-"  gorm:"not null;size:256;comment:User login password"`                                                                         // User login password
	NickName       string         `json:"nickName" gorm:"not null;size:50;default:System user;comment:User nickname"`                                                      // User nickname
	SideMode       string         `json:"sideMode" gorm:"size:20;default:dark;comment:User side theme"`                                                                    // User side theme
	HeaderImg      string         `json:"headerImg" gorm:"default:https://www.sealtightroofingexperts.com/wp-content/uploads/2023/04/avataaars-1.png;comment:User avatar"` // User avatar
	BaseColor      string         `json:"baseColor" gorm:"size:10;default:#fff;comment:Base color"`                                                                        // Base color
	ActiveColor    string         `json:"activeColor" gorm:"size:10;default:#1890ff;comment:Active color"`                                                                 // Active color
	EmployeeId     uint           `json:"employeeId" validate:"required"`                                                                                                  //
	SysAuthorityId uint           `json:"authorityId" gorm:"default:1;comment:User role ID"`                                                                               // User role ID
	SysAuthority   SysAuthority   `json:"authority"`                                                                                                                       //
	Authorities    []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`                                                                                //
	CompanyID      uint           `json:"companyId"`                                                                                                                       //
	Phone          string         `json:"phone"  gorm:"comment:User phone number"`                                                                                         // User phone number
	Email          string         `json:"email"  gorm:"not null;size:256;uniqueIndex;comment:User email"`                                                                  // User email
	Enable         int            `json:"enable" gorm:"default:1;comment:Whether the user is frozen (1 normal, 2 frozen)"`                                                 // Whether the user is frozen (1 normal, 2 frozen)
	common.ModelTime
}

// TableName returns the table name for SysUser.
func (SysUser) TableName() string {
	return "sys_users"
}
