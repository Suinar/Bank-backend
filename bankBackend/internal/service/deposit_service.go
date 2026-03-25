package core

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type DepositService struct {
	repository repository.IDepositRepository
}

func NewDepositService(repository repository.IDepositRepository) *DepositService {
	return &DepositService{repository: repository}
}

func (s *DepositService) GetAll(ctx context.Context) ([]core.Deposit, error) {
	deposits, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	filtered := make([]core.Deposit, len(deposits))

	for _, deposit := range deposits {
		if deposit.Status != 0 {
			filtered = append(filtered, deposit)
		}
	}

	return filtered, nil
}

func (s *DepositService) GetByUser(ctx context.Context, userId uint64) ([]core.Deposit, error) {
	deposits, err := s.repository.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	filtered := make([]core.Deposit, len(deposits))

	for _, deposit := range deposits {
		if deposit.Status != 0 {
			filtered = append(filtered, deposit)
		}
	}

	return filtered, nil
}

func (s *DepositService) GetById(ctx context.Context, id uint64) (*core.Deposit, error) {
	deposit, err := s.repository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return deposit, nil
}

func (s *DepositService) Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error) {
}

func (s *DepositService) Repay(ctx context.Context, id uint64, amount int) error {}

func (s *DepositService) Delete(ctx context.Context, id uint64) error {}
