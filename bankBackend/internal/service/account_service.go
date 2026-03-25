package core

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repository repository.IAccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) GetAll(ctx context.Context) ([]core.Account, error) {}

func (s *AccountService) GetByUser(ctx context.Context, userId uint64) ([]core.Account, error) {
	accounts, err := s.repository.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	filtered := make([]core.Account, 0, len(accounts))

	for _, account := range accounts {
		if account.Status != 0 {
			filtered = append(filtered, account)
		}
	}

	return filtered, nil
}

func (s *AccountService) GetById(ctx context.Context, id uint64) (*core.Account, error) {
	account, err := s.repository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return account, nil
}

func (s *AccountService) Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error) {
}

func (s *AccountService) Blocking(ctx context.Context, id uint64) error {}

func (s *AccountService) Close(ctx context.Context, id uint64) error {}

func (s *AccountService) Update(ctx context.Context, id uint64, input *core.AccountUpdateInput) (*core.Account, error) {
}

func (s *AccountService) Delete(ctx context.Context, id uint64) error {}
