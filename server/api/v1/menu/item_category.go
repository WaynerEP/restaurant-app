package menu

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/menu"
	resModel "github.com/WaynerEP/restaurant-app/server/models/menu/response"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ItemCategoryApi struct{}

func (e *ItemCategoryApi) CreateItemCategory(c *gin.Context) {
	var itemCategory menu.ItemCategory
	err := c.ShouldBindJSON(&itemCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(itemCategory)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = itemCategoryService.CreateItemCategory(itemCategory)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *ItemCategoryApi) DeleteItemCategory(c *gin.Context) {
	var itemCategory menu.ItemCategory
	err := c.ShouldBindJSON(&itemCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(itemCategory)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = itemCategoryService.DeleteItemCategory(itemCategory)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *ItemCategoryApi) UpdateItemCategory(c *gin.Context) {
	var itemCategory menu.ItemCategory
	err := c.ShouldBindJSON(&itemCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if itemCategory.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(itemCategory)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = itemCategoryService.UpdateItemCategory(&itemCategory)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *ItemCategoryApi) GetItemCategory(c *gin.Context) {
	var itemCategory menu.ItemCategory
	err := c.ShouldBindQuery(&itemCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if itemCategory.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := itemCategoryService.GetItemCategory(itemCategory.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.ItemCategoryResponse{ItemCategory: data}, "Retrieved successfully", c)
}

func (e *ItemCategoryApi) GetItemCategoryList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ItemCategoryList, total, err := itemCategoryService.GetItemCategoryInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     ItemCategoryList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
