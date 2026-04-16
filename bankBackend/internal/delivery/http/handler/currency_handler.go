package core

import (
	"strconv"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	service "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
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
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, currencies)
}

func (h *CurrencyHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	currency, err := h.currencyService.GetById(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) GetByIso(ctx *gin.Context) {
	isoCode := ctx.Param("iso_code")

	currency, err := h.currencyService.GetByIso(ctx, isoCode)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) GetBySymbol(ctx *gin.Context) {
	symbol := ctx.Param("symbol")

	if len(symbol) == 0 && len(symbol) > 1 {
		ResponseBadRequest(ctx, core.MoreThanOneCharacter)
		return
	}
	parseSymbol := []rune(symbol)[0]

	currency, err := h.currencyService.GetBySymbol(ctx, parseSymbol)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) Convert(ctx *gin.Context) {
	currencyIdFrom := ctx.Param("currency_id_from")
	currencyIdTo := ctx.Param("currency_id_to")

	amount := ctx.Param("amount")

	parseAmount, err := strconv.Atoi(amount)
	if err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	parsedIdFrom, err := IdParseHandler(ctx, currencyIdFrom)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	parsedIdTo, err := IdParseHandler(ctx, currencyIdTo)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ranking, err := h.convertService.Convert(ctx, parsedIdFrom, parseAmount, parsedIdTo)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, ranking)
}

func (h *CurrencyHandler) Create(ctx *gin.Context) {
	var input core.CurrencyCreateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	currency, err := h.currencyService.Create(ctx, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var input core.CurrencyUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	currency, err := h.currencyService.Update(ctx, parsedId, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, currency)
}

func (h *CurrencyHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	err = h.currencyService.Delete(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, nil)
}

func (h *CurrencyHandler) GetAllRanking(ctx *gin.Context) {
	currencyIdFrom := ctx.Param("currency_id_from")

	parsedIdFrom, err := IdParseHandler(ctx, currencyIdFrom)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ranking, err := h.exchangeRateService.GetAllRanking(ctx, parsedIdFrom)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, ranking)
}

func (h *CurrencyHandler) GetRelativeRanking(ctx *gin.Context) {
	currencyIdFrom := ctx.Param("currency_id_from")
	currencyIdTo := ctx.Param("currency_id_to")

	parsedIdFrom, err := IdParseHandler(ctx, currencyIdFrom)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	parsedIdTo, err := IdParseHandler(ctx, currencyIdTo)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ranking, err := h.exchangeRateService.GetRelativeRanking(ctx, parsedIdFrom, parsedIdTo)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, ranking)
}
