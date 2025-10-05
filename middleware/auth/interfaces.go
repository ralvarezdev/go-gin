package auth

import (
	"github.com/gin-gonic/gin"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
)

type (
	// Authenticator interface
	Authenticator interface {
		Authenticate(token gojwttoken.Token) gin.HandlerFunc
	}
)
