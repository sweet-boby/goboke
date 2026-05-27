package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// ContentTypeMiddleware validates content type for POST/PUT requests
func ContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Check content type for POST/PUT requests
		// Must be application/json
		// Return 415 if invalid content type
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			contentType := c.GetHeader("Content-Type")

			if !strings.HasPrefix(contentType, "application/json") {
				c.JSON(415, gin.H{
					"error": "Content-Type must be application/json",
					"code":  "INVALID_CONTENT_TYPE",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
