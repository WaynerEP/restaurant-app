package request

import (
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
