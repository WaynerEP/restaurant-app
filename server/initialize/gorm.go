package initialize

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/contact"
	"github.com/WaynerEP/restaurant-app/server/models/inventory"
	"github.com/WaynerEP/restaurant-app/server/models/menu"
	"github.com/WaynerEP/restaurant-app/server/models/order"
	"github.com/WaynerEP/restaurant-app/server/models/reservation"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables performs automatic table migration on the global database.
func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		//system
		system.Company{},
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysOperationRecord{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
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
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
