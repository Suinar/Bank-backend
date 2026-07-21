package credit

import (
	"strconv"

	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	log "github.com/kVinsom/Bank-backend/internal/logging/handler/credit"
	service "github.com/kVinsom/Bank-backend/internal/service/credit"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CreditHandler translates credit HTTP requests into service calls.
type CreditHandler struct {
	service service.ICreditService
}

// NewCreditHandler creates a credit HTTP handler.
func NewCreditHandler(service service.ICreditService) *CreditHandler {
	return &CreditHandler{service: service}
}

// GetAll godoc
// @Summary List credits
// @Tags credits
// @Produce json
// @Success 200 {array} core.Credit
// @Failure 500 {string} string
// @Router /credits [get]
func (h *CreditHandler) GetAll(ctx *gin.Context) {
	const operation = "get_all"
	defer log.RequestStarted(ctx, operation)()
	credits, err := h.service.GetAll(ctx)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, credits)
}

// GetByUser godoc
// @Summary List credits by user
// @Tags credits
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} core.Credit
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /credits/user/{user_id} [get]
func (h *CreditHandler) GetByUser(ctx *gin.Context) {
	const operation = "get_by_user"
	defer log.RequestStarted(ctx, operation)()
	userId := ctx.Param("user_id")

	parsedId, err := common.IdParseHandler(ctx, userId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	credits, err := h.service.GetByUser(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, credits)
}

// GetById godoc
// @Summary Get credit by ID
// @Tags credits
// @Produce json
// @Param id path int true "Credit ID"
// @Success 200 {object} core.Credit
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /credits/{id} [get]
func (h *CreditHandler) GetById(ctx *gin.Context) {
	const operation = "get_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	credits, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, credits)
}

// Create godoc
// @Summary Create credit
// @Tags credits
// @Accept json
// @Produce json
// @Param input body core.CreditCreateInput true "Credit data"
// @Success 200 {object} core.Credit
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /credits [post]
func (h *CreditHandler) Create(ctx *gin.Context) {
	const operation = "create"
	defer log.RequestStarted(ctx, operation)()
	var input core.CreditCreateInput

	if err := ctx.ShouldBind(&input); err != nil {
		log.RequestError(ctx, operation, err)
		common.ResponseBadRequest(ctx, err)
		return
	}

	credit, err := h.service.Create(ctx, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, credit)
}

// RepayById godoc
// @Summary Repay credit
// @Tags credits
// @Param id path int true "Credit ID"
// @Param amount path int true "Repayment amount"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /credits/{id}/repay/{amount} [post]
func (h *CreditHandler) RepayById(ctx *gin.Context) {
	const operation = "repay_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")
	amount := ctx.Param("amount")

	intAmount, err := strconv.Atoi(amount)
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

	err = h.service.Repay(ctx, parsedId, intAmount)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// DeleteById godoc
// @Summary Delete credit
// @Tags credits
// @Param id path int true "Credit ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /credits/{id} [delete]
func (h *CreditHandler) DeleteById(ctx *gin.Context) {
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
