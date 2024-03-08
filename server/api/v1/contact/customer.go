package contact

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	responseCust "github.com/WaynerEP/restaurant-app/server/models/contact/response"

	"github.com/WaynerEP/restaurant-app/server/models/contact"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerApi struct{}

func (e *CustomerApi) CreateCustomer(c *gin.Context) {
	var customerModel contact.Customer
	err := c.ShouldBindJSON(&customerModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(customerModel)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerModel.CreatedBy = utils.GetUserID(c)
	err = customerService.CreateCustomer(customerModel)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *CustomerApi) DeleteCustomer(c *gin.Context) {
	var customerModel contact.Customer
	err := c.ShouldBindJSON(&customerModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(customerModel)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = customerService.DeleteCustomer(customerModel)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *CustomerApi) UpdateCustomer(c *gin.Context) {
	var customerModel contact.Customer
	err := c.ShouldBindJSON(&customerModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if customerModel.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(customerModel)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = customerService.UpdateCustomer(&customerModel)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *CustomerApi) GetCustomer(c *gin.Context) {
	var customerModel contact.Customer
	err := c.ShouldBindQuery(&customerModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if customerModel.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := customerService.GetCustomer(customerModel.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(responseCust.CustomerResponse{Customer: data}, "Retrieved successfully", c)
}

func (e *CustomerApi) GetCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := customerService.GetCustomerInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
