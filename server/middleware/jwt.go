package middleware

import (
	"errors"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/service"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// JWT authentication retrieves the token from the "x-token" header.
		// The front end should store the token in a cookie or local storage, and expiration should be coordinated with the backend.
		token := utils.GetToken(c)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "Not logged in or unauthorized access", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "Your account is logged in from another location or the token is invalid", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// ParseToken decodes the information contained in the token.
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "Authorization has expired", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		// Optionally check if the logged-in user is disabled by the administrator.
		// Logic for user deletion needs optimization (if required, uncomment the code).

		// if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		// 	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		// 	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		// 	c.Abort()
		// }

		// Set claims in the context and proceed to the next middleware.
		c.Set("claims", claims)
		c.Next()

		// Refresh the token if it is close to expiration.
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))

			// Update Redis JWT and blacklist the old token if multipoint is enabled.
			if global.GVA_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else {
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
	}
}
