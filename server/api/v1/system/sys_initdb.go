package system

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/system/request"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type DBApi struct{}

// InitDB .
func (i *DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("Database configuration already exists!")
		response.FailWithMessage("Database configuration already exists", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("Parameter validation failed!", zap.Error(err))
		response.FailWithMessage("Parameter validation failed", c)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error("Failed to automatically create the database!", zap.Error(err))
		response.FailWithMessage("Failed to automatically create the database. Please check the backend logs and perform initialization after verification", c)
		return
	}
	response.OkWithMessage("Automatically created the database successfully", c)
}

// CheckDB .
func (i *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "Go to initialize the database"
		needInit = true
	)

	if global.GVA_DB != nil {
		message = "The database does not need initialization"
		needInit = false
	}
	global.GVA_LOG.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, c)
}
