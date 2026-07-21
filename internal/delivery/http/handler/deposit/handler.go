package deposit

import (
	"strconv"

	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	log "github.com/kVinsom/Bank-backend/internal/logging/handler/deposit"
	service "github.com/kVinsom/Bank-backend/internal/service/deposit"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// DepositHandler translates deposit HTTP requests into service calls.
type DepositHandler struct {
	service service.IDepositService
}

// NewDepositHandler creates a deposit HTTP handler.
func NewDepositHandler(service service.IDepositService) *DepositHandler {
	return &DepositHandler{service: service}
}

// GetAll godoc
// @Summary List deposits
// @Tags deposits
// @Produce json
// @Success 200 {array} core.Deposit
// @Failure 500 {string} string
// @Router /deposits [get]
func (h *DepositHandler) GetAll(ctx *gin.Context) {
	const operation = "get_all"
	defer log.RequestStarted(ctx, operation)()
	deposits, err := h.service.GetAll(ctx)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, deposits)
}

// GetByUser godoc
// @Summary List deposits by user
// @Tags deposits
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} core.Deposit
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /deposits/user/{user_id} [get]
func (h *DepositHandler) GetByUser(ctx *gin.Context) {
	const operation = "get_by_user"
	defer log.RequestStarted(ctx, operation)()
	userId := ctx.Param("user_id")

	parsedId, err := common.IdParseHandler(ctx, userId)
	if err != nil {
		log.RequestError(ctx, operation, err)
	}

	deposits, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, deposits)
}

// GetById godoc
// @Summary Get deposit by ID
// @Tags deposits
// @Produce json
// @Param id path int true "Deposit ID"
// @Success 200 {object} core.Deposit
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /deposits/{id} [get]
func (h *DepositHandler) GetById(ctx *gin.Context) {
	const operation = "get_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	deposit, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, deposit)
}

// Create godoc
// @Summary Create deposit
// @Tags deposits
// @Accept json
// @Produce json
// @Param input body core.DepositCreateInput true "Deposit data"
// @Success 200 {object} core.Deposit
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /deposits [post]
func (h *DepositHandler) Create(ctx *gin.Context) {
	const operation = "create"
	defer log.RequestStarted(ctx, operation)()
	var input core.DepositCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		log.RequestError(ctx, operation, err)
		common.ResponseBadRequest(ctx, err)
		return
	}

	deposit, err := h.service.Create(ctx, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, deposit)
}

// ReplenishById godoc
// @Summary Replenish deposit
// @Tags deposits
// @Param id path int true "Deposit ID"
// @Param amount path int true "Replenishment amount"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /deposits/{id}/replenish/{amount} [post]
func (h *DepositHandler) ReplenishById(ctx *gin.Context) {
	const operation = "replenish_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")
	amount := ctx.PostForm("amount")

	parsedAmount, err := strconv.Atoi(amount)
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

	err = h.service.Replenish(ctx, parsedId, parsedAmount)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// DeleteById godoc
// @Summary Delete deposit
// @Tags deposits
// @Param id path int true "Deposit ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /deposits/{id} [delete]
func (h *DepositHandler) DeleteById(ctx *gin.Context) {
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
