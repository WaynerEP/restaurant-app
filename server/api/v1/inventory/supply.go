package inventory

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/inventory"
	resModel "github.com/WaynerEP/restaurant-app/server/models/inventory/response"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SupplyApi struct{}

func (e *SupplyApi) CreateSupply(c *gin.Context) {
	var supplyModel inventory.Supply
	err := c.ShouldBindJSON(&supplyModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(supplyModel)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = supplyService.CreateSupply(supplyModel)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *SupplyApi) DeleteSupply(c *gin.Context) {
	var supplyModel inventory.Supply
	err := c.ShouldBindJSON(&supplyModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(supplyModel)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = supplyService.DeleteSupply(supplyModel)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *SupplyApi) UpdateSupply(c *gin.Context) {
	var supplyModel inventory.Supply
	err := c.ShouldBindJSON(&supplyModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if supplyModel.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(supplyModel)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = supplyService.UpdateSupply(&supplyModel)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *SupplyApi) GetSupply(c *gin.Context) {
	var supplyModel inventory.Supply
	err := c.ShouldBindQuery(&supplyModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if supplyModel.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := supplyService.GetSupply(supplyModel.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.SupplyResponse{Supply: data}, "Retrieved successfully", c)
}

func (e *SupplyApi) GetSupplyList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	supplyList, total, err := supplyService.GetSupplyInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     supplyList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
