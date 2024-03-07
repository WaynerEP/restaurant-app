package system

import (
	"context"
	sysModel "github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/service/system"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initCompany struct{}

const initOrderCompany = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderCompany, &initCompany{})
}

func (i initCompany) InitializerName() string {
	return sysModel.Company{}.TableName()
}

func (i *initCompany) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.Company{})
}

func (i *initCompany) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.Company{})
}

func (i *initCompany) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.Company{
		{
			ID:                   1,
			RUC:                  "20606730072",
			Logo:                 "",
			CompanyName:          "ECOSECURITY CONSULTING S.A.C.",
			TradeName:            "ECO CONSULTING",
			PhoneCompany:         "+51 964 200 304",
			EmailCompany:         "ecosecurit.consulting@gmail.com",
			FiscalAddress:        "CAL. RAMÓN CASTILLO NRO. 168 A.H. VILLA HERMOSA LA LIBERTAD CHEPEN CHEPEN",
			Website:              "https://www.ecosecurity.com",
			RegimenID:            1,
			EconomicActivity:     "XXXXX",
			ActivitySecondaryOne: "XXXXX",
			ActivitySecondaryTwo: "XXXXX",
			//Banks:                   datatypes.JSON([]byte(`[{"bank": "BANCO DE LA NACIÓN", "description": "CTA. DETRACCIÓN", "bankAccount": "00-812-039660", "interbankAccount": "188112000812039000000"}]`)),
			Regime:                  "Régimen especial",
			ApplicationVersion:      "peru",
			Timezone:                "America/Lima",
			Profile:                 "Negocio",
			DecimalPrecision:        2,
			CalculationScale:        "2",
			MultiTax:                utils.Pointer[bool](false),
			EmployeesNumber:         20,
			Sector:                  "Otros",
			ShowInvoiceTotalInWords: "yes",
			ShowRetentionInvoice:    "yes",
			ShowNewLineCharOnPdf:    "no",
			ShowItemReferenceOnPdf:  false,
			Coupon:                  "",
			CurrencyCode:            "PEN",
			MultiCurrency:           true,
			DecimalSeparator:        ",",
			CompanySetting: sysModel.CompanySetting{
				CanStampInvoices:                    utils.Pointer[bool](false),
				ElectronicInvoicing:                 utils.Pointer[bool](false),
				IsTypesAndNumerationsFeatureEnabled: utils.Pointer[bool](true),
				ElectronicInvoicingWizardStep:       "",
				ElectronicInvoicingWizardProgress:   "unstarted",
				StatusElectronicInvoicingOnProvider: "",
				UblVersion:                          "2.1",
				FinishedOseConfiguration:            "",
				IsEnabledOnProvider:                 false,
			},
			//InvoicePreferences: datatypes.JSON([]byte(`{"defaultAnotation": null, "defaultTermsAndConditions": ""}`)),
		},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.Company{}.TableName()+" table data initialization failed!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCompany) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("id = ?", 1).
		First(&sysModel.Company{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
