package core

import (
	"strconv"

	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
	service "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type DepositHandler struct {
	service service.IDepositService
}

func NewDepositHandler(service service.IDepositService) *DepositHandler {
	return &DepositHandler{service: service}
}

func (h *DepositHandler) GetAll(ctx *gin.Context) {
	deposits, err := h.service.GetAll(ctx)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, deposits)
}

func (h *DepositHandler) GetByUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	parsedId, err := IdParseHandler(ctx, userId)

	deposits, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, deposits)
}

func (h *DepositHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	deposit, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, deposit)
}

func (h *DepositHandler) Create(ctx *gin.Context) {
	var input core.DepositCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	deposit, err := h.service.Create(ctx, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, deposit)
}

func (h *DepositHandler) ReplenishById(ctx *gin.Context) {
	id := ctx.Param("id")
	amount := ctx.PostForm("amount")

	parsedAmount, err := strconv.Atoi(amount)
	if err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	err = h.service.Replenish(ctx, parsedId, parsedAmount)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, nil)
}

func (h *DepositHandler) DeleteById(ctx *gin.Context) {
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
