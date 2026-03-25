package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func IdParseHandler(ctx *gin.Context, id string) (uint64, error) {
	uintId, err := strconv.ParseUint(id, 10, 64)

	return uintId, err
}
