package common

import (
	"errors"

	"github.com/gin-gonic/gin"

	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
)

// ErrorHandler maps domain and transport errors to a stable HTTP response.
func ErrorHandler(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, coreErrors.NotFound):
		ResponseNotFound(ctx, err)
	case errors.Is(err, coreErrors.BadRequest):
		ResponseBadRequest(ctx, err)
	default:
		ResponseInternalServerError(ctx, err)
	}
}
