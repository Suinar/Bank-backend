package user

import (
	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	log "github.com/kVinsom/Bank-backend/internal/logging/handler/user"
	service "github.com/kVinsom/Bank-backend/internal/service/user"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// UserHandler translates user HTTP requests into service calls.
type UserHandler struct {
	service service.IUserService
}

// NewUserHandler creates a user HTTP handler.
func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAll godoc
// @Summary List users
// @Tags users
// @Produce json
// @Success 200 {array} core.User
// @Failure 500 {string} string
// @Router /users [get]
func (h *UserHandler) GetAll(ctx *gin.Context) {
	const operation = "get_all"
	defer log.RequestStarted(ctx, operation)()
	users, err := h.service.GetAll(ctx)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, users)
}

// GetById godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} core.User
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [get]
func (h *UserHandler) GetById(ctx *gin.Context) {
	const operation = "get_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	user, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

// GetByEmail godoc
// @Summary Get user by email
// @Tags users
// @Produce json
// @Param email path string true "Email address"
// @Success 200 {object} core.User
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /users/email/{email} [get]
func (h *UserHandler) GetByEmail(ctx *gin.Context) {
	const operation = "get_by_email"
	defer log.RequestStarted(ctx, operation)()
	email := ctx.Param("email")

	user, err := h.service.GetByEmail(ctx, email)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

// GetByPhoneNumber godoc
// @Summary Get user by phone number
// @Tags users
// @Produce json
// @Param phone_number path string true "Phone number"
// @Success 200 {object} core.User
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /users/phone/{phone_number} [get]
func (h *UserHandler) GetByPhoneNumber(ctx *gin.Context) {
	const operation = "get_by_phone_number"
	defer log.RequestStarted(ctx, operation)()
	phoneNumber := ctx.Param("phone_number")

	user, err := h.service.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

// Create godoc
// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param input body core.UserCreateInput true "User data"
// @Success 200 {object} core.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users [post]
func (h *UserHandler) Create(ctx *gin.Context) {
	const operation = "create"
	defer log.RequestStarted(ctx, operation)()
	var input core.UserCreateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.RequestError(ctx, operation, err)
		common.ResponseBadRequest(ctx, err)
		return
	}

	user, err := h.service.Create(ctx, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

// UpdateById godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body core.UserUpdateInput true "User changes"
// @Success 200 {object} core.User
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [put]
func (h *UserHandler) UpdateById(ctx *gin.Context) {
	const operation = "update_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")
	var input core.UserUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
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

	user, err := h.service.Update(ctx, parsedId, &input)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

// ChangePasswordById godoc
// @Summary Change user password
// @Tags users
// @Accept application/x-www-form-urlencoded
// @Param id path int true "User ID"
// @Param hash_password formData string true "New password"
// @Success 200
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /users/{id}/password [patch]
func (h *UserHandler) ChangePasswordById(ctx *gin.Context) {
	const operation = "change_password_by_id"
	defer log.RequestStarted(ctx, operation)()
	id := ctx.Param("id")
	hashPassword := ctx.PostForm("hash_password")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.ChangePassword(ctx, parsedId, hashPassword)
	if err != nil {
		log.RequestError(ctx, operation, err)
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// DeleteById godoc
// @Summary Delete user
// @Tags users
// @Param id path int true "User ID"
// @Success 200
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteById(ctx *gin.Context) {
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
