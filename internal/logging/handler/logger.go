package handler

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestStarted(ctx *gin.Context, handlerName, operation string) func() {
	startedAt := time.Now()
	log.Printf(
		"%s handler: request started operation=%s method=%s path=%s client_ip=%s",
		handlerName, operation, ctx.Request.Method, ctx.Request.URL.Path, ctx.ClientIP(),
	)

	return func() {
		log.Printf(
			"%s handler: request completed operation=%s method=%s path=%s status=%d duration=%s",
			handlerName, operation, ctx.Request.Method, ctx.Request.URL.Path,
			ctx.Writer.Status(), time.Since(startedAt),
		)
	}
}

func RequestError(ctx *gin.Context, handlerName, operation string, err error) {
	log.Printf(
		"%s handler: request failed operation=%s method=%s path=%s error=%q",
		handlerName, operation, ctx.Request.Method, ctx.Request.URL.Path, err,
	)
}
