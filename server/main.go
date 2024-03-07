package main

import (
	"database/sql"
	"github.com/WaynerEP/restaurant-app/server/core"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// Initialize Viper configuration - init config
	global.GVA_VP = core.Viper()

	//initialize other settings
	initialize.OtherInit()

	// Initialize the Zap logging library
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)

	//initialize gorm DB
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // Initialize tables
		// Close the database connection before the program ends
		db, _ := global.GVA_DB.DB()
		defer func(db *sql.DB) {
			_ = db.Close()
		}(db)
	}
	// Initialize Validator
	initialize.SetupValidator()

	//run server(routers, cors, etc)
	core.RunWindowsServer()
}
