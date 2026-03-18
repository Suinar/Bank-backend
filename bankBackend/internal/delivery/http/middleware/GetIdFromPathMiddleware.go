package middleware

import (
	"errors"
	"strconv"

	response "github.com/Suinar/Bank-backend/bankBackend/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

func GetIdFromPathMiddleware(jsonNameId string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("jsonNameId")

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			if errors.Is(err, strconv.ErrSyntax) {
				response.ResponseBadRequest(ctx, err)
				return
			}
			response.ResponseInternalServerError(ctx, err)
			return
		}

		ctx.Set(jsonNameId, id)

		ctx.Next()
	}
}
