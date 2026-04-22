package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	service "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	service service.ICardService
}

func NewCardHandler(service service.ICardService) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) GetAll(ctx *gin.Context) {
	cards, err := h.service.GetAll(ctx)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, cards)
}

func (h *CardHandler) GetByUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	parsedId, err := IdParseHandler(ctx, userId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	cards, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, cards)
}

func (h *CardHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	card, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, card)
}

func (h *CardHandler) GetByNumber(ctx *gin.Context) {
	number := ctx.Param("number")

	card, err := h.service.GetByNumber(ctx, number)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, card)
}

func (h *CardHandler) Create(ctx *gin.Context) {
	var input core.CardCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	card, err := h.service.Create(ctx, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, card)
}

func (h *CardHandler) BlockingById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	err = h.service.Blocking(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, nil)
}

func (h *CardHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	err = h.service.Delete(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, nil)
}
