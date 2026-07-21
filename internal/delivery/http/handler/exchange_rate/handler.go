package exchange_rate

import (
	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	log "github.com/kVinsom/Bank-backend/internal/logging/handler/exchange_rate"
	service "github.com/kVinsom/Bank-backend/internal/service/exhange_rate"
)

// ExchangeRateHandler serves exchange-rate query endpoints.
type ExchangeRateHandler struct {
	service service.IExchangeRateService
}

// NewExchangeRateHandler creates an exchange-rate HTTP handler.
func NewExchangeRateHandler(service service.IExchangeRateService) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		service: service,
	}
}

// GetAllRanking godoc
// @Summary List exchange-rate rankings
// @Tags exchange-rates
// @Produce json
// @Param currency_iso_from path int true "Source numeric ISO code"
// @Success 200 {array} RankingResponse
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /exchange-rates/{currency_iso_from} [get]
func (h *ExchangeRateHandler) GetAllRanking(ctx *gin.Context) {
	const operation = "get_all_ranking"
	defer log.RequestStarted(ctx, operation)()
	currencyIsoFrom := ctx.Param("currency_iso_from")

	parsedIsoFrom, err := common.IsoParseHandler(ctx, currencyIsoFrom)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	ranking, err := h.service.GetAllRanking(ctx, parsedIsoFrom)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, ranking)
}

// GetRelativeRanking godoc
// @Summary Get relative exchange rate
// @Tags exchange-rates
// @Produce json
// @Param currency_iso_from path int true "Source numeric ISO code"
// @Param currency_iso_to path int true "Target numeric ISO code"
// @Success 200 {object} RankingResponse
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /exchange-rates/{currency_iso_from}/{currency_iso_to} [get]
func (h *ExchangeRateHandler) GetRelativeRanking(ctx *gin.Context) {
	const operation = "get_relative_ranking"
	defer log.RequestStarted(ctx, operation)()
	currencyIsoFrom := ctx.Param("currency_iso_from")
	currencyIsoTo := ctx.Param("currency_iso_to")

	parsedIsoFrom, err := common.IsoParseHandler(ctx, currencyIsoFrom)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	parsedIsoTo, err := common.IsoParseHandler(ctx, currencyIsoTo)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	ranking, err := h.service.GetRelativeRanking(ctx, parsedIsoFrom, parsedIsoTo)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, ranking)
}
