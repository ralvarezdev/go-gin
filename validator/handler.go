package validator

import (
	"github.com/gin-gonic/gin"
	goginresponse "github.com/ralvarezdev/go-gin/response"
	"net/http"
)

type (
	// Handler interface
	Handler interface {
		HandleError(ctx *gin.Context, err error)
	}

	// DefaultHandler struct
	DefaultHandler struct{}
)

// NewDefaultHandler creates a new default handler
func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

// HandleError handles the error
func (d *DefaultHandler) HandleError(ctx *gin.Context, err error) {
	ctx.JSON(
		http.StatusUnauthorized,
		goginresponse.NewErrorResponse(err),
	)
	ctx.Abort()
}
