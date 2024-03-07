package plugin

import (
	"github.com/gin-gonic/gin"
)

const (
	OnlyFuncName = "Plugin"
)

// Plugin is an interface for plugin mode
type Plugin interface {
	// Register registers routes
	Register(group *gin.RouterGroup)

	// RouterPath returns the registered route
	RouterPath() string
}
