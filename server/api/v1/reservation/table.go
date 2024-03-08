package reservation

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/reservation"
	resModel "github.com/WaynerEP/restaurant-app/server/models/reservation/response"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TableApi struct{}

func (e *TableApi) GetTablesByFloorEnvironmentId(c *gin.Context) {
	idParam, err := utils.GetIdFromParam(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := tableService.GetTablesByFloorEnvironmentId(idParam)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

func (e *TableApi) GetOptionsForSelect(c *gin.Context) {
	list, err := tableService.GetOptionsForSelect()
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

func (e *TableApi) CreateFloorEnvTable(c *gin.Context) {
	var table reservation.FloorEnvironmentTable
	err := c.ShouldBindJSON(&table)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(table)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tableService.CreateFloorEnvironmentTable(table)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *TableApi) CreateTable(c *gin.Context) {
	var table reservation.Table
	err := c.ShouldBindJSON(&table)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(table)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = tableService.CreateTable(table)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *TableApi) DeleteFloorEnvironmentTable(c *gin.Context) {
	var tbl reservation.FloorEnvironmentTable
	err := c.ShouldBindJSON(&tbl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(tbl)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = tableService.DeleteFloorEnvironmentTable(tbl)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar: "+err.Error(), c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}
func (e *TableApi) DeleteTable(c *gin.Context) {
	var table reservation.Table
	err := c.ShouldBindJSON(&table)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(table)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = tableService.DeleteTable(table)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar: "+err.Error(), c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *TableApi) UpdateTable(c *gin.Context) {
	var table reservation.Table
	err := c.ShouldBindJSON(&table)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if table.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(table)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = tableService.UpdateTable(&table)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *TableApi) GetTable(c *gin.Context) {
	var table reservation.Table
	err := c.ShouldBindQuery(&table)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if table.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := tableService.GetTable(table.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.TableResponse{Table: data}, "Retrieved successfully", c)
}

func (e *TableApi) GetTableList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	envList, total, err := tableService.GetTableInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     envList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
