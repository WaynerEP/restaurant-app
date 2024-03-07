package initialize

import (
	"context"
	"github.com/WaynerEP/restaurant-app/server/models/contact"
	"github.com/WaynerEP/restaurant-app/server/models/inventory"
	"github.com/WaynerEP/restaurant-app/server/models/menu"
	"github.com/WaynerEP/restaurant-app/server/models/order"
	"github.com/WaynerEP/restaurant-app/server/models/reservation"
	sysModel "github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/service/system"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

const initOrderEnsureTables = system.InitOrderExternal - 1

type ensureTables struct{}

// Auto-run
func init() {
	system.RegisterInit(initOrderEnsureTables, &ensureTables{})
}

func (ensureTables) InitializerName() string {
	return "ensure_tables_created"
}

func (e *ensureTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (e *ensureTables) DataInserted(ctx context.Context) bool {
	return true
}

func (e *ensureTables) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	tables := []interface{}{
		//system
		sysModel.Company{},
		sysModel.SysApi{},
		sysModel.SysUser{},
		sysModel.SysBaseMenu{},
		sysModel.SysAuthority{},
		sysModel.JwtBlacklist{},
		sysModel.SysOperationRecord{},
		sysModel.SysBaseMenuParameter{},
		sysModel.SysBaseMenuBtn{},
		sysModel.SysAuthorityBtn{},
		adapter.CasbinRule{},
		//tables
		reservation.Floor{},
		reservation.Environment{},
		reservation.Table{},
		reservation.FloorEnvironment{},
		reservation.FloorEnvironmentTable{},
		//supply
		inventory.SupplyCategory{},
		inventory.UnitMeasure{},
		inventory.Supply{},
		//menu
		menu.ItemCategory{},
		menu.Item{},
		menu.ItemSupply{},
		menu.Recipe{},
		menu.NutritionalValue{},
		//customers
		contact.Customer{},
		//orders
		order.MenuOrder{},
		order.PaymentMethod{},
		order.PaymentOrder{},
		order.MenuOrderItem{},
	}
	for _, t := range tables {
		_ = db.AutoMigrate(&t)
		// The view 'authority_menu' might be treated as a table during creation, causing conflicts (seems to be fixed in the updated version of gorm)
		// Since AutoMigrate() is mostly error-agnostic, it is explicitly ignored
	}
	return ctx, nil
}

func (e *ensureTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	tables := []interface{}{
		//system
		sysModel.Company{},
		sysModel.SysApi{},
		sysModel.SysUser{},
		sysModel.SysBaseMenu{},
		sysModel.SysAuthority{},
		sysModel.JwtBlacklist{},
		sysModel.SysOperationRecord{},
		sysModel.SysBaseMenuParameter{},
		sysModel.SysBaseMenuBtn{},
		sysModel.SysAuthorityBtn{},
		adapter.CasbinRule{},
		//tables
		reservation.Floor{},
		reservation.Environment{},
		reservation.Table{},
		reservation.FloorEnvironment{},
		reservation.FloorEnvironmentTable{},
		//supply
		inventory.SupplyCategory{},
		inventory.UnitMeasure{},
		inventory.Supply{},
		//menu
		menu.ItemCategory{},
		menu.Item{},
		menu.ItemSupply{},
		menu.Recipe{},
		menu.NutritionalValue{},
		//customers
		contact.Customer{},
		//orders
		order.MenuOrder{},
		order.PaymentMethod{},
		order.PaymentOrder{},
		order.MenuOrderItem{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}
