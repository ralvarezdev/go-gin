package redis

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	goratelimiterredis "github.com/ralvarezdev/go-rate-limiter/redis"
)

type (
	// Middleware struct
	Middleware struct {
		rateLimiter goratelimiterredis.RateLimiter
	}
)

// NewMiddleware creates a new rate limiter middleware
//
// Parameters:
//
// rateLimiter goratelimiterredis.RateLimiter: the rate limiter
//
// Returns:
//
// *Middleware: the middleware instance
// error: if the rate limiter is nil
func NewMiddleware(rateLimiter goratelimiterredis.RateLimiter) (
	*Middleware,
	error,
) {
	// Check if the rate limiter is nil
	if rateLimiter == nil {
		return nil, goratelimiterredis.ErrNilRateLimiter
	}

	return &Middleware{
		rateLimiter,
	}, nil
}

// Limit limits the number of requests per IP address
//
// Returns:
//
//	gin.HandlerFunc: the middleware handler function
func (m Middleware) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client IP address
		ip := c.ClientIP()

		// Limit the number of requests per IP address
		if err := m.rateLimiter.Limit(ip); err != nil {
			// Check if the rate limit is exceeded
			if errors.Is(err, goratelimiterredis.ErrTooManyRequests) {
				c.AbortWithStatus(http.StatusTooManyRequests)
				return
			}
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Next()
	}
}
