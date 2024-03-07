package response

import "github.com/WaynerEP/restaurant-app/server/models/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	ID             uint                `json:"id"`
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId uint                `json:"oldAuthorityId" validate:"required"`
}
