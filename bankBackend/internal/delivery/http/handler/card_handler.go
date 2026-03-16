package core

import (
	service "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	service service.ICardRepository
}

func NewCardHandler(service service.ICardRepository) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) GetAll(*gin.Context) {}

func (h *CardHandler) GetByUser(*gin.Context) {}

func (h *CardHandler) GetById(*gin.Context) {}

func (h *CardHandler) GetByNumber(*gin.Context) {}

func (h *CardHandler) GetByCreateTime(*gin.Context) {}

func (h *CardHandler) Create(*gin.Context) {}

func (h *CardHandler) BlockingById(*gin.Context) {}

func (h *CardHandler) DeleteById(*gin.Context) {}
