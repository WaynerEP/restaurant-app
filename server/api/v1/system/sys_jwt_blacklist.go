package system

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct{}

// JsonInBlacklist .
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := utils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlacklist(jwt)
	if err != nil {
		global.GVA_LOG.Error("Failed to invalidate JWT!", zap.Error(err))
		response.FailWithMessage("Failed to invalidate JWT", c)
		return
	}
	utils.ClearToken(c)
	response.OkWithMessage("JWT invalidated successfully", c)
}
