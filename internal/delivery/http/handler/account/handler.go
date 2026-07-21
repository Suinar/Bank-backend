package account

import (
	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	log "github.com/kVinsom/Bank-backend/internal/logging/handler/account"
	service "github.com/kVinsom/Bank-backend/internal/service/account"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// AccountHandler translates account HTTP requests into service calls.
type AccountHandler struct {
	service service.IAccountService
}

// NewAccountHandler creates an account HTTP handler.
func NewAccountHandler(service service.IAccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

// GetAll godoc
// @Summary List accounts
// @Tags accounts
// @Produce json
// @Success 200 {array} core.Account
// @Failure 500 {string} string
// @Router /accounts [get]
func (h *AccountHandler) GetAll(ctx *gin.Context) {
	const operation = "get_all"
	defer log.RequestStarted(ctx, operation)()

	accounts, err := h.service.GetAll(ctx)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, accounts)
}

// GetByUser godoc
// @Summary List accounts by user
// @Tags accounts
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} core.Account
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /accounts/user/{user_id} [get]
func (h *AccountHandler) GetByUser(ctx *gin.Context) {
	const operation = "get_by_user"
	defer log.RequestStarted(ctx, operation)()

	id := ctx.Param("user_id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	accounts, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, accounts)
}

// GetById godoc
// @Summary Get account by ID
// @Tags accounts
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} core.Account
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /accounts/{id} [get]
func (h *AccountHandler) GetById(ctx *gin.Context) {
	const operation = "get_by_id"
	defer log.RequestStarted(ctx, operation)()

	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, account)
}

// Create godoc
// @Summary Create account
// @Tags accounts
// @Accept json
// @Produce json
// @Param input body core.AccountCreateInput true "Account data"
// @Success 200 {object} core.Account
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /accounts [post]
func (h *AccountHandler) Create(ctx *gin.Context) {
	const operation = "create"
	defer log.RequestStarted(ctx, operation)()

	var input core.AccountCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.Create(ctx, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, account)
}

// BlockingById godoc
// @Summary Block account
// @Tags accounts
// @Param id path int true "Account ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /accounts/{id}/block [patch]
func (h *AccountHandler) BlockingById(ctx *gin.Context) {
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
}

// CloseById godoc
// @Summary Close account
// @Tags accounts
// @Param id path int true "Account ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /accounts/{id}/close [patch]
func (h *AccountHandler) CloseById(ctx *gin.Context) {
	const operation = "close_by_id"
	defer log.RequestStarted(ctx, operation)()

	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.Close(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// UpdateById godoc
// @Summary Update account
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param input body core.AccountUpdateInput true "Account changes"
// @Success 200 {object} core.Account
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /accounts/{id} [put]
func (h *AccountHandler) UpdateById(ctx *gin.Context) {
	const operation = "update_by_id"
	defer log.RequestStarted(ctx, operation)()

	id := ctx.Param("id")
	var input core.AccountUpdateInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	account, err := h.service.Update(ctx, parsedId, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, account)
}

// DeleteById godoc
// @Summary Delete account
// @Tags accounts
// @Param id path int true "Account ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /accounts/{id} [delete]
func (h *AccountHandler) DeleteById(ctx *gin.Context) {
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
