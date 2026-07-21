package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
)

// IsoParseHandler parses a numeric ISO currency code from a route parameter.
func IsoParseHandler(_ *gin.Context, iso string) (int, error) {
	parsedISO, err := strconv.Atoi(iso)
	if err != nil {
		return 0, coreErrors.BadRequest
	}

	return parsedISO, nil
}
