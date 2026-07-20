package currency

import (
	"strconv"

	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	service "github.com/kVinsom/Bank-backend/internal/service/currency"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type CurrencyHandler struct {
	currencyService     service.ICurrencyService
	convertService      service.IConvertService
	exchangeRateService service.IExchangeRateService
}

func NewCurrencyHandler(currencyService service.ICurrencyService,
	convertService service.IConvertService,
	exchangeRateService service.IExchangeRateService) *CurrencyHandler {
	return &CurrencyHandler{currencyService: currencyService,
		convertService:      convertService,
		exchangeRateService: exchangeRateService,
	}
}

func (h *CurrencyHandler) GetAll(ctx *gin.Context) {
	currencies, err := h.currencyService.GetAll(ctx)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currencies)
}

func (h *CurrencyHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	currency, err := h.currencyService.GetById(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) GetByIso(ctx *gin.Context) {
	isoCode := ctx.Param("iso_code")

	currency, err := h.currencyService.GetByIso(ctx, isoCode)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) GetBySymbol(ctx *gin.Context) {
	symbol := ctx.Param("symbol")

	if len(symbol) == 0 && len(symbol) > 1 {
		common.ResponseBadRequest(ctx, coreErrors.MoreThanOneCharacter)
		return
	}
	parseSymbol := []rune(symbol)[0]

	currency, err := h.currencyService.GetBySymbol(ctx, parseSymbol)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) Convert(ctx *gin.Context) {
	currencyIdFrom := ctx.Param("currency_id_from")
	currencyIdTo := ctx.Param("currency_id_to")

	amount := ctx.Param("amount")

	parseAmount, err := strconv.Atoi(amount)
	if err != nil {
		common.ResponseBadRequest(ctx, err)
		return
	}

	parsedIdFrom, err := common.IdParseHandler(ctx, currencyIdFrom)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	parsedIdTo, err := common.IdParseHandler(ctx, currencyIdTo)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	ranking, err := h.convertService.Convert(ctx, parsedIdFrom, parseAmount, parsedIdTo)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, ranking)
}

func (h *CurrencyHandler) Create(ctx *gin.Context) {
	var input core.CurrencyCreateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.ResponseBadRequest(ctx, err)
		return
	}

	currency, err := h.currencyService.Create(ctx, &input)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var input core.CurrencyUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	currency, err := h.currencyService.Update(ctx, parsedId, &input)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.currencyService.Delete(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

func (h *CurrencyHandler) GetAllRanking(ctx *gin.Context) {
	currencyIdFrom := ctx.Param("currency_id_from")

	parsedIdFrom, err := common.IdParseHandler(ctx, currencyIdFrom)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	ranking, err := h.exchangeRateService.GetAllRanking(ctx, parsedIdFrom)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, ranking)
}

func (h *CurrencyHandler) GetRelativeRanking(ctx *gin.Context) {
	currencyIdFrom := ctx.Param("currency_id_from")
	currencyIdTo := ctx.Param("currency_id_to")

	parsedIdFrom, err := common.IdParseHandler(ctx, currencyIdFrom)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	parsedIdTo, err := common.IdParseHandler(ctx, currencyIdTo)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	ranking, err := h.exchangeRateService.GetRelativeRanking(ctx, parsedIdFrom, parsedIdTo)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, ranking)
}
