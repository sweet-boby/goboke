package middleware

import (
	"goboke/internal/dto"
	"goboke/internal/util/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware: JWT Authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, dto.APIResponse{
				Success: false,
				Error:   "Authorization header required",
			})
			c.Abort()
			return
		}

		// TODO: Extract token from "Bearer <token>" format
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// TODO: Validate token using validateToken function
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, dto.APIResponse{
				Success: false,
				Error:   "Invalid token",
			})
			c.Abort()
			return
		}
		// TODO: Set user info in context for route handlers
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
