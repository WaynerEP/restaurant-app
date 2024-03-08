package reservation

import (
	"errors"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/request"
	"github.com/WaynerEP/restaurant-app/server/models/reservation"
	"gorm.io/gorm"
)

type EnvService struct{}

func (s *EnvService) GetEnvironmentsByFloorId(floorId uint) (list []reservation.FloorEnvironment, err error) {
	err = global.GVA_DB.Where("floor_id = ?", floorId).Preload("Environment").Find(&list).Error
	return list, err
}

func (s *EnvService) CreateFloorEnvironment(e reservation.FloorEnvironment) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (s *EnvService) GetOptionsForSelect() (list []reservation.Environment, err error) {
	err = global.GVA_DB.Select("id", "name", "description").Find(&list).Error
	return list, err
}

func (s *EnvService) CreateEnvironment(e reservation.Environment) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (s *EnvService) DeleteFloorEnvironment(e reservation.FloorEnvironment) error {
	if !errors.Is(global.GVA_DB.Where("floor_environment_id = ?", e.ID).First(&reservation.FloorEnvironmentTable{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("No se puede eliminar el ambiente de piso asociado a varias mesas")
	}
	err := global.GVA_DB.Unscoped().Delete(&e).Error
	return err
}

func (s *EnvService) DeleteEnvironment(e reservation.Environment) (err error) {
	if !errors.Is(global.GVA_DB.Where("environment_id = ?", e.ID).First(&reservation.FloorEnvironment{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("No se puede eliminar el ambiente porque está asociada a un piso")
	}
	err = global.GVA_DB.Unscoped().Delete(&e).Error
	return err
}

func (s *EnvService) UpdateEnvironment(e *reservation.Environment) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (s *EnvService) GetEnvironment(id uint) (Env reservation.Environment, err error) {
	err = global.GVA_DB.First(&Env, id).Error
	return
}

func (s *EnvService) GetEnvironmentInfoList(info request.PageInfo) (list []reservation.Environment, total int64, err error) {
	limit := info.GetLimit()
	page := info.GetPage()
	offset := limit * (page - 1)

	db := global.GVA_DB.Model(&reservation.Environment{})
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	}
	return list, total, err
}
