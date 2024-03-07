package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// LogLayout log layout
type LogLayout struct {
	Time      time.Time
	Metadata  map[string]interface{} // stores custom original data
	Path      string                 // request path
	Query     string                 // query parameters
	Body      string                 // request body data
	IP        string                 // IP address
	UserAgent string                 // user agent
	Error     string                 // error message
	Cost      time.Duration          // time spent processing the request
	Source    string                 // source identifier
}

type Logger struct {
	// Filter user-defined filter
	Filter func(c *gin.Context) bool
	// FilterKeyword keyword filter (key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess authentication process
	AuthProcess func(c *gin.Context, layout *LogLayout)
	// Log processing
	Print func(LogLayout)
	// Source unique identifier for the service
	Source string
}

func (l Logger) SetLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		var body []byte
		if l.Filter != nil && !l.Filter(c) {
			body, _ = c.GetRawData()
			// Put the original body back
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
		cost := time.Since(start)
		layout := LogLayout{
			Time:      time.Now(),
			Path:      path,
			Query:     query,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Error:     strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
			Cost:      cost,
			Source:    l.Source,
		}
		if l.Filter != nil && !l.Filter(c) {
			layout.Body = string(body)
		}
		if l.AuthProcess != nil {
			// Process authentication-related information
			l.AuthProcess(c, &layout)
		}
		if l.FilterKeyword != nil {
			// Perform custom key/value filtering or desensitization
			l.FilterKeyword(&layout)
		}
		// Custom log processing
		l.Print(layout)
	}
}

func DefaultLogger() gin.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			// Standard output, collected by Kubernetes
			v, _ := json.Marshal(layout)
			fmt.Println(string(v))
		},
		Source: "GVA",
	}.SetLoggerMiddleware()
}
