package core

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type CardService struct {
	repository repository.ICardRepository
}

func NewCardService(repository repository.ICardRepository) *CardService {
	return &CardService{repository: repository}
}

func (s *CardService) GetAll(ctx context.Context) ([]core.Card, error) {
	cards, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	filtered := make([]core.Card, len(cards))

	for _, card := range cards {
		if card.Status != 0 {
			filtered = append(filtered, card)
		}
	}

	return filtered, nil
}

func (s *CardService) GetByUser(ctx context.Context, userId string) ([]core.Card, error) {}

func (s *CardService) GetById(ctx context.Context, id string) (*core.Card, error) {}

func (s *CardService) GetByNumber(ctx context.Context, number string) (*core.Card, error) {}

func (s *CardService) Blocking(ctx context.Context, id string) error {}

func (s *CardService) Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error) {}

func (s *CardService) Delete(ctx context.Context, id string) error {}
