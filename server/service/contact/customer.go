package contact

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/contact"
)

type CustomerService struct{}

func (exa *CustomerService) CreateCustomer(e contact.Customer) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *CustomerService) DeleteCustomer(e contact.Customer) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (exa *CustomerService) UpdateCustomer(e *contact.Customer) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *CustomerService) GetCustomer(id uint) (customer contact.Customer, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

func (exa *CustomerService) GetCustomerInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&contact.Customer{})
	var customerList []contact.Customer
	err = db.Count(&total).Error
	if err != nil {
		return customerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&customerList).Error
	}
	return customerList, total, err
}
