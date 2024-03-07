package system

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"time"
)

// SysOperationRecord If it contains time.Time, please import the time package manually
type SysOperationRecord struct {
	common.ModelId
	Ip           string        `json:"ip" form:"ip" gorm:"comment:Request IP"`                             // Request IP
	Method       string        `json:"method" form:"method" gorm:"comment:Request Method"`                 // Request Method
	Path         string        `json:"path" form:"path" gorm:"comment:Request Path"`                       // Request Path
	Status       int           `json:"status" form:"status" gorm:"comment:Request Enable"`                 // Request Status
	Latency      time.Duration `json:"latency" form:"latency" gorm:"comment:Latency" swaggertype:"string"` // Latency
	Agent        string        `json:"agent" form:"agent" gorm:"comment:Agent"`                            // Agent
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"comment:Error Message"`    // Error Message
	Body         string        `json:"body" form:"body" gorm:"type:text;comment:Request Body"`             // Request Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;comment:Response Body"`            // Response Body
	SysUserID    int           `json:"userId" form:"userId"  gorm:"comment:User ID"`                       // User ID
	SysUser      SysUser       `json:"user" validate:"-"`
	common.ModelTime
}
