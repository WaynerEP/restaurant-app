package system

import (
	"context"
	sysModel "github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/service/system"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderAuthority = initOrderCasbin + 1

type initAuthority struct{}

// auto run
func init() {
	system.RegisterInit(initOrderAuthority, &initAuthority{})
}

func (i *initAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysAuthority{})
}

func (i *initAuthority) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysAuthority{})
}

func (i initAuthority) InitializerName() string {
	return sysModel.SysAuthority{}.TableName()
}

func (i *initAuthority) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysAuthority{
		{ID: 1, AuthorityName: "Administrador", ParentId: utils.Pointer[uint](0), DefaultRouter: "dashboard"},
		{ID: 2, AuthorityName: "Empleado", ParentId: utils.Pointer[uint](0), DefaultRouter: "dashboard"},
		{ID: 3, AuthorityName: "Invitado", ParentId: utils.Pointer[uint](0), DefaultRouter: "dashboard"},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "Failed to initialize data for %s table!", sysModel.SysAuthority{}.TableName())
	}
	// data authority
	if err := db.Model(&entities[0]).Association("DataAuthorityId").Replace(
		[]*sysModel.SysAuthority{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}); err != nil {
		return ctx, errors.Wrapf(err, "Failed to initialize data for %s table!",
			db.Model(&entities[0]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	}
	if err := db.Model(&entities[1]).Association("DataAuthorityId").Replace(
		[]*sysModel.SysAuthority{
			{ID: 2},
			{ID: 3},
		}); err != nil {
		return ctx, errors.Wrapf(err, "Failed to initialize data for %s table!",
			db.Model(&entities[1]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	}

	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("id = ?", "3").
		First(&sysModel.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
