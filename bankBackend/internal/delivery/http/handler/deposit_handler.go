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

	deposits, err := h.service.GetByUser(ctx, userId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, deposits)
}

func (h *DepositHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	deposit, err := h.service.GetById(ctx, id)
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

func (h *DepositHandler) RepayById(ctx *gin.Context) {
	id := ctx.Param("id")
	amount := ctx.PostForm("amount")

	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	err = h.service.Repay(ctx, id, intAmount)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, nil)
}

func (h *DepositHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.Delete(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, nil)
}
