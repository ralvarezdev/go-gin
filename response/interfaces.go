package response

import (
	"github.com/gin-gonic/gin"
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
)
