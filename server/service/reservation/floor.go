package reservation

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/reservation"
)

type FloorService struct{}

func (s *FloorService) GetTreeFloorEnvironmentTables() (list []reservation.Floor, err error) {
	var treeList []reservation.Floor
	err = global.GVA_DB.
		Preload("FloorEnvironments.Environment").
		Preload("FloorEnvironments.FloorEnvironmentTables.Table").
		Find(&treeList).Error
	return treeList, err
}

func (s *FloorService) GetOptionsForSelect() (list []reservation.Floor, err error) {
	err = global.GVA_DB.Select("id", "name", "description").Find(&list).Error
	return list, err
}

func (s *FloorService) CreateFloor(e reservation.Floor) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (s *FloorService) DeleteFloor(e reservation.Floor) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (s *FloorService) UpdateFloor(e *reservation.Floor) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (s *FloorService) GetFloor(id uint) (floor reservation.Floor, err error) {
	err = global.GVA_DB.First(&floor, id).Error
	return
}

func (s *FloorService) GetFloorInfoList(info request.PageInfo) (list []reservation.Floor, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&reservation.Floor{})
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	}
	return list, total, err
}
