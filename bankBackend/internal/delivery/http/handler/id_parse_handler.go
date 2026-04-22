package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func IdParseHandler(ctx *gin.Context, id string) (int64, error) {
	intId, err := strconv.ParseInt(id, 10, 64)

	return intId, err
}
