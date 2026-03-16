package core

import (
	service "github.com/Suinar/Bank-backend/bankBackend/internal/service"
	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	service service.ICurrencyService
}

func NewCurrencyHandler(service service.ICurrencyService) *CurrencyHandler {
	return &CurrencyHandler{service: service}
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
