package system

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	systemRes "github.com/WaynerEP/restaurant-app/server/models/system/response"
	"github.com/WaynerEP/restaurant-app/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

// CreateAuthority creates a new authority.
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority, authBack system.SysAuthority
	var err error

	if err = c.ShouldBindJSON(&authority); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if verifyErr := utils.Verify(authority); verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if authority.ParentId == nil {
		authority.ParentId = utils.Pointer[uint](0)
	}
	if authBack, err = authorityService.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed"+err.Error(), c)
		return
	}
	err = casbinService.FreshCasbin()
	if err != nil {
		global.GVA_LOG.Error("Creation succeeded, but failed to refresh permissions.", zap.Error(err))
		response.FailWithMessage("La creación se realizó correctamente, pero no se pudieron actualizar los permisos.."+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "Creation succeeded", c)
}

// CopyAuthority copies an authority.
func (a *AuthorityApi) CopyAuthority(c *gin.Context) {
	var copyInfo systemRes.SysAuthorityCopyResponse
	err := c.ShouldBindJSON(&copyInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(copyInfo)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	authBack, err := authorityService.CopyAuthority(copyInfo)
	if err != nil {
		global.GVA_LOG.Error("Copy failed!", zap.Error(err))
		response.FailWithMessage("Copy failed: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "Copy succeeded", c)
}

// DeleteAuthority deletes an authority.
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority system.SysAuthority
	var err error
	if err = c.ShouldBindJSON(&authority); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authority.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	// Before deleting a role, check if any users are currently using this role.
	if err = authorityService.DeleteAuthority(&authority); err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed: "+err.Error(), c)
		return
	}
	_ = casbinService.FreshCasbin()
	response.OkWithMessage("Deletion succeeded", c)
}

// UpdateAuthority updates an authority.
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var auth system.SysAuthority
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	inputErrs := utils.Verify(auth)
	if inputErrs != nil {
		response.FailWithValidationErrors(inputErrs, c)
		return
	}
	if auth.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	authority, err := authorityService.UpdateAuthority(auth)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "Update succeeded", c)
}

// GetAuthorityList gets a list of authorities.
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := authorityService.GetAuthorityInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to get!", zap.Error(err))
		response.FailWithMessage("Failed to get"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get succeeded", c)
}

// SetDataAuthority sets data authority.
func (a *AuthorityApi) SetDataAuthority(c *gin.Context) {
	var auth system.SysAuthority
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if auth.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	err = authorityService.SetDataAuthority(auth)
	if err != nil {
		global.GVA_LOG.Error("Setting failed!", zap.Error(err))
		response.FailWithMessage("Setting failed"+err.Error(), c)
		return
	}
	response.OkWithMessage("Setting succeeded", c)
}
