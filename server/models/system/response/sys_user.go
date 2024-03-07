package response

import "github.com/WaynerEP/restaurant-app/server/models/system"

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	UserData  system.SysUser `json:"userData"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
