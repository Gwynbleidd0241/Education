package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		server, _ := os.Hostname()

		fmt.Printf("Server: %s, Method: %s, Path: %s, Status: %d, Duration: %v\n", server, method, path, status, duration)
	}
}
