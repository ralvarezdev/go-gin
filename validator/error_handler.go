package validator

import (
	"github.com/gin-gonic/gin"
	goginresponse "github.com/ralvarezdev/go-gin/response"
	"net/http"
)

type (
	// ErrorHandler function
	ErrorHandler func(ctx *gin.Context, err error)
)

// DefaultErrorHandler function
func DefaultErrorHandler(ctx *gin.Context, err error) {
	ctx.JSON(
		http.StatusUnauthorized,
		goginresponse.ErrorResponse{Error: err.Error()},
	)
	ctx.Abort()
}
