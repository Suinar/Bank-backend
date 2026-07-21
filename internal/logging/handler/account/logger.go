package account

import (
	"github.com/gin-gonic/gin"

	handlerlog "github.com/kVinsom/Bank-backend/internal/logging/handler"
)

// RequestStarted logs an account handler request and returns a completion logger.
func RequestStarted(ctx *gin.Context, operation string) func() {
	return handlerlog.RequestStarted(ctx, "account", operation)
}

// RequestError logs an error produced while handling an account request.
func RequestError(ctx *gin.Context, operation string, err error) {
	handlerlog.RequestError(ctx, "account", operation, err)
}
