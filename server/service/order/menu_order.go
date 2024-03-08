package order

import (
	"errors"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/order"
	"github.com/WaynerEP/restaurant-app/server/utils"
)

var orderStatus = []string{
	"Pendiente",
	"Aprobado",
	"Rechazado",
	"En preparación",
	"Listo para servir",
	"Entregado",
	"Cancelado",
	"En espera",
	"Completado",
}

type MenuOrderService struct{}

func (s *MenuOrderService) ReadyMenuOrder(orderId uint, updatedBy uint) (err error) {
	err = global.GVA_DB.Model(&order.MenuOrder{}).
		Where("id = ? AND status = ?", orderId, "Pendiente").
		Updates(map[string]interface{}{"status": "Listo para servir", "updatedBy": updatedBy}).Error
	return err
}

func (s *MenuOrderService) ApproveMenuOrder(orderId uint, updatedBy uint) (err error) {
	err = global.GVA_DB.Model(&order.MenuOrder{}).
		Where("id = ? AND status = ?", orderId, "Pendiente").
		Updates(map[string]interface{}{"status": "Aprobado", "updatedBy": updatedBy}).Error
	return err
}

func (s *MenuOrderService) RejectMenuOrder(orderId uint, reasonRejection string, updatedBy uint) (err error) {
	err = global.GVA_DB.Model(&order.MenuOrder{}).
		Where("id = ? AND status = ?", orderId, "Pendiente").
		Updates(map[string]interface{}{"status": "Rechazado", "reasonRejection": reasonRejection, "updatedBy": updatedBy}).Error
	return err
}

func (s *MenuOrderService) UpdateStatusMenuOrder(id uint, newStatus string) (err error) {
	if !utils.InArray(orderStatus, newStatus) {
		return errors.New("Invalid order status: " + newStatus)
	}
	err = global.GVA_DB.Model(&order.MenuOrder{}).
		Where("id = ?", id).
		Update("status", newStatus).Error
	return err
}

func (s *MenuOrderService) CreateMenuOrder(order order.MenuOrder) (err error) {
	order.AssignToDefaultCustomer()
	var subtotal float64
	for _, item := range order.MenuOrderItems {
		subtotal += float64(item.Quantity) * item.UnitPrice
	}
	order.Subtotal = subtotal
	order.Discount = 0
	order.Taxes = 0
	order.Status = GetValidOrderStatus(order.Status)
	order.Total = subtotal

	err = global.GVA_DB.Create(&order).Error
	return err
}

func (s *MenuOrderService) DeleteMenuOrder(e order.MenuOrder) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (s *MenuOrderService) UpdateMenuOrder(order *order.MenuOrder) (err error) {
	order.AssignToDefaultCustomer()
	var subtotal float64
	for _, item := range order.MenuOrderItems {
		subtotal += float64(item.Quantity) * item.UnitPrice
	}
	order.Subtotal = subtotal
	order.Discount = 0
	order.Taxes = 0
	order.Status = GetValidOrderStatus(order.Status)
	order.Total = subtotal
	err = global.GVA_DB.Save(order).Error
	return err
}

func (s *MenuOrderService) GetMenuOrder(id uint) (menuOrder order.MenuOrder, err error) {
	err = global.GVA_DB.First(&menuOrder, id).Error
	return
}

func (s *MenuOrderService) GetMenuOrderInfoList(info request.PageInfo) (list []order.MenuOrder, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&order.MenuOrder{})
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	}
	return list, total, err
}

func GetValidOrderStatus(status string) string {
	if !utils.InArray(orderStatus, status) {
		return orderStatus[0]
	}
	return status
}
