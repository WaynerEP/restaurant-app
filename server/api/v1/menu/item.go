package menu

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	resModel "github.com/WaynerEP/restaurant-app/server/models/menu/response"

	"github.com/WaynerEP/restaurant-app/server/models/menu"
	"github.com/WaynerEP/restaurant-app/server/utils"
)

type ItemApi struct{}

func (e *ItemApi) CreateItem(c *gin.Context) {
	var item menu.Item
	err := c.ShouldBindJSON(&item)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(item)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = itemService.CreateItem(item)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *ItemApi) DeleteItem(c *gin.Context) {
	var item menu.Item
	err := c.ShouldBindJSON(&item)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(item)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = itemService.DeleteItem(item)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *ItemApi) UpdateItem(c *gin.Context) {
	var item menu.Item
	err := c.ShouldBindJSON(&item)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if item.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(item)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = itemService.UpdateItem(&item)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *ItemApi) GetItem(c *gin.Context) {
	var item menu.Item
	err := c.ShouldBindQuery(&item)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if item.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := itemService.GetItem(item.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.ItemResponse{Item: data}, "Retrieved successfully", c)
}

func (e *ItemApi) GetItemList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ItemList, total, err := itemService.GetItemInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     ItemList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
