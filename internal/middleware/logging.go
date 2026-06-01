package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs all requests with timing information
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Capture start time

		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("[%s] %s %s - %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			duration)
		// TODO: Calculate duration and log request
		// Format: [REQUEST_ID] METHOD PATH STATUS DURATION IP USER_AGENT
	}
}
