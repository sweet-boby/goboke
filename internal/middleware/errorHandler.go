package middleware

import (
	"fmt"
	"goboke/internal/dto"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware handles panics and errors
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// TODO: Handle panics gracefully
		// Return consistent error response format
		// Include request ID in response
		var apiErr dto.APIError

		switch err := recovered.(type) {
		case dto.APIError:
			apiErr = err
		case error:
			apiErr = dto.APIError{
				StatusCode: 500,
				Code:       "INTERNAL_ERROR",
				Message:    "Internal server error",
				Details:    err.Error(),
			}
		default:
			apiErr = dto.APIError{
				StatusCode: 500,
				Code:       "PANIC",
				Message:    "Internal server error",
				Details:    fmt.Sprintf("%v", recovered),
			}
		}

		c.JSON(apiErr.StatusCode, gin.H{
			"success":    false,
			"error":      apiErr.Message,
			"code":       apiErr.Code,
			"message":    apiErr.Details,
			"request_id": c.GetString("request_id"),
		})
	})
}
