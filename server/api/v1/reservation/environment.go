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

type EnvApi struct{}

func (e *EnvApi) GetEnvironmentsByFloorId(c *gin.Context) {
	idParam, err := utils.GetIdFromParam(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := envService.GetEnvironmentsByFloorId(idParam)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

func (e *EnvApi) GetOptionsForSelect(c *gin.Context) {
	list, err := envService.GetOptionsForSelect()
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

func (e *EnvApi) CreateFloorEnvironment(c *gin.Context) {
	var floorEnv reservation.FloorEnvironment
	err := c.ShouldBindJSON(&floorEnv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(floorEnv)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = envService.CreateFloorEnvironment(floorEnv)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *EnvApi) CreateEnvironment(c *gin.Context) {
	var env reservation.Environment
	err := c.ShouldBindJSON(&env)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(env)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = envService.CreateEnvironment(env)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *EnvApi) DeleteFloorEnvironment(c *gin.Context) {
	var env reservation.FloorEnvironment
	err := c.ShouldBindJSON(&env)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(env)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = envService.DeleteFloorEnvironment(env)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar: "+err.Error(), c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}
func (e *EnvApi) DeleteEnvironment(c *gin.Context) {
	var env reservation.Environment
	err := c.ShouldBindJSON(&env)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(env)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = envService.DeleteEnvironment(env)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar: "+err.Error(), c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *EnvApi) UpdateEnvironment(c *gin.Context) {
	var env reservation.Environment
	err := c.ShouldBindJSON(&env)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if env.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(env)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = envService.UpdateEnvironment(&env)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *EnvApi) GetEnvironment(c *gin.Context) {
	var env reservation.Environment
	err := c.ShouldBindQuery(&env)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if env.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := envService.GetEnvironment(env.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.EnvironmentResponse{Environment: data}, "Retrieved successfully", c)
}

func (e *EnvApi) GetEnvironmentList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	envList, total, err := envService.GetEnvironmentInfoList(pageInfo)
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
