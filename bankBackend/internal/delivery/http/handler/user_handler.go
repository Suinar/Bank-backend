package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service core.UserService
}

func NewUserHandler(service core.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(*gin.Context) {}

func (h *UserHandler) GetById(*gin.Context) {}

func (h *UserHandler) GetByEmail(*gin.Context) {}

func (h *UserHandler) GetByPhoneNumber(*gin.Context) {}

func (h *UserHandler) GetMe(*gin.Context) {}

func (h *UserHandler) Create(*gin.Context) {}

func (h *UserHandler) UpdateById(*gin.Context) {}

func (h *UserHandler) ChangePasswordById(*gin.Context) {}

func (h *UserHandler) DeleteById(*gin.Context) {}
