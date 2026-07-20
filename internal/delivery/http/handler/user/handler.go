package user

import (
	"github.com/gin-gonic/gin"

	common "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/common"
	service "github.com/kVinsom/Bank-backend/internal/service/user"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.service.GetAll(ctx)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, users)
}

func (h *UserHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	user, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

func (h *UserHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := h.service.GetByEmail(ctx, email)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

func (h *UserHandler) GetByPhoneNumber(ctx *gin.Context) {
	phoneNumber := ctx.Param("phone_number")

	user, err := h.service.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var input core.UserCreateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.ResponseBadRequest(ctx, err)
		return
	}

	user, err := h.service.Create(ctx, &input)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

func (h *UserHandler) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var input core.UserUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	user, err := h.service.Update(ctx, parsedId, &input)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, user)
}

func (h *UserHandler) ChangePasswordById(ctx *gin.Context) {
	id := ctx.Param("id")
	hashPassword := ctx.PostForm("hash_password")

	parsedId, err := common.IdParseHandler(ctx, id)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	err = h.service.ChangePassword(ctx, parsedId, hashPassword)
	if err != nil {
		common.ErrorHandler(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

func (h *UserHandler) DeleteById(ctx *gin.Context) {
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
