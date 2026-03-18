package core

import (
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, core.NotFound):
		ResponseNotFound(ctx, err)
	case errors.Is(err, core.BadRequest):
		ResponseBadRequest(ctx, err)
	default:
		ResponseInternalServerError(ctx, err)
	}
}
