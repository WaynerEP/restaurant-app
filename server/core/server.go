package core

import (
	"fmt"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/initialize"
	"github.com/WaynerEP/restaurant-app/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// Initialize Redis service
		initialize.Redis()
	}
	// Load JWT data from the database
	if global.GVA_DB != nil {
		system.LoadAll()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	// listen and serve on 0.0.0.0:8080
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
