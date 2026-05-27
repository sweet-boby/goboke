package middleware

import (
	"goboke/internal/dto"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates API keys for protected routes
func AuthMiddleware() gin.HandlerFunc {
	// TODO: Define valid API keys and their roles
	// "admin-key-123" -> "admin"
	// "user-key-456" -> "user"
	keys := map[string]string{
		"admin-key-123": "admin",
		"user-key-456":  "user",
	}

	return func(c *gin.Context) {
		// TODO: Get API key from X-API-Key header
		// TODO: Validate API key
		// TODO: Set user role in context
		// TODO: Return 401 if invalid or missing
		apiKey := c.GetHeader("X-API-Key")

		if apiKey == "" {
			c.JSON(401, dto.APIResponse{
				Success: false,
				Error:   "API key required",
			})
			c.Abort() // Stop middleware chain
			return
		}
		key, ok := keys[apiKey]
		// Validate API key
		if ok == false {
			c.JSON(401, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		// Store user info in context
		c.Set("user_role", key)

		c.Next() // Continue to next middleware/handler
	}
}
