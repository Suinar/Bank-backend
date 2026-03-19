package core

import (
	"context"
	"time"

	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
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

func (s *DepositService) Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error) {
}

func (s *DepositService) Repay(ctx context.Context, id string, amount int) error {}

func (s *DepositService) Delete(ctx context.Context, id string) error {}
