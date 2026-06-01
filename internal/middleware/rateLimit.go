package middleware

import (
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var rateLimiters = make(map[string]*rate.Limiter)
var mu sync.Mutex

// RateLimitMiddleware implements rate limiting per IP
func RateLimitMiddleware() gin.HandlerFunc {
	// TODO: Implement rate limiting
	// Limit: 100 requests per IP per minute
	// Use golang.org/x/time/rate package
	// Set headers: X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset
	// Return 429 if rate limit exceeded
	requestsPerSecond := 100
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		mu.Lock()
		limiter, exists := rateLimiters[clientIP]
		if !exists {
			limiter = rate.NewLimiter(rate.Limit(requestsPerSecond), requestsPerSecond*2)
			rateLimiters[clientIP] = limiter
		}
		mu.Unlock()

		remaining := int(limiter.Tokens())
		reset := time.Now().Add(time.Minute).Unix()

		c.Header("X-RateLimit-Limit", strconv.Itoa(requestsPerSecond))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(reset, 10))
		if !limiter.Allow() {

			c.JSON(429, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}
