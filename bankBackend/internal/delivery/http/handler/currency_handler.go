package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	service core.CurrencyService
}

func NewCurrencyHandler(service core.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{}
}

func (h *CurrencyHandler) GetAll(*gin.Context) {}

func (h *CurrencyHandler) GetById(*gin.Context) {}

func (h *CurrencyHandler) GetByIsoCode(*gin.Context) {}

func (h *CurrencyHandler) GetByNumberCode(*gin.Context) {}

func (h *CurrencyHandler) GetBySymbol(*gin.Context) {}

func (h *CurrencyHandler) Convert(*gin.Context) {}

func (h *CurrencyHandler) Create(*gin.Context) {}

func (h *CurrencyHandler) UpdateById(*gin.Context) {}

func (h *CurrencyHandler) DeleteById(*gin.Context) {}

func (h *CurrencyHandler) GetAllRanking(*gin.Context) {}

func (h *CurrencyHandler) GetRelativeRanking(*gin.Context) {}