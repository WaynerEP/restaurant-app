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

type UnitOfMeasureApi struct{}

func (e *UnitOfMeasureApi) CreateUnitOfMeasure(c *gin.Context) {
	var measure inventory.UnitMeasure
	err := c.ShouldBindJSON(&measure)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(measure)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = unitOfMeasureService.CreateUnitOfMeasure(measure)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *UnitOfMeasureApi) DeleteUnitOfMeasure(c *gin.Context) {
	var measure inventory.UnitMeasure
	err := c.ShouldBindJSON(&measure)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(measure)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = unitOfMeasureService.DeleteUnitOfMeasure(measure)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *UnitOfMeasureApi) UpdateUnitOfMeasure(c *gin.Context) {
	var measure inventory.UnitMeasure
	err := c.ShouldBindJSON(&measure)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if measure.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(measure)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = unitOfMeasureService.UpdateUnitOfMeasure(&measure)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *UnitOfMeasureApi) GetUnitOfMeasure(c *gin.Context) {
	var measure inventory.UnitMeasure
	err := c.ShouldBindQuery(&measure)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if measure.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := unitOfMeasureService.GetUnitOfMeasure(measure.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.UnitOfMeasureResponse{UnitOfMeasure: data}, "Retrieved successfully", c)
}

func (e *UnitOfMeasureApi) GetUnitOfMeasureList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	UnitMeasureList, total, err := unitOfMeasureService.GetUnitOfMeasureInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     UnitMeasureList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
