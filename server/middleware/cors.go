package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors creates and returns a CORS middleware configured according to the specifications.
func Cors() gin.HandlerFunc {
	// Custom CORS configuration
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{"*"}                                                                                                                     // Allow any origin (you can adjust this)
	conf.AddAllowHeaders("Authorization", "AccessToken", "X-CSRF-Token", "Token", "X-Token", "X-User-Id")                                                 // Allowed headers
	conf.AddExposeHeaders("Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "New-Token", "New-Expires-At") // Exposed headers
	conf.AllowCredentials = true                                                                                                                          // Allow credentials

	// Create and return the CORS middleware with the configuration
	return cors.New(conf)
}
