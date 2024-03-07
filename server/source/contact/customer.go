package contact

import (
	"context"
	custModel "github.com/WaynerEP/restaurant-app/server/models/contact"
	"github.com/WaynerEP/restaurant-app/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderCustomer = system.InitOrderInternal + 1

type initCustomer struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCustomer, &initCustomer{})
}

func (i *initCustomer) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&custModel.Customer{})
}

func (i *initCustomer) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&custModel.Customer{})
}

func (i *initCustomer) InitializerName() string {
	return custModel.Customer{}.TableName()
}

func (i *initCustomer) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []custModel.Customer{
		{
			DocumentType:   "N/A", // Tipo de documento genérico
			DocumentNumber: "N/A", // Número de documento genérico
			Name:           "Cliente General",
			Lastname:       "Cliente General",
			Email:          "generalcustomer@example.com",
			Phone:          "000000000", // Número de teléfono genérico
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, custModel.Customer{}.TableName()+" table data initialization failed!")
	}
	return next, err
}

func (i *initCustomer) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record custModel.Customer
	if errors.Is(db.First(&record, 1).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return record.ID == 1
}
