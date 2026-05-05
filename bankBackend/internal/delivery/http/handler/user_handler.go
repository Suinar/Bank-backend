package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	service "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
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
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, users)
}

func (h *UserHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	user, err := h.service.GetById(ctx, parsedId)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, user)
}

func (h *UserHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := h.service.GetByEmail(ctx, email)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, user)
}

func (h *UserHandler) GetByPhoneNumber(ctx *gin.Context) {
	phoneNumber := ctx.Param("phone_number")

	user, err := h.service.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, user)
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var input core.UserCreateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	user, err := h.service.Create(ctx, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return

		// todo: notification
	}

	ResponseSuccess(ctx, user)
}

func (h *UserHandler) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var input core.UserUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	user, err := h.service.Update(ctx, parsedId, &input)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	ResponseSuccess(ctx, user)
}

func (h *UserHandler) ChangePasswordById(ctx *gin.Context) {
	id := ctx.Param("id")
	hashPassword := ctx.PostForm("hash_password")

	parsedId, err := IdParseHandler(ctx, id)
	if err != nil {
		ErrorHandler(ctx, err)
		return
	}

	err = h.service.ChangePassword(ctx, parsedId, hashPassword)
	if err != nil {
		ErrorHandler(ctx, err)
		return

		// todo: notification
	}

	ResponseSuccess(ctx, nil)
}

func (h *UserHandler) DeleteById(ctx *gin.Context) {
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

		// todo: notification
	}

	ResponseSuccess(ctx, nil)
}
