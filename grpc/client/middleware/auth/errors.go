package auth

import "errors"

var (
	InvalidAuthorizationHeaderError = errors.New("invalid authorization header")
	NilGRPCInterceptionsError       = errors.New("grpc interceptions cannot be nil")
)
