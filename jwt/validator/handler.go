package validator

import (
	"github.com/gin-gonic/gin"
	goginresponse "github.com/ralvarezdev/go-gin/response"
	"net/http"
)

// ErrorHandler handles the possible JWT validation error
type ErrorHandler func(ctx *gin.Context, err error)

// DefaultErrorHandler function
func DefaultErrorHandler(ctx *gin.Context, err error) {
	ctx.JSON(
		http.StatusUnauthorized,
		goginresponse.NewErrorResponse(err),
	)
	ctx.Abort()
}
