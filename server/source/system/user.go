package system

import (
	"context"
	sysModel "github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/service/system"
	"github.com/WaynerEP/restaurant-app/server/utils"

	"github.com/gofrs/uuid/v5"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderUser = initOrderAuthority + 1

type initUser struct{}

// auto run
func init() {
	system.RegisterInit(initOrderUser, &initUser{})
}

func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysUser{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i *initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	password := utils.BcryptHash("12345678")
	adminPassword := utils.BcryptHash("12345678")

	entities := []sysModel.SysUser{
		{
			UUID:           uuid.Must(uuid.NewV4()),
			Username:       "admin",
			Password:       adminPassword,
			NickName:       "Usuario Admin",
			HeaderImg:      "",
			EmployeeId:     1,
			SysAuthorityId: 1,
			CompanyID:      1,
			Phone:          "987458978",
			Email:          "admin@gmail.com",
		},
		{
			UUID:           uuid.Must(uuid.NewV4()),
			Username:       "guest",
			Password:       password,
			NickName:       "Usuario Invitado",
			HeaderImg:      "",
			EmployeeId:     2,
			SysAuthorityId: 2,
			CompanyID:      1,
			Phone:          "987458965",
			Email:          "guest@gmail.com",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+" table data initialization failed!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	authorityEntities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "Failed to create [User-Authority] association, could not find authority table initialization data")
	}
	if err = db.Model(&entities[0]).Association("Authorities").Replace(authorityEntities); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("Authorities").Replace(authorityEntities[:1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysUser
	if errors.Is(db.Where("username = ?", "admin").
		Preload("Authorities").First(&record).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return len(record.Authorities) > 0 && record.Authorities[0].ID == 1
}
