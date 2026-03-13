package middleware

import (
	"strconv"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

func GetIdFromPathMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			core.ResponseBadRequest(ctx, err)
			return
		}

		ctx.Set("user_id", id)

		ctx.Next()
	}
}
