package response

import (
	"github.com/gin-gonic/gin"
	gogin "github.com/ralvarezdev/go-gin"
	"net/http"
)

// SendInternalServerError sends an internal server error response
func SendInternalServerError(ctx *gin.Context) {
	ctx.JSON(
		http.StatusInternalServerError,
		NewJSONErrorResponseFromString(gogin.InternalServerError),
	)
	ctx.Abort()
}
