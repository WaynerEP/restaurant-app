package initialize

import (
	"fmt"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/middleware"
	"github.com/WaynerEP/restaurant-app/server/plugin/email"
	"github.com/WaynerEP/restaurant-app/server/utils/plugin"

	"github.com/gin-gonic/gin"
)

// PluginInit initializes the provided plugins for a given router group.
func PluginInit(group *gin.RouterGroup, plugins ...plugin.Plugin) {
	for i := range plugins {
		pluginGroup := group.Group(plugins[i].RouterPath())
		plugins[i].Register(pluginGroup)
	}
}

// InstallPlugin installs plugins for the specified router.
func InstallPlugin(router *gin.Engine) {
	publicGroup := router.Group("")
	fmt.Println("Installing plugins without authentication ==>", publicGroup)
	privateGroup := router.Group("")
	fmt.Println("Installing authenticated plugins ==>", privateGroup)
	privateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	// Adding plugins linked to roles and permissions. Example: local example mode and online repository mode can be switched using the import above. The effect is the same.
	PluginInit(privateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))
}
