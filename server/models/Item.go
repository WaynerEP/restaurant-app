package models

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
	"gorm.io/datatypes"
)

type Product struct {
	common.ModelId
	Name              string             `json:"name" form:"name"  gorm:"size:100; not null; unique;" validate:"required,unique_db,min=3"`
	InternalCode      string             `json:"internalCode" form:"internalCode"  gorm:"size:80; not null; unique;"`
	SunatCode         string             `json:"sunatCode" form:"sunatCode"  gorm:"size:8;unique;"`
	Description       string             `json:"description" form:"description" `
	Reference         string             `json:"reference" form:"reference"  gorm:"size:80;"`
	Type              string             `json:"productType" form:"productType"  gorm:"size:15; not null; default:'P'"`
	CalculationScale  int                `json:"calculationScale" gorm:"not null;default:2"`
	UnitMeasureCode   string             `json:"unitMeasureCode" form:"unitMeasureCode"  gorm:"size:10;not null" validate:"required"`
	UnitMeasure       UnitMeasure        `json:"unitMeasure,omitempty" gorm:"foreignKey:UnitMeasureCode;references:Code" validate:"-"`
	CategoryID        uint               `json:"categoryId" form:"categoryId"  gorm:"not null" validate:"required"`
	Category          Category           `json:"category" form:"-" validate:"-"`
	BasePrice         float64            `json:"basePrice" form:"basePrice"  gorm:"not null" validate:"required"`
	TotalPrice        float64            `json:"totalPrice" form:"totalPrice"  gorm:"not null; CHECK:total_price>0;"  validate:"required,gt=0"`
	InitialCost       float64            `json:"initialCost" form:"initialCost"  gorm:"default:0; comment:'Valor de adquisición de tus productos de venta. para servicios no'"`
	Tax               float64            `json:"tax" gorm:"not null"`
	TaxID             int                `json:"taxId"`
	HasVariants       *bool              `json:"hasVariants" form:"hasVariants"  gorm:"default:0;" `
	FeaturedImageId   *string            `json:"featuredImageId" form:"featuredImageId"`
	Images            datatypes.JSON     `json:"images" form:"images"`
	Status            *bool              `json:"status" form:"status"  gorm:"default:1;"`
	ProductWarehouses []ProductWarehouse `json:"productWarehouses,omitempty" form:"productWarehouses"`
	InvoiceProducts   []InvoiceProduct   `json:"invoiceProducts,omitempty" form:"invoiceProducts"`
	Variants          []Variant          `json:"variants,omitempty" form:"variants"`
	common.ControlBy
	common.ModelTime
}

type Products []Product

type Image struct {
	Value string `json:"value"`
	URL   string `json:"url"`
}

// ProductWarehouse Tabla intermedia con los almacenes
type ProductWarehouse struct {
	ProductID   uint      `json:"productID" form:"productID" gorm:"column:product_id;primaryKey;autoIncrement:false"`
	WarehouseID uint      `json:"warehouseID" form:"warehouseID" gorm:"column:warehouse_id;primaryKey;autoIncrement:false"`
	Warehouse   Warehouse `json:"warehouse" form:"-" validate:"-"`
	StockQty    int       `json:"stockQty" form:"stockQty" `
	MinQty      int       `json:"minQty" form:"minQty"`
	MaxQty      int       `json:"maxQty" form:"maxQty"`
	common.ModelTime
}

type Variant struct {
	ID          uint   `json:"id" form:"id" gorm:"primaryKey;autoIncrement" `
	ProductID   uint   `json:"productId" form:"productId"  gorm:"not null"`
	NameVariant string `json:"nameVariant" form:"nameVariant" gorm:"size:100; not null"`
	Value       string `json:"value" form:"value" gorm:"size:100; not null"`
}
