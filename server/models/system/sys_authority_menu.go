package system

// SysMenu represents a system menu entity.
type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:Menu ID"`
	AuthorityId uint                   `json:"-" gorm:"comment:Role ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
	Btns        map[string]uint        `json:"btns" gorm:"-"`
}

// SysAuthorityMenu represents the mapping between SysMenu and SysAuthority.
type SysAuthorityMenu struct {
	MenuId      string `json:"menuId" gorm:"comment:Menu ID;column:sys_base_menu_id"`
	AuthorityId string `json:"-" gorm:"comment:Role ID;column:sys_authority_id"`
}

// TableName returns the table name for SysAuthorityMenu.
func (s SysAuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
