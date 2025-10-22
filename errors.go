package gogin

import (
	"errors"
	"net/http"
)

var (
	InternalServerError           = http.StatusText(http.StatusInternalServerError)
	ServiceUnavailable            = http.StatusText(http.StatusServiceUnavailable)
	Unauthorized                  = http.StatusText(http.StatusUnauthorized)
	ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
	ErrUnauthenticated            = errors.New("missing or invalid bearer token on authentication header")
	ErrInDevelopment              = errors.New("in development")
)
