package route

import (
	"github.com/gin-gonic/gin"
	gojwtgrpcauth "github.com/ralvarezdev/go-gin/grpc/client/middleware/auth"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
)

type (
	// Handler interface
	Handler interface {
		NewAuthenticated(route, grpcMethod string, handler gin.HandlerFunc) (
			string,
			gin.HandlerFunc,
			gin.HandlerFunc,
		)
		NewUnauthenticated(route string, handler gin.HandlerFunc) (string, gin.HandlerFunc)
	}

	// DefaultHandler struct
	DefaultHandler struct {
		authentication    gojwtgrpcauth.Authentication
		grpcInterceptions *map[string]gojwtinterception.Interception
	}
)

// NewDefaultHandler creates a new default response handler
func NewDefaultHandler(
	authentication gojwtgrpcauth.Authentication,
	grpcInterceptions *map[string]gojwtinterception.Interception,
) *DefaultHandler {
	return &DefaultHandler{authentication: authentication, grpcInterceptions: grpcInterceptions}
}

// NewAuthenticated creates an authenticated endpoint
func (d *DefaultHandler) NewAuthenticated(route, grpcMethod string, handler gin.HandlerFunc) (
	string,
	gin.HandlerFunc,
	gin.HandlerFunc,
) {
	// Create the endpoint
	return route, d.authentication.Authenticate(
		grpcMethod,
		d.grpcInterceptions,
	), handler
}

// NewUnauthenticated creates an unauthenticated endpoint
func (d *DefaultHandler) NewUnauthenticated(route string, handler gin.HandlerFunc) (
	string,
	gin.HandlerFunc,
) {
	// Create the endpoint
	return route, handler
}
