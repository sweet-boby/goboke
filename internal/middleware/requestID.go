package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDMiddleware generates a unique request ID for each request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Generate UUID for request ID
		// Use github.com/google/uuid package
		// Store in context as "request_id"
		// Add to response header as "X-Request-ID"

		// Check if request ID already exists in header
		requestID := c.GetHeader("X-Request-ID")

		if requestID == "" {
			// Generate new UUID
			requestID = uuid.New().String()
		}

		// Store in context for other middleware/handlers
		c.Set("request_id", requestID)

		// Add to response headers
		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}
