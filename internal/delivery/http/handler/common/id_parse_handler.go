package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// IdParseHandler parses and validates a positive resource identifier.
func IdParseHandler(_ *gin.Context, id string) (int64, error) {
	intId, err := strconv.ParseInt(id, 10, 64)

	return intId, err
}
