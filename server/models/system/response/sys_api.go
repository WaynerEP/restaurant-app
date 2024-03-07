package response

import "github.com/WaynerEP/restaurant-app/server/models/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}
