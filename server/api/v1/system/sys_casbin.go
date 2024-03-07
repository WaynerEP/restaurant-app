package system

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/system/request"
	systemRes "github.com/WaynerEP/restaurant-app/server/models/system/response"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// UpdateCasbin .
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(cmr)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		global.GVA_LOG.Error("Failed to update!", zap.Error(err))
		response.FailWithMessage("Failed to update", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// GetPolicyPathByAuthorityId .
func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *gin.Context) {
	var casbin request.CasbinInReceive
	err := c.ShouldBindJSON(&casbin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(casbin)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "Obtained successfully", c)
}
