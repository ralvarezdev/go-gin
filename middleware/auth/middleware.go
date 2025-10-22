package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtgin "github.com/ralvarezdev/go-jwt/gin"
	gojwtginctx "github.com/ralvarezdev/go-jwt/gin/context"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"

	gogin "github.com/ralvarezdev/go-gin"
	goginjwtvalidator "github.com/ralvarezdev/go-gin/jwt/validator"
	goginresponse "github.com/ralvarezdev/go-gin/response"
)

type (
	// Middleware struct
	Middleware struct {
		validator                gojwtvalidator.Validator
		responseHandler          goginresponse.Handler
		jwtValidatorErrorHandler goginjwtvalidator.ErrorHandler
	}
)

// NewMiddleware creates a new authentication middleware
//
// Parameters:
//
//   - validator: The JWT validator
//   - responseHandler: The response handler
//   - jwtValidatorErrorHandler: The JWT validator error handler
//
// Returns:
//
//   - *Middleware: The authentication middleware
//   - error: An error if the validator, response handler or validator handler is nil
func NewMiddleware(
	validator gojwtvalidator.Validator,
	responseHandler goginresponse.Handler,
	jwtValidatorErrorHandler goginjwtvalidator.ErrorHandler,
) (*Middleware, error) {
	// Check if either the validator, response handler or validator handler is nil
	if validator == nil {
		return nil, gojwtvalidator.ErrNilValidator
	}
	if responseHandler == nil {
		return nil, goginresponse.ErrNilHandler
	}
	if jwtValidatorErrorHandler == nil {
		return nil, goginjwtvalidator.ErrNilHandler
	}

	return &Middleware{
		validator,
		responseHandler,
		jwtValidatorErrorHandler,
	}, nil
}

// Authenticate return the middleware function that authenticates the request
//
// Parameters:
//
//   - token: The JWT token
//
// Returns:
//
//   - gin.HandlerFunc: The middleware function
func (m Middleware) Authenticate(token gojwttoken.Token) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the authorization from the header
		authorization := ctx.GetHeader(gojwtgin.AuthorizationHeaderKey)

		// Check if the authorization is a bearer token
		parts := strings.Split(authorization, " ")

		// Return an error if the authorization is missing or invalid
		if len(parts) < 2 || parts[0] != gojwt.BearerPrefix {
			m.responseHandler.HandleError(
				ctx,
				goginresponse.NewErrorResponseWithCode(
					gogin.ErrInvalidAuthorizationHeader,
					http.StatusUnauthorized,
				),
			)
			return
		}

		// Get the raw token from the header
		rawToken := parts[1]

		// Validate the token and get the validated claims
		claims, err := m.validator.ValidateClaims(rawToken, token)
		if err != nil {
			m.jwtValidatorErrorHandler(ctx, err)
			return
		}

		// Set the token claims to the context
		gojwtginctx.SetCtxTokenClaims(ctx, claims)
		gojwtginctx.SetCtxToken(ctx, rawToken)

		// Continue
		ctx.Next()
	}
}
