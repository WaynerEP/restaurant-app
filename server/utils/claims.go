package utils

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	systemReq "github.com/WaynerEP/restaurant-app/server/models/system/request"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"

	"net"
)

func ClearToken(c *gin.Context) {
	// Add cookie x-token to the web source
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	c.SetCookie("x-token", "", -1, "/", host, true, false)
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// Add cookie x-token to the web source
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	c.SetCookie("x-token", token, maxAge, "/", host, true, false)
}

func GetTokenCookie(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

// GetClaims retrieves the claims parsed from jwt in the Gin context.
func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("Failed to retrieve jwt parsing information from Gin context. Please check if the request header contains x-token and if claims has the specified structure.")
	}
	return claims, err
}

// GetUserID retrieves the user ID extracted from the JWT in the Gin context
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetCompanyId retrieves the user ID extracted from the JWT in the Gin context
func GetCompanyId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.CompanyId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.CompanyId
	}
}

// GetUserUuid retrieves the user UUID parsed from jwt in the Gin context.
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId retrieves the user role ID parsed from jwt in the Gin context.
func GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo retrieves the user information parsed from jwt in the Gin context.
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// GetUserName retrieves the username parsed from jwt in the Gin context.
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.Username
	}
}
