package core

import (
	"context"
	"time"

	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repository repository.IAccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) GetAll(ctx context.Context) ([]core.Account, error) {}

func (s *AccountService) GetByUser(ctx context.Context, userId string) ([]core.Account, error) {}

func (s *AccountService) GetById(ctx context.Context, id string) (*core.Account, error) {}

func (s *AccountService) GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Account, error) {
}

func (s *AccountService) Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error) {
}

func (s *AccountService) Blocking(ctx context.Context, id string) error {}

func (s *AccountService) Close(ctx context.Context, id string) error {}

func (s *AccountService) Update(ctx context.Context, id string, input *core.AccountUpdateInput) (*core.Account, error) {
}

func (s *AccountService) Delete(ctx context.Context, id string) error {}
