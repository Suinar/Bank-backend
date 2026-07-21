package currency

import (
	"github.com/gin-gonic/gin"
	handlerlog "github.com/kVinsom/Bank-backend/internal/logging/handler"
)

func RequestStarted(ctx *gin.Context, operation string) func() {
	return handlerlog.RequestStarted(ctx, "currency", operation)
}
func RequestError(ctx *gin.Context, operation string, err error) {
	handlerlog.RequestError(ctx, "currency", operation, err)
}
