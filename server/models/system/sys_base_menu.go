package system

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
)

// SysBaseMenu represents a system base menu entity.
type SysBaseMenu struct {
	common.ModelId
	MenuLevel      uint                                                        `json:"-"`
	ParentId       string                                                      `json:"parentId" gorm:"comment:Parent menu ID" validate:"required"`                    // Parent menu ID
	Path           string                                                      `json:"path" gorm:"comment:Route path" validate:"required"`                            // Route path
	Name           string                                                      `json:"name" gorm:"comment:Route name" validate:"required,unique_db"`                  // Route name
	Hidden         bool                                                        `json:"hidden" gorm:"comment:Whether hidden in the list"`                              // Whether hidden in the list
	Component      string                                                      `json:"component" gorm:"comment:Corresponding frontend file path" validate:"required"` // Corresponding frontend file path
	Sort           int                                                         `json:"sort" gorm:"comment:Sorting flag" validate:"required,gt=0"`                     // Sorting flag
	Meta           `json:"meta" gorm:"embedded;comment:Additional properties"` // Additional properties
	SysAuthorities []SysAuthority                                              `json:"authorities" gorm:"many2many:sys_authority_menus;"`
	Children       []SysBaseMenu                                               `json:"children" gorm:"-"`
	Parameters     []SysBaseMenuParameter                                      `json:"parameters"`
	MenuBtn        []SysBaseMenuBtn                                            `json:"menuBtn"`
	common.ModelTime
}

// Meta represents additional properties for a system base menu.
type Meta struct {
	ActiveName  string `json:"activeName" gorm:"comment:Highlighted menu"`
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:Whether to cache"`                                  // Whether to cache
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:Whether it is a basic route (under development)"` // Whether it is a basic route (under development)
	Title       string `json:"title" gorm:"comment:Menu name" validate:"required"`                         // Menu name
	Icon        string `json:"icon" gorm:"comment:Menu icon"`                                              // Menu icon
	CloseTab    bool   `json:"closeTab" gorm:"comment:Automatically close tab"`                            // Automatically close tab
}

// SysBaseMenuParameter represents parameters for a system base menu.
type SysBaseMenuParameter struct {
	common.ModelId
	SysBaseMenuID uint   `json:"sysBaseMenuId" gorm:"not null"`
	Type          string `json:"type" gorm:"not null;size:20;comment:Whether the address bar carries parameters as params or query"` // Whether the address bar carries parameters as params or query
	Key           string `json:"key" gorm:"not null;size:20;comment:Key for address bar carrying parameters"`                        // Key for address bar carrying parameters
	Value         string `json:"value" gorm:"not null;size:20;comment:Value for address bar carrying parameters"`                    // Value for address bar carrying parameters
	common.ModelTime
}

// TableName returns the table name for SysBaseMenu.
func (SysBaseMenu) TableName() string {
	return "sys_base_menus"
}
