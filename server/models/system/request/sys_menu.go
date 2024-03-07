package request

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"github.com/WaynerEP/restaurant-app/server/models/system"
)

// AddMenuAuthorityInfo Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus" validate:"-"`
	AuthorityId uint                 `json:"authorityId" validate:"required"`
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		ModelId:   common.ModelId{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "Dashboard",
			Icon:  "setting",
		},
	}}
}
