package internal

import (
	"fmt"
	"github.com/WaynerEP/restaurant-app/server/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter is the constructor for the writer.
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf prints formatted log messages.
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GVA_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.GVA_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.GVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
