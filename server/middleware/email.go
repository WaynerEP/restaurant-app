package middleware

import (
	"bytes"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/models/system"
	"github.com/WaynerEP/restaurant-app/server/plugin/email/utils"
	"github.com/WaynerEP/restaurant-app/server/service"
	utils2 "github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"strconv"
	"time"
)

var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

func ErrorToEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var username string
		claims, _ := utils2.GetClaims(c)
		if claims.Username != "" {
			username = claims.Username
		} else {
			id, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			user, err := userService.FindUserById(id)
			if err != nil {
				username = "Unknown"
			}
			username = user.Username
		}
		body, _ := io.ReadAll(c.Request.Body)
		// Write back to the request body, ioutil.ReadAll will clear the data in c.Request.Body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
		}
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		status := c.Writer.Status()
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		str := "Received request body: " + record.Body + "\n" + "Request method: " + record.Method + "\n" + "Error message: " + record.ErrorMessage + "\n" + "Time taken: " + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + " called " + record.Path + " and encountered an error"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				global.GVA_LOG.Error("ErrorToEmail Failed, err:", zap.Error(err))
			}
		}
	}
}
