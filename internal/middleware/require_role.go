package middleware

import (
	"goboke/internal/dto"

	"github.com/gin-gonic/gin"
)

// Middleware: Role-based authorization
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Get user role from context (set by authMiddleware)
		role, ok := c.Get("role")
		if !ok {
			c.JSON(400, dto.APIResponse{
				Success: false,
				Message: "requireRole fail",
			})
			c.Abort()
			return
		}
		// TODO: Check if user role is in allowed roles
		for _, item := range roles {
			if item == role {
				c.Next()
				return
			}
		}
		// TODO: Return 403 if not authorized
		c.JSON(403, dto.APIResponse{
			Success: false,
			Message: "not permission user",
		})
		c.Abort()
	}
}
