package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service core.AccountService
}

func NewAccountHandler(service core.AccountService) *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) GetAll(*gin.Context) {}

func (h *AccountHandler) GetByUser(*gin.Context) {}

func (h *AccountHandler) GetById(*gin.Context) {}

func (h *AccountHandler) GetByCreateTime(*gin.Context) {}

func (h *AccountHandler) Create(*gin.Context) {}

func (h *AccountHandler) BlockingById(*gin.Context) {}

func (h *AccountHandler) CloseById(*gin.Context) {}

func (h *AccountHandler) UpdateById(*gin.Context) {}

func (h *AccountHandler) DeleteById(*gin.Context) {}
