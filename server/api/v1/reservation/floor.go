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

type FloorApi struct{}

func (e *FloorApi) GetTreeFloor(c *gin.Context) {
	list, err := floorService.GetTreeFloorEnvironmentTables()
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

func (e *FloorApi) GetOptionsForSelect(c *gin.Context) {
	list, err := floorService.GetOptionsForSelect()
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

func (e *FloorApi) CreateFloor(c *gin.Context) {
	var floor reservation.Floor
	err := c.ShouldBindJSON(&floor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(floor)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = floorService.CreateFloor(floor)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *FloorApi) DeleteFloor(c *gin.Context) {
	var floor reservation.Floor
	err := c.ShouldBindJSON(&floor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if floor.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	err = floorService.DeleteFloor(floor)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar: "+err.Error(), c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *FloorApi) UpdateFloor(c *gin.Context) {
	var floor reservation.Floor
	err := c.ShouldBindJSON(&floor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if floor.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(floor)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = floorService.UpdateFloor(&floor)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *FloorApi) GetFloor(c *gin.Context) {
	var floor reservation.Floor
	err := c.ShouldBindQuery(&floor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if floor.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := floorService.GetFloor(floor.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.FloorResponse{Floor: data}, "Retrieved successfully", c)
}

func (e *FloorApi) GetFloorList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	floorList, total, err := floorService.GetFloorInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     floorList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
