package system

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
)

// SysApi represents a system API entity.
type SysApi struct {
	common.ModelId
	Path        string `json:"path" gorm:"not null;unique;size:256;comment:API path" validate:"required"`                                                                                                                           // API path
	Description string `json:"description" gorm:"size:256;comment:Description of the API" validate:"required"`                                                                                                                      // Description of the API
	ApiGroup    string `json:"apiGroup" gorm:"size:256;not null;comment:API group" validate:"required"`                                                                                                                             // API group
	Method      string `json:"method" gorm:"not null;size:15;default:POST;check:method in ('GET','POST','PUT','PATCH', 'DELETE','OPTIONS','HEAD');comment:Method" validate:"required,oneof=GET POST PUT PATCH DELETE OPTIONS HEAD"` // Method: Create POST (default) | View GET | Update PUT | Delete DELETE
	common.ModelTime
}

func (SysApi) TableName() string {
	return "sys_apis"
}
