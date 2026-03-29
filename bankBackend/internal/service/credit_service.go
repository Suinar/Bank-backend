package core

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type CreditService struct {
	repository repository.ICreditRepository
}

func NewCreditService(repository repository.ICreditRepository) *CreditService {
	return &CreditService{repository: repository}
}

func (s *CreditService) GetAll(ctx context.Context) ([]core.Credit, error) {
	credits, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	filtered := make([]core.Credit, len(credits))

	for _, credit := range credits {
		if credit.Status != core.CreditStatusClosed {
			filtered = append(filtered, credit)
		}
	}

	return filtered, nil
}

func (s *CreditService) GetByUser(ctx context.Context, userId uint64) ([]core.Credit, error) {
	credits, err := s.repository.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	filtered := make([]core.Credit, len(credits))

	for _, credit := range credits {
		if credit.Status != core.CreditStatusClosed {
			filtered = append(filtered, credit)
		}
	}

	return filtered, nil
}

func (s *CreditService) GetById(ctx context.Context, id uint64) (*core.Credit, error) {
	credit, err := s.repository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return credit, nil
}

func (s *CreditService) Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error) {
	return nil, nil
}

func (s *CreditService) Repay(ctx context.Context, id uint64, amount int) error {
	// todo: repay service

	return nil
}

func (s *CreditService) Delete(ctx context.Context, id uint64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// todo: notification

	return nil
}
