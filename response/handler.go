package response

import (
	"github.com/gin-gonic/gin"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
)

type (
	// DefaultHandler struct
	DefaultHandler struct {
		mode *goflagsmode.Flag
	}
)

// NewDefaultHandler creates a new default request handler
//
// Parameters:
//
//   - mode: the flag mode
//
// Returns:
//
//   - *DefaultHandler: a pointer to the DefaultHandler struct
//   - error: an error if the mode is nil
func NewDefaultHandler(mode *goflagsmode.Flag) (*DefaultHandler, error) {
	// Check if the flag mode is nil
	if mode == nil {
		return nil, goflagsmode.ErrNilModeFlag
	}
	return &DefaultHandler{mode}, nil
}

// HandleSuccess handles the success response
//
// Parameters:
//
//   - ctx: the gin context
//   - response: the response to be sent
func (d DefaultHandler) HandleSuccess(
	ctx *gin.Context,
	response *Response,
) {
	if response != nil && response.code != nil {
		ctx.JSON(*response.code, response.data)
	} else {
		SendInternalServerError(ctx)
	}
}

// HandleErrorProne handles the response that may contain an error
//
// Parameters:
//
//   - ctx: the gin context
//   - successResponse: the success response
//   - errorResponse: the error response
func (d DefaultHandler) HandleErrorProne(
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
//
// Parameters:
//
//   - ctx: the gin context
//   - response: the response to be sent
func (d DefaultHandler) HandleError(
	ctx *gin.Context,
	response *Response,
) {
	if response != nil && response.code != nil {
		ctx.JSON(*response.code, response.data)
		ctx.Abort()
	} else {
		SendInternalServerError(ctx)
	}
}
