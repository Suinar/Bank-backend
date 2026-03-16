package core

import (
	"context"
	"time"

	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type CreditService struct {
	repository repository.ICreditRepository
}

func NewCreditService(repository repository.ICreditRepository) *CreditService {
	return &CreditService{repository: repository}
}

func (s *CreditService) GetAll(ctx context.Context) ([]core.Credit, error) {}

func (s *CreditService) GetByUser(ctx context.Context, userId string) ([]core.Credit, error) {}

func (s *CreditService) GetById(ctx context.Context, id string) (*core.Credit, error) {}

func (s *CreditService) GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Credit, error) {
}

func (s *CreditService) GetByRepayTime(ctx context.Context, repayTime time.Time) (*core.Credit, error) {
}

func (s *CreditService) Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error) {
}

func (s *CreditService) Repay(ctx context.Context, id string, amount int) error {}

func (s *CreditService) Delete(ctx context.Context, id string) error {}
