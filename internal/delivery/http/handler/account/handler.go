package account

import (
	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	service "github.com/kVinsom/Bank-backend/internal/service/account"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type AccountHandler struct {
	service service.IAccountService
}

func NewAccountHandler(service service.IAccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) GetAll(ctx *gin.Context) {
	accounts, err := h.service.GetAll(ctx)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, accounts)
}

func (h *AccountHandler) GetByUser(ctx *gin.Context) {
	id := ctx.Param("user_id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	accounts, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, accounts)
}

func (h *AccountHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, account)
}

func (h *AccountHandler) Create(ctx *gin.Context) {
	var input core.AccountCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.Create(ctx, &input)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, account)
}

func (h *AccountHandler) BlockingById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.Blocking(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}
}

func (h *AccountHandler) CloseById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.Close(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

func (h *AccountHandler) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var input core.AccountUpdateInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		common.ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.Update(ctx, parsedId, &input)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, account)
}

func (h *AccountHandler) DeleteById(ctx *gin.Context) {
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
