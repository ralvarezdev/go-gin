package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gogin "github.com/ralvarezdev/go-gin"
)

// SendInternalServerError sends an internal server error response
//
// Parameters:
//
//   - ctx: the gin context
func SendInternalServerError(ctx *gin.Context) {
	ctx.JSON(
		http.StatusInternalServerError,
		NewJSONErrorResponseFromString(gogin.InternalServerError),
	)
	ctx.Abort()
}
