package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	service "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
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
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, accounts)
}

func (h *AccountHandler) GetByUser(ctx *gin.Context) {
	id := ctx.Param("user_id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	accounts, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, accounts)
}

func (h *AccountHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, account)
}

func (h *AccountHandler) Create(ctx *gin.Context) {
	var input core.AccountCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.Create(ctx, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, account)
}

func (h *AccountHandler) BlockingById(ctx *gin.Context) {
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
}

func (h *AccountHandler) CloseById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	err = h.service.Close(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, nil)
}

func (h *AccountHandler) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var input core.AccountUpdateInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.Update(ctx, parsedId, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, account)
}

func (h *AccountHandler) DeleteById(ctx *gin.Context) {
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
