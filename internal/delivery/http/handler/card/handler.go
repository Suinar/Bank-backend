package card

import (
	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	log "github.com/kVinsom/Bank-backend/internal/logging/handler/card"
	service "github.com/kVinsom/Bank-backend/internal/service/card"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CardHandler translates card HTTP requests into service calls.
type CardHandler struct {
	service service.ICardService
}

// NewCardHandler creates a card HTTP handler.
func NewCardHandler(service service.ICardService) *CardHandler {
	return &CardHandler{service: service}
}

// GetAll godoc
// @Summary List cards
// @Tags cards
// @Produce json
// @Success 200 {array} core.Card
// @Failure 500 {string} string
// @Router /cards [get]
func (h *CardHandler) GetAll(ctx *gin.Context) {
	const operation = "get_all"
	defer log.RequestStarted(ctx, operation)()
	cards, err := h.service.GetAll(ctx)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, cards)
}

// GetByUser godoc
// @Summary List cards by user
// @Tags cards
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} core.Card
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /cards/user/{user_id} [get]
func (h *CardHandler) GetByUser(ctx *gin.Context) {
	const operation = "get_by_user"
	defer log.RequestStarted(ctx, operation)()
	userId := ctx.Param("user_id")

	parsedId, err := common.IdParseHandler(ctx, userId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	cards, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, cards)
}

// GetById godoc
// @Summary Get card by ID
// @Tags cards
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} core.Card
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /cards/{id} [get]
func (h *CardHandler) GetById(ctx *gin.Context) {
	const operation = "get_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	card, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, card)
}

// GetByNumber godoc
// @Summary Get card by number
// @Tags cards
// @Produce json
// @Param number path string true "Card number"
// @Success 200 {object} core.Card
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /cards/number/{number} [get]
func (h *CardHandler) GetByNumber(ctx *gin.Context) {
	const operation = "get_by_number"
	defer log.RequestStarted(ctx, operation)()
	number := ctx.Param("number")

	card, err := h.service.GetByNumber(ctx, number)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, card)
}

// Create godoc
// @Summary Create card
// @Tags cards
// @Accept json
// @Produce json
// @Param input body core.CardCreateInput true "Card data"
// @Success 200 {object} core.Card
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /cards [post]
func (h *CardHandler) Create(ctx *gin.Context) {
	const operation = "create"
	defer log.RequestStarted(ctx, operation)()
	var input core.CardCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		log.RequestError(ctx, operation, err)
		common.ResponseBadRequest(ctx, err)
		return
	}

	card, err := h.service.Create(ctx, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, card)
}

// BlockingById godoc
// @Summary Block card
// @Tags cards
// @Param id path int true "Card ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /cards/{id}/block [patch]
func (h *CardHandler) BlockingById(ctx *gin.Context) {
	const operation = "block_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.Blocking(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// DeleteById godoc
// @Summary Delete card
// @Tags cards
// @Param id path int true "Card ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /cards/{id} [delete]
func (h *CardHandler) DeleteById(ctx *gin.Context) {
	const operation = "delete_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.Delete(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}
