package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type CreditHandler struct {
	service core.CreditService
}

func NewCreditHandler(service core.CreditService) *CreditHandler {
	return &CreditHandler{service: service}
}

func (h *CreditHandler) GetAll(*gin.Context) {}

func (h *CreditHandler) GetByUser(*gin.Context) {}

func (h *CreditHandler) GetById (*gin.Context) {}

func (h *CreditHandler) GetByCreateTime(*gin.Context) {}

func (h *CreditHandler) GetByRepayTime(*gin.Context) {}

func (h *CreditHandler) Create(*gin.Context) {}

func (h *CreditHandler) RepayById(*gin.Context) {}

func (h *CreditHandler) DeleteById(*gin.Context) {}