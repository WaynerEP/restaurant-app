package order

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	requestPage "github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/order"
	"github.com/WaynerEP/restaurant-app/server/models/order/request"
	resModel "github.com/WaynerEP/restaurant-app/server/models/order/response"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuOrderApi struct{}

func (e *MenuOrderApi) ReadyMenuOrder(c *gin.Context) {
	var statusOrderReq request.OrderStatusRequest
	err := c.ShouldBindJSON(&statusOrderReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if statusOrderReq.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	updatedBy := utils.GetUserID(c)
	err = menuOrderService.ReadyMenuOrder(statusOrderReq.ID, updatedBy)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update Status successful", c)
}
func (e *MenuOrderApi) ApproveMenuOrder(c *gin.Context) {
	var statusOrderReq request.OrderStatusRequest
	err := c.ShouldBindJSON(&statusOrderReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if statusOrderReq.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	updatedBy := utils.GetUserID(c)
	err = menuOrderService.ApproveMenuOrder(statusOrderReq.ID, updatedBy)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update Status successful", c)
}

func (e *MenuOrderApi) RejectMenuOrder(c *gin.Context) {
	var statusOrderReq request.OrderStatusRequest
	err := c.ShouldBindJSON(&statusOrderReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(statusOrderReq)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	updatedBy := utils.GetUserID(c)
	err = menuOrderService.RejectMenuOrder(statusOrderReq.ID, statusOrderReq.ReasonRejection, updatedBy)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update Status successful", c)
}
func (e *MenuOrderApi) UpdateStatusMenuOrder(c *gin.Context) {
	var statusOrderReq request.OrderStatusRequest
	err := c.ShouldBindJSON(&statusOrderReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(statusOrderReq)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = menuOrderService.UpdateStatusMenuOrder(statusOrderReq.ID, statusOrderReq.Status)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update Status successful", c)
}

func (e *MenuOrderApi) CreateMenuOrder(c *gin.Context) {
	var menuOrder order.MenuOrder
	err := c.ShouldBindJSON(&menuOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(menuOrder)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	menuOrder.CreatedBy = utils.GetUserID(c)
	err = menuOrderService.CreateMenuOrder(menuOrder)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Creation successful", c)
}

func (e *MenuOrderApi) DeleteMenuOrder(c *gin.Context) {
	var menuOrder order.MenuOrder
	err := c.ShouldBindJSON(&menuOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verifyErr := utils.Verify(menuOrder)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	err = menuOrderService.DeleteMenuOrder(menuOrder)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Error al eliminar", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

func (e *MenuOrderApi) UpdateMenuOrder(c *gin.Context) {
	var menuOrder order.MenuOrder
	err := c.ShouldBindJSON(&menuOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if menuOrder.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	verifyErr := utils.Verify(menuOrder)
	if verifyErr != nil {
		response.FailWithValidationErrors(verifyErr, c)
		return
	}
	menuOrder.UpdatedBy = utils.GetUserID(c)
	err = menuOrderService.UpdateMenuOrder(&menuOrder)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

func (e *MenuOrderApi) GetMenuOrder(c *gin.Context) {
	var menuOrder order.MenuOrder
	err := c.ShouldBindQuery(&menuOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if menuOrder.ID <= 0 {
		response.FailWithMessage("Se requiere un identificador válido para la operación", c)
		return
	}
	data, err := menuOrderService.GetMenuOrder(menuOrder.ID)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve", c)
		return
	}
	response.OkWithDetailed(resModel.MenuOrderResponse{MenuOrder: data}, "Retrieved successfully", c)
}

func (e *MenuOrderApi) GetMenuOrderList(c *gin.Context) {
	var pageInfo requestPage.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menuOrderList, total, err := menuOrderService.GetMenuOrderInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve!", zap.Error(err))
		response.FailWithMessage("Failed to retrieve "+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menuOrderList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Retrieved successfully", c)
}
