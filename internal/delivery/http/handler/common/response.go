package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseSuccess writes a successful JSON response.
func ResponseSuccess(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

// ResponseBadRequest writes a client-error response.
func ResponseBadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
}

// ResponseInternalServerError writes an unexpected server-error response.
func ResponseInternalServerError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
}

// ResponseNotFound writes a missing-resource response.
func ResponseNotFound(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
}
