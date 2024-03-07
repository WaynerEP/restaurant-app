package system

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
)

type JwtBlacklist struct {
	common.ModelId
	Jwt string `gorm:"type:text;comment:jwt"`
	common.ModelTime
}
