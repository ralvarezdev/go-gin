package response

import (
	"github.com/gin-gonic/gin"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
)

type (
	// Handler interface for handling the responses
	Handler interface {
		HandleSuccess(ctx *gin.Context, response *Response)
		HandleErrorProne(
			ctx *gin.Context,
			successResponse *Response,
			errorResponse *Response,
		)
		HandleError(ctx *gin.Context, response *Response)
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

// HandleSuccess handles the success response
func (d *DefaultHandler) HandleSuccess(
	ctx *gin.Context,
	response *Response,
) {
	if response != nil && response.Code != nil {
		ctx.JSON(*response.Code, response.Data)
	} else {
		SendInternalServerError(ctx)
	}
}

// HandleErrorProne handles the response that may contain an error
func (d *DefaultHandler) HandleErrorProne(
	ctx *gin.Context,
	successResponse *Response,
	errorResponse *Response,
) {
	// Check if the error response is nil
	if errorResponse != nil {
		d.HandleError(ctx, errorResponse)
		return
	}

	// Handle the success response
	d.HandleSuccess(ctx, successResponse)
}

// HandleError handles the error response
func (d *DefaultHandler) HandleError(
	ctx *gin.Context,
	response *Response,
) {
	if response != nil && response.Code != nil {
		ctx.JSON(*response.Code, response.Data)
		ctx.Abort()
	} else {
		SendInternalServerError(ctx)
	}
}
