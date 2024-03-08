package menu

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/menu"
)

type ItemCategoryService struct{}

func (s *ItemCategoryService) CreateItemCategory(e menu.ItemCategory) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (s *ItemCategoryService) DeleteItemCategory(e menu.ItemCategory) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (s *ItemCategoryService) UpdateItemCategory(e *menu.ItemCategory) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (s *ItemCategoryService) GetItemCategory(id uint) (customer menu.ItemCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

func (s *ItemCategoryService) GetItemCategoryInfoList(info request.PageInfo) (itemsCatList []menu.ItemCategory, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&menu.ItemCategory{})
	err = db.Count(&total).Error
	if err != nil {
		return itemsCatList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&itemsCatList).Error
	}
	return itemsCatList, total, err
}
