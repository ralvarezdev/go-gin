package auth

import (
	"github.com/gin-gonic/gin"
	gogin "github.com/ralvarezdev/go-gin"
	gogingrpcclientresponse "github.com/ralvarezdev/go-gin/grpc/client/response"
	gogintypes "github.com/ralvarezdev/go-gin/types"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtgin "github.com/ralvarezdev/go-jwt/gin"
	gojwtginctx "github.com/ralvarezdev/go-jwt/gin/context"
	gojwtinterception "github.com/ralvarezdev/go-jwt/token/interception"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	gologger "github.com/ralvarezdev/go-logger"
	"google.golang.org/grpc/status"
	"strings"
)

// Middleware struct
type Middleware struct {
	validator       gojwtvalidator.Validator
	logger          *Logger
	responseHandler gogingrpcclientresponse.Handler
}

// NewMiddleware creates a new authentication middleware
func NewMiddleware(
	validator gojwtvalidator.Validator,
	logger *Logger,
	responseHandler gogingrpcclientresponse.Handler,
) (*Middleware, error) {
	// Check if either the validator, logger, or response handler is nil
	if validator == nil {
		return nil, gojwtvalidator.NilValidatorError
	}
	if logger == nil {
		return nil, gologger.NilLoggerError
	}
	if responseHandler == nil {
		return nil, gogingrpcclientresponse.NilHandlerError
	}

	return &Middleware{
		validator:       validator,
		logger:          logger,
		responseHandler: responseHandler,
	}, nil
}

// Authenticate return the middleware function that authenticates the request
func (m *Middleware) Authenticate(
	grpcMethod string, grpcInterceptions *map[string]gojwtinterception.Interception,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if the gRPC interceptions is nil
		if grpcInterceptions == nil {
			if grpcInterceptions == nil {
				m.logger.MissingGRPCInterceptions()
			}
			ctx.JSON(500, gogintypes.NewErrorResponse(gogin.InternalServerError))
			return
		}

		// Get the request URI and method
		requestURI := ctx.Request.RequestURI

		// Get the gRPC method interception
		interception, ok := (*grpcInterceptions)[grpcMethod]
		if !ok {
			m.logger.MissingGRPCMethod(requestURI)
			ctx.JSON(500, gogintypes.NewErrorResponse(gogin.InternalServerError))
			ctx.Abort()
			return
		}

		// Check if there is None interception
		if interception == gojwtinterception.None {
			ctx.Next()
			return
		}

		// Get the authorization from the header
		authorization := ctx.GetHeader(gojwtgin.AuthorizationHeaderKey)

		// Check if the authorization is a bearer token
		parts := strings.Split(authorization, " ")

		// Return an error if the authorization is missing or invalid
		if len(parts) < 2 || parts[0] != gojwt.BearerPrefix {
			ctx.JSON(
				401, gin.H{"error": InvalidAuthorizationHeaderError.Error()},
			)
			ctx.Abort()
			return
		}

		// Get the raw token from the header
		rawToken := parts[1]

		// Validate the token and get the validated claims
		claims, err := m.validator.GetValidatedClaims(rawToken, interception)
		if err != nil {
			// Check if the error is a gRPC status error
			if _, ok := status.FromError(err); ok {
				m.responseHandler.HandleErrorResponse(ctx, err)
			} else {
				ctx.JSON(401, gin.H{"error": err.Error()})
			}

			ctx.Abort()
			return
		}

		// Set the raw token and token claims to the context
		gojwtginctx.SetCtxRawToken(ctx, &rawToken)
		gojwtginctx.SetCtxTokenClaims(ctx, claims)

		// Continue
		ctx.Next()
	}
}
