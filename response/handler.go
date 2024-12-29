package response

import (
	"github.com/gin-gonic/gin"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
)

type (
	// Handler interface for handling the responses
	Handler interface {
		HandleResponse(
			ctx *gin.Context,
			successCode int,
			response interface{},
			errorCode int,
			err error,
		)
		HandleErrorResponse(ctx *gin.Context, errorCode error, err error)
	}

	// DefaultHandler struct
	DefaultHandler struct {
		mode *goflagsmode.Flag
	}
)

// NewDefaultHandler creates a new default request handler
func NewDefaultHandler(mode *goflagsmode.Flag) (*DefaultHandler, error) {
	// Check if the flag mode is nil
	if mode == nil {
		return nil, goflagsmode.ErrNilModeFlag
	}
	return &DefaultHandler{mode: mode}, nil
}

// HandleResponse handles the response
func (d *DefaultHandler) HandleResponse(
	ctx *gin.Context,
	successCode int,
	response interface{},
	errorCode int,
	err error,
) {
	// Check if the error is nil
	if err == nil {
		ctx.JSON(successCode, response)
		return
	}

	// Handle the error response
	d.HandleErrorResponse(ctx, errorCode, err)
}

// HandleErrorResponse handles the error response
func (d *DefaultHandler) HandleErrorResponse(
	ctx *gin.Context,
	errorCode int,
	err error,
) {
	ctx.JSON(errorCode, NewErrorResponse(err))
}
