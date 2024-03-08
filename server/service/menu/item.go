package menu

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/menu"
)

type ItemService struct{}

func (s *ItemService) CreateItem(e menu.Item) (err error) {
	err = global.GVA_DB.Omit("ItemCategory").Create(&e).Error
	return err
}

func (s *ItemService) DeleteItem(e menu.Item) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (s *ItemService) UpdateItem(e *menu.Item) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (s *ItemService) GetItem(id uint) (item menu.Item, err error) {
	err = global.GVA_DB.Preload("NutritionalValue").Preload("ItemCategory").First(&item, id).Error
	return
}

func (s *ItemService) GetItemInfoList(info request.PageInfo) (itemList []menu.Item, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&menu.Item{})
	err = db.Count(&total).Error
	if err != nil {
		return itemList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	}
	return itemList, total, err
}
