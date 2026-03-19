package core

import (
	"strconv"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	service "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type CreditHandler struct {
	service service.ICreditService
}

func NewCreditHandler(service service.ICreditService) *CreditHandler {
	return &CreditHandler{service: service}
}

func (h *CreditHandler) GetAll(ctx *gin.Context) {
	credits, err := h.service.GetAll(ctx)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, credits)
}

func (h *CreditHandler) GetByUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	credits, err := h.service.GetByUser(ctx, userId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, credits)
}

func (h *CreditHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	credits, err := h.service.GetById(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, credits)
}

func (h *CreditHandler) Create(ctx *gin.Context) {
	var input core.CreditCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	credit, err := h.service.Create(ctx, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, credit)
}

func (h *CreditHandler) RepayById(ctx *gin.Context) {
	id := ctx.Param("id")
	amount := ctx.Param("amount")

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

func (h *CreditHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.Delete(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, nil)
}
