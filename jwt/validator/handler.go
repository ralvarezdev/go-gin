package validator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goginresponse "github.com/ralvarezdev/go-gin/response"
)

type (
	// ErrorHandler handles the possible JWT validation error
	ErrorHandler func(ctx *gin.Context, err error)
)

// DefaultErrorHandler function
//
// Parameters:
//
//   - ctx: the gin context
//   - err: the error to be handled
func DefaultErrorHandler(ctx *gin.Context, err error) {
	ctx.JSON(
		http.StatusUnauthorized,
		goginresponse.NewErrorResponse(err),
	)
	ctx.Abort()
}
