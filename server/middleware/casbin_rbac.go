package middleware

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/service"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler is an interceptor
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.GVA_CONFIG.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			// Get the request path
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.GVA_CONFIG.System.RouterPrefix)
			// Get the request method
			act := c.Request.Method
			// Get the user's role
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin() // Check if the policy exists
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "Insufficient permissions", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
