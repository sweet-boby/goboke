package handler

import "github.com/gin-gonic/gin"

// ping handles GET /ping - health check endpoint
func Ping(c *gin.Context) {
	// TODO: Return simple pong response with request ID
	data, _ := c.Get("request_id")
	c.JSON(200, gin.H{
		"success":    true,
		"request_id": data,
	})
}
