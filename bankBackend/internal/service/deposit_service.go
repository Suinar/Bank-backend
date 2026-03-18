package core

import (
	"context"
	"time"

	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type DepositService struct {
	repository repository.IDepositRepository
}

func NewDepositService(repository repository.IDepositRepository) *DepositService {
	return &DepositService{repository: repository}
}

func (s *DepositService) GetAll(ctx context.Context) ([]core.Deposit, error) {}

func (s *DepositService) GetByUser(ctx context.Context, userID string) ([]core.Deposit, error) {}

func (s *DepositService) GetById(ctx context.Context, depositId string) (*core.Deposit, error) {}

func (s *DepositService) GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Deposit, error) {
}

func (s *DepositService) GetByCompletionTime(ctx context.Context, createTime time.Time) (*core.Deposit, error) {
}

func (s *DepositService) Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error) {
}

func (s *DepositService) Repay(ctx context.Context, id string, amount int) error {}

func (s *DepositService) Delete(ctx context.Context, id string) error {}
