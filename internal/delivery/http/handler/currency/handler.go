package currency

import (
	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	log "github.com/kVinsom/Bank-backend/internal/logging/handler/currency"
	service "github.com/kVinsom/Bank-backend/internal/service/currency"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CurrencyHandler translates currency HTTP requests into service calls.
type CurrencyHandler struct {
	currencyService service.ICurrencyService
}

// NewCurrencyHandler creates a currency HTTP handler.
func NewCurrencyHandler(
	currencyService service.ICurrencyService) *CurrencyHandler {
	return &CurrencyHandler{
		currencyService: currencyService,
	}
}

// GetAll godoc
// @Summary List currencies
// @Tags currencies
// @Produce json
// @Success 200 {array} core.Currency
// @Failure 500 {string} string
// @Router /currencies [get]
func (h *CurrencyHandler) GetAll(ctx *gin.Context) {
	const operation = "get_all"
	defer log.RequestStarted(ctx, operation)()
	currencies, err := h.currencyService.GetAll(ctx)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currencies)
}

// GetById godoc
// @Summary Get currency by ID
// @Tags currencies
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} core.Currency
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /currencies/{id} [get]
func (h *CurrencyHandler) GetById(ctx *gin.Context) {
	const operation = "get_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	currency, err := h.currencyService.GetById(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

// GetByIso godoc
// @Summary Get currency by ISO code
// @Tags currencies
// @Produce json
// @Param iso_code path string true "ISO code"
// @Success 200 {object} core.Currency
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /currencies/iso/{iso_code} [get]
func (h *CurrencyHandler) GetByIso(ctx *gin.Context) {
	const operation = "get_by_iso"
	defer log.RequestStarted(ctx, operation)()
	isoCode := ctx.Param("iso_code")

	currency, err := h.currencyService.GetByIso(ctx, isoCode)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

// GetBySymbol godoc
// @Summary Get currency by symbol
// @Tags currencies
// @Produce json
// @Param symbol path string true "Currency symbol"
// @Success 200 {object} core.Currency
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /currencies/symbol/{symbol} [get]
func (h *CurrencyHandler) GetBySymbol(ctx *gin.Context) {
	const operation = "get_by_symbol"
	defer log.RequestStarted(ctx, operation)()
	symbol := ctx.Param("symbol")

	if len([]rune(symbol)) == 0 || len([]rune(symbol)) > 1 {
		common.ResponseBadRequest(ctx, coreErrors.MoreThanOneCharacter)
		return
	}
	parseSymbol := []rune(symbol)[0]

	currency, err := h.currencyService.GetBySymbol(ctx, parseSymbol)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

// Create godoc
// @Summary Create currency
// @Tags currencies
// @Accept json
// @Produce json
// @Param input body core.CurrencyCreateInput true "Currency data"
// @Success 200 {object} core.Currency
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /currencies [post]
func (h *CurrencyHandler) Create(ctx *gin.Context) {
	const operation = "create"
	defer log.RequestStarted(ctx, operation)()
	var input core.CurrencyCreateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.RequestError(ctx, operation, err)
		common.ResponseBadRequest(ctx, err)
		return
	}

	currency, err := h.currencyService.Create(ctx, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

// UpdateById godoc
// @Summary Update currency
// @Tags currencies
// @Accept json
// @Produce json
// @Param id path int true "Currency ID"
// @Param input body core.CurrencyUpdateInput true "Currency changes"
// @Success 200 {object} core.Currency
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /currencies/{id} [put]
func (h *CurrencyHandler) UpdateById(ctx *gin.Context) {
	const operation = "update_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")
	var input core.CurrencyUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.RequestError(ctx, operation, err)
		common.ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	currency, err := h.currencyService.Update(ctx, parsedId, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

// DeleteById godoc
// @Summary Delete currency
// @Tags currencies
// @Param id path int true "Currency ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /currencies/{id} [delete]
func (h *CurrencyHandler) DeleteById(ctx *gin.Context) {
	const operation = "delete_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.currencyService.Delete(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}
