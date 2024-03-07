package system

// SysAuthorityBtn represents a system authority button entity.
type SysAuthorityBtn struct {
	AuthorityId      uint           `gorm:"comment:Role ID"`
	SysMenuID        uint           `gorm:"comment:Menu ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:Menu button ID"`
	SysBaseMenuBtn   SysBaseMenuBtn `gorm:"comment:Button details"`
}
