package system

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"gorm.io/datatypes"
)

type Company struct {
	ID                      uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Logo                    string         `json:"logo"`
	RUC                     string         `json:"ruc" gorm:"size:11; not null; unique " validate:"omitempty,len=11,numeric"`
	CompanyName             string         `json:"companyName" gorm:"size:150; not null;"` // razón social
	TradeName               string         `json:"tradeName" gorm:"size:255;"`
	PhoneCompany            string         `json:"phoneCompany" gorm:"size:25;not null"`
	EmailCompany            string         `json:"emailCompany" gorm:"size:255;not null; unique" validate:"omitempty,email"`
	FiscalAddress           string         `json:"fiscalAddress" gorm:"size: 255"`
	RegimenID               uint32         `json:"regimenID"`
	District                string         `json:"district" gorm:"size:100"`
	Province                string         `json:"province" gorm:"size:100"`
	Department              string         `json:"department" gorm:"size:100"`
	Ubigeo                  string         `json:"ubigeo" gorm:"size:6"`
	Urbanization            string         `json:"urbanization"`
	Website                 string         `json:"website"`
	EconomicActivity        string         `json:"economicActivity" gorm:"size:200; not null"`
	ActivitySecondaryOne    string         `json:"activitySecondaryOne" gorm:"size:200"`
	ActivitySecondaryTwo    string         `json:"activitySecondaryTwo" gorm:"size:200"`
	Banks                   datatypes.JSON `json:"banks" gorm:"type:json"`
	Status                  string         `json:"status"`
	Regime                  string         `json:"regime" gorm:"size:40"`
	ApplicationVersion      string         `json:"applicationVersion" gorm:"size:25;default:'peru'"`
	Timezone                string         `json:"timezone" gorm:"size:30"`
	Profile                 string         `json:"profile" gorm:"size:20"`
	DecimalPrecision        int            `json:"decimalPrecision" gorm:"not null;default:2"`
	CalculationScale        string         `json:"calculationScale" gorm:"not null;default:2"`
	MultiTax                *bool          `json:"multiTax" gorm:"default:1"`
	EmployeesNumber         uint           `json:"employeesNumber" gorm:"check:employees_number>0"`
	Sector                  string         `json:"sector" gorm:"size:50"`
	ShowInvoiceTotalInWords string         `json:"showInvoiceTotalInWords" gorm:"size:3"`
	ShowRetentionInvoice    string         `json:"showRetentionInvoice" gorm:"size:3"`
	ShowNewLineCharOnPdf    string         `json:"showNewLineCharOnPdf" gorm:"size:50"`
	ShowItemReferenceOnPdf  bool           `json:"showItemReferenceOnPdf" gorm:"default:0"`
	Coupon                  string         `json:"coupon" gorm:"size:50"`
	CurrencyCode            string         `json:"currencyCode" gorm:"size:10; not null" validate:"required"`
	MultiCurrency           bool           `json:"multiCurrency" gorm:"default:1"`
	DecimalSeparator        string         `json:"decimalSeparator" gorm:"size:1;default:','"`
	CompanySettingID        uint           `json:"companySettingId"`
	CompanySetting          CompanySetting `json:"companySettings"`
	InvoicePreferences      datatypes.JSON `json:"invoicePreferences" gorm:"type:json"`
	common.ModelTime
	common.ControlBy
}

type CompanySetting struct {
	common.ModelId
	CanStampInvoices                    *bool  `json:"canStampInvoices"`
	ElectronicInvoicing                 *bool  `json:"electronicInvoicing"`
	IsTypesAndNumerationsFeatureEnabled *bool  `json:"isTypesAndNumerationsFeatureEnabled" gorm:"default:1"`
	ElectronicInvoicingWizardStep       string `json:"electronicInvoicingWizardStep" gorm:"size:10"`
	ElectronicInvoicingWizardProgress   string `json:"electronicInvoicingWizardProgress" gorm:"size:10"`
	StatusElectronicInvoicingOnProvider string `json:"statusElectronicInvoicingOnProvider" gorm:"size:10"`
	UblVersion                          string `json:"ublVersion" gorm:"size:5;default:'2.1'"`
	FinishedOseConfiguration            string `json:"finishedOseConfiguration" gorm:"size:10"`
	IsEnabledOnProvider                 bool   `json:"isEnabledOnProvider"`
}

func (Company) TableName() string {
	return "companies"
}

type InvoicePreferences struct {
	DefaultAnnotation         string `json:"defaultAnnotation"`
	DefaultTermsAndConditions string `json:"defaultTermsAndConditions"`
}
