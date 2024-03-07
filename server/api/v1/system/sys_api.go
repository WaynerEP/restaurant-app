package system

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	systemReq "github.com/WaynerEP/restaurant-app/server/models/system/request"
	systemRes "github.com/WaynerEP/restaurant-app/server/models/system/response"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApiApi struct{}

// CreateApi creates a basic API.
func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(api)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = apiService.CreateApi(api)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation succeeded", c)
}

// DeleteApi deletes an API.
func (s *SystemApiApi) DeleteApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if api.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	err = apiService.DeleteApi(api)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion succeeded", c)
}

// GetApiList gets a paginated list of APIs.
func (s *SystemApiApi) GetApiList(c *gin.Context) {
	var pageInfo systemReq.SearchApiParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.Page <= 0 {
		pageInfo.Page = 1
	}
	switch {
	case pageInfo.PageSize > 100:
		pageInfo.PageSize = 100
	case pageInfo.PageSize <= 0:
		pageInfo.PageSize = 10
	}
	list, total, err := apiService.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("Failed to get!", zap.Error(err))
		response.FailWithMessage("Failed to get", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get succeeded", c)
}

// GetApiById gets an API by ID.
func (s *SystemApiApi) GetApiById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(idInfo)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	api, err := apiService.GetApiById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to get!", zap.Error(err))
		response.FailWithMessage("Failed to get", c)
		return
	}
	response.OkWithDetailed(systemRes.SysAPIResponse{Api: api}, "Get succeeded", c)
}

// UpdateApi updates a basic API.
func (s *SystemApiApi) UpdateApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(api)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = apiService.UpdateApi(api)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update succeeded", c)
}

// GetAllApis gets all APIs without pagination.
func (s *SystemApiApi) GetAllApis(c *gin.Context) {
	apis, err := apiService.GetAllApis()
	if err != nil {
		global.GVA_LOG.Error("Failed to get!", zap.Error(err))
		response.FailWithMessage("Failed to get", c)
		return
	}
	response.OkWithDetailed(systemRes.SysAPIListResponse{Apis: apis}, "Get succeeded", c)
}

// DeleteApisByIds deletes selected APIs by IDs.
func (s *SystemApiApi) DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.DeleteApisByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion succeeded", c)
}

// FreshCasbin refreshes the Casbin cache.
func (s *SystemApiApi) FreshCasbin(c *gin.Context) {
	err := casbinService.FreshCasbin()
	if err != nil {
		global.GVA_LOG.Error("Refresh failed!", zap.Error(err))
		response.FailWithMessage("Refresh failed", c)
		return
	}
	response.OkWithMessage("Refresh succeeded", c)
}
