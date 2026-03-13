package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

func ResponseBadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
}

func ResponseInternalServerError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
}

func ResponseNotFound(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
}
