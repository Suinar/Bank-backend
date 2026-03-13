package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type DepositHandler struct {
	service core.DepositService
}

func NewDepositHandler(service core.DepositService) *DepositHandler {
	return &DepositHandler{service: service}
}

func (h *DepositHandler) GetAll(*gin.Context) {}

func (h *DepositHandler) GetByUser(*gin.Context) {}

func (h *DepositHandler) GetById(*gin.Context) {}

func (h *DepositHandler) GetByCreateTime(*gin.Context) {}

func (h *DepositHandler) GetByCompletionTime(*gin.Context) {}

func (h *DepositHandler) Create(*gin.Context) {}

func (h *DepositHandler) RepayById(*gin.Context) {}

func (h *DepositHandler) DeleteById(*gin.Context) {}