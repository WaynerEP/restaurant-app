package inventory

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/inventory"
)

type SupplyCategoryService struct{}

func (s *SupplyCategoryService) CreateSupplyCategory(e inventory.SupplyCategory) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (s *SupplyCategoryService) DeleteSupplyCategory(e inventory.SupplyCategory) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (s *SupplyCategoryService) UpdateSupplyCategory(e *inventory.SupplyCategory) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (s *SupplyCategoryService) GetSupplyCategory(id uint) (supplyCategory inventory.SupplyCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&supplyCategory).Error
	return
}

func (s *SupplyCategoryService) GetSupplyCategoryInfoList(info request.PageInfo) (categoryList []inventory.SupplyCategory, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&inventory.SupplyCategory{})
	err = db.Count(&total).Error
	if err != nil {
		return categoryList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&categoryList).Error
	}
	return categoryList, total, err
}
