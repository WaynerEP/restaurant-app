package system

// SysUserAuthority is the junction table between SysUser and SysAuthority.
type SysUserAuthority struct {
	SysUserId               uint `gorm:"column:sys_user_id"`
	SysAuthorityAuthorityId uint `gorm:"column:sys_authority_id"`
}

// TableName returns the table name for SysUserAuthority.
func (s *SysUserAuthority) TableName() string {
	return "sys_user_authority"
}
