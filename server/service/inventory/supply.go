package inventory

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/inventory"
)

type SupplyService struct{}

func (s *SupplyService) CreateSupply(e inventory.Supply) (err error) {
	err = global.GVA_DB.Omit("SupplyCategory").Create(&e).Error
	return err
}

func (s *SupplyService) DeleteSupply(e inventory.Supply) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (s *SupplyService) UpdateSupply(e *inventory.Supply) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (s *SupplyService) GetSupply(id uint) (customer inventory.Supply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

func (s *SupplyService) GetSupplyInfoList(info request.PageInfo) (supplyList []inventory.Supply, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&inventory.Supply{})
	err = db.Count(&total).Error
	if err != nil {
		return supplyList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&supplyList).Error
	}
	return supplyList, total, err
}
