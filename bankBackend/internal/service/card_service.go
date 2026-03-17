package core

import (
	"context"
	"time"

	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type CardService struct {
	repository repository.ICardRepository
}

func NewCardService(repository repository.ICardRepository) *CardService {
	return &CardService{repository: repository}
}

func (s *CardService) GetAll(ctx context.Context) ([]core.Card, error) {}

func (s *CardService) GetByUser(ctx context.Context, userId uint64) ([]core.Card, error) {}

func (s *CardService) GetById(ctx context.Context, id uint64) (*core.Card, error) {}

func (s *CardService) GetByNumber(ctx context.Context, number string) (*core.Card, error) {}

func (s *CardService) GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Card, error) {
}

func (s *CardService) Blocking(ctx context.Context, id uint64) error {}

func (s *CardService) Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error) {}

func (s *CardService) Delete(ctx context.Context, id uint64) error {}
