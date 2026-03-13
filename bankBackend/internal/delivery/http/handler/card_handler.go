package core

import (
	core "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	service core.CardRepository
}

func NewCardHandler(service core.CardRepository) *CardHandler {
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
