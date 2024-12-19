package auth

import (
	"github.com/gin-gonic/gin"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
)

// Authentication interface
type Authentication interface {
	Authenticate(
		grpcMethod string,
		grpcInterceptions *map[string]gojwtinterception.Interception,
	) gin.HandlerFunc
}
