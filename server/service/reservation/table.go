package reservation

import (
	"errors"

	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/reservation"
	"gorm.io/gorm"
)

type TableService struct{}

func (s *TableService) GetTablesByFloorEnvironmentId(floorEnvId uint) (list []reservation.FloorEnvironmentTable, err error) {
	err = global.GVA_DB.Where("floor_environment_id = ?", floorEnvId).Preload("Table").Find(&list).Error
	return list, err
}

func (s *TableService) GetOptionsForSelect() (list []reservation.Table, err error) {
	err = global.GVA_DB.Select("id", "table_number", "description").Find(&list).Error
	return list, err
}

func (s *TableService) CreateFloorEnvironmentTable(e reservation.FloorEnvironmentTable) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (s *TableService) CreateTable(e reservation.Table) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (s *TableService) DeleteFloorEnvironmentTable(e reservation.FloorEnvironmentTable) error {
	if !errors.Is(global.GVA_DB.Where("floor_environment_table_id = ?", e.ID).First(&reservation.MenuOrderFloorEnvironmentTable{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("No se puede eliminar la mesa porque está asociada a pedidos.")
	}
	err := global.GVA_DB.Unscoped().Delete(&e).Error
	return err
}

func (s *TableService) DeleteTable(e reservation.Table) (err error) {
	if !errors.Is(global.GVA_DB.Where("table_id = ?", e.ID).First(&reservation.FloorEnvironmentTable{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("No se puede eliminar la mesa porque está asociada a un ambiente de piso")
	}
	err = global.GVA_DB.Unscoped().Delete(&e).Error
	return err
}

func (s *TableService) UpdateTable(e *reservation.Table) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (s *TableService) GetTable(id uint) (Table reservation.Table, err error) {
	err = global.GVA_DB.First(&Table, id).Error
	return
}

func (s *TableService) GetTableInfoList(info request.PageInfo) (list []reservation.Table, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&reservation.Table{})
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	}
	return list, total, err
}
