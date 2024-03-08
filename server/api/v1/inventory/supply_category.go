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

type SupplyCategoryApi struct{}

func (e *SupplyCategoryApi) CreateSupplyCategory(c *gin.Context) {
	var supplyCategory inventory.SupplyCategory
	err := c.ShouldBindJSON(&supplyCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(supplyCategory)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = supplyCategoryService.CreateSupplyCategory(supplyCategory)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *SupplyCategoryApi) DeleteSupplyCategory(c *gin.Context) {
	var supplyCategory inventory.SupplyCategory
	err := c.ShouldBindJSON(&supplyCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(supplyCategory)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = supplyCategoryService.DeleteSupplyCategory(supplyCategory)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *SupplyCategoryApi) UpdateSupplyCategory(c *gin.Context) {
	var supCategory inventory.SupplyCategory
	err := c.ShouldBindJSON(&supCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if supCategory.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(supCategory)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = supplyCategoryService.UpdateSupplyCategory(&supCategory)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *SupplyCategoryApi) GetSupplyCategory(c *gin.Context) {
	var supCategory inventory.SupplyCategory
	err := c.ShouldBindQuery(&supCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if supCategory.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := supplyCategoryService.GetSupplyCategory(supCategory.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.SupplyCategoryResponse{SupplyCategory: data}, "Retrieved successfully", c)
}

func (e *SupplyCategoryApi) GetSupplyCategoryList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	supplyCategoryList, total, err := supplyCategoryService.GetSupplyCategoryInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     supplyCategoryList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
