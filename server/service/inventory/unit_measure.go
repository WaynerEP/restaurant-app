package inventory

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/inventory"
)

type UnitOfMeasureService struct{}

func (s *UnitOfMeasureService) CreateUnitOfMeasure(e inventory.UnitMeasure) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (s *UnitOfMeasureService) DeleteUnitOfMeasure(e inventory.UnitMeasure) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (s *UnitOfMeasureService) UpdateUnitOfMeasure(e *inventory.UnitMeasure) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (s *UnitOfMeasureService) GetUnitOfMeasure(id uint) (customer inventory.UnitMeasure, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

func (s *UnitOfMeasureService) GetUnitOfMeasureInfoList(info request.PageInfo) (unitMeasureList []inventory.UnitMeasure, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&inventory.UnitMeasure{})
	err = db.Count(&total).Error
	if err != nil {
		return unitMeasureList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&unitMeasureList).Error
	}
	return unitMeasureList, total, err
}
