package system

import (
	v1 "github.com/WaynerEP/restaurant-app/server/api/v1"
	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	operationRecordApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("createSysOperationRecord", operationRecordApi.CreateSysOperationRecord)             // Create SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecord", operationRecordApi.DeleteSysOperationRecord)           // Delete SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", operationRecordApi.DeleteSysOperationRecordByIds) // Batch delete SysOperationRecord
		operationRecordRouter.GET("findSysOperationRecord", operationRecordApi.FindSysOperationRecord)                  // Get SysOperationRecord by ID
		operationRecordRouter.GET("getSysOperationRecordList", operationRecordApi.GetSysOperationRecordList)            // Get SysOperationRecord list
	}
}
