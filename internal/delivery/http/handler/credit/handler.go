package credit

import (
	"strconv"

	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	service "github.com/kVinsom/Bank-backend/internal/service/credit"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
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
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, credits)
}

func (h *CreditHandler) GetByUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	parsedId, err := common.IdParseHandler(ctx, userId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	credits, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, credits)
}

func (h *CreditHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	credits, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, credits)
}

func (h *CreditHandler) Create(ctx *gin.Context) {
	var input core.CreditCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		common.ResponseBadRequest(ctx, err)
		return
	}

	credit, err := h.service.Create(ctx, &input)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, credit)
}

func (h *CreditHandler) RepayById(ctx *gin.Context) {
	id := ctx.Param("id")
	amount := ctx.Param("amount")

	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		common.ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.Repay(ctx, parsedId, intAmount)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

func (h *CreditHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.Delete(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}
