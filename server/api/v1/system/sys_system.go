package system

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	systemRes "github.com/WaynerEP/restaurant-app/server/models/system/response"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CSystemApi struct{}

// GetSystemConfig retrieves the system configuration.
func (s *CSystemApi) GetSystemConfig(c *gin.Context) {
	config, err := systemConfigService.GetSystemConfig()
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve system configuration!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve system configuration", c)
		return
	}
	response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "Retrieved successfully", c)
}

// SetSystemConfig sets the content of the configuration file.
func (s *CSystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	err := c.ShouldBindJSON(&sys)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = systemConfigService.SetSystemConfig(sys)
	if err != nil {
		global.GVA_LOG.Error("Failed to set system configuration!", zap.Error(err))
		response.FailWithMessage("Failed to set system configuration", c)
		return
	}
	response.OkWithMessage("Set successfully", c)
}

// ReloadSystem restarts the system.
func (s *CSystemApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.GVA_LOG.Error("Failed to restart the system!", zap.Error(err))
		response.FailWithMessage("Failed to restart the system", c)
		return
	}
	response.OkWithMessage("System restarted successfully", c)
}

// GetServerInfo retrieves information about the server.
func (s *CSystemApi) GetServerInfo(c *gin.Context) {
	server, err := systemConfigService.GetServerInfo()
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve server information!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve server information", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "Retrieved successfully", c)
}
