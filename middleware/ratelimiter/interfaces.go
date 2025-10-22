package ratelimiter

import (
	"github.com/gin-gonic/gin"
)

type (
	// RateLimiter interface
	RateLimiter interface {
		Limit() gin.HandlerFunc
	}
)
