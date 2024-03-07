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

type AuthorityMenuApi struct{}

// GetMenu .
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c))
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "Retrieved successfully", c)
}

// GetBaseMenuTree .
func (a *AuthorityMenuApi) GetBaseMenuTree(c *gin.Context) {
	menus, err := menuService.GetBaseMenuTree()
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "Retrieved successfully", c)
}

// AddMenuAuthority .
func (a *AuthorityMenuApi) AddMenuAuthority(c *gin.Context) {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	err := c.ShouldBindJSON(&authorityMenu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if verifyErr := utils.Verify(authorityMenu); verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.GVA_LOG.Error("Failed to add!", zap.Error(err))
		response.FailWithMessage("Failed to add", c)
	} else {
		response.OkWithMessage("Added successfully", c)
	}
}

// GetMenuAuthority .
func (a *AuthorityMenuApi) GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(param)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	menus, err := menuService.GetMenuAuthority(&param)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(gin.H{"menus": menus}, "Retrieved successfully", c)
}

// AddBaseMenu .
func (a *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(menu)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	/*	verifyErr = utils.Verify(menu.Meta)
		if err != nil {
			response.FailWithValidationErrors(verifyErr, c)
			return
		}*/
	err = menuService.AddBaseMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("Failed to add!", zap.Error(err))
		response.FailWithMessage("Failed to add", c)
		return
	}
	response.OkWithMessage("Added successfully", c)
}

// DeleteBaseMenu .
func (a *AuthorityMenuApi) DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(menu)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = baseMenuService.DeleteBaseMenu(menu.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to delete!", zap.Error(err))
		response.FailWithMessage("Failed to delete", c)
		return
	}
	response.OkWithMessage("Deleted successfully", c)
}

// UpdateBaseMenu .
func (a *AuthorityMenuApi) UpdateBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(menu)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	/*	verifyErr = utils.Verify(menu.Meta)
		if err != nil {
			response.FailWithValidationErrors(verifyErr, c)
			return
		}*/
	err = baseMenuService.UpdateBaseMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("Failed to update!", zap.Error(err))
		response.FailWithMessage("Failed to update", c)
		return
	}
	response.OkWithMessage("Updated successfully", c)
}

// GetBaseMenuById .
func (a *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(idInfo)
	if err != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	menu, err := baseMenuService.GetBaseMenuById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "Retrieved successfully", c)
}

// GetMenuList .
func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
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
	menuList, total, err := menuService.GetInfoList()
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menuList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
