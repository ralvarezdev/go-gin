package auth

import (
	"github.com/gin-gonic/gin"
	gogin "github.com/ralvarezdev/go-gin"
	goginjwtvalidator "github.com/ralvarezdev/go-gin/jwt/validator"
	goginresponse "github.com/ralvarezdev/go-gin/response"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtgin "github.com/ralvarezdev/go-jwt/gin"
	gojwtginctx "github.com/ralvarezdev/go-jwt/gin/context"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	"net/http"
	"strings"
)

// Middleware struct
type Middleware struct {
	validator                gojwtvalidator.Validator
	responseHandler          goginresponse.Handler
	jwtValidatorErrorHandler goginjwtvalidator.ErrorHandler
}

// NewMiddleware creates a new authentication middleware
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
		validator:                validator,
		responseHandler:          responseHandler,
		jwtValidatorErrorHandler: jwtValidatorErrorHandler,
	}, nil
}

// Authenticate return the middleware function that authenticates the request
func (m *Middleware) Authenticate(interception gojwtinterception.Interception) gin.HandlerFunc {
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
		claims, err := m.validator.GetValidatedClaims(rawToken, interception)
		if err != nil {
			m.jwtValidatorErrorHandler(ctx, err)
			return
		}

		// Set the raw token and token claims to the context
		gojwtginctx.SetCtxRawToken(ctx, &rawToken)
		gojwtginctx.SetCtxTokenClaims(ctx, claims)

		// Continue
		ctx.Next()
	}
}
