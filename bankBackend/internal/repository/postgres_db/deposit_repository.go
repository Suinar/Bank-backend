package core

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type DepositRepository struct {
	db *sqlx.DB
}

func NewDepositRepository(db *sqlx.DB) *DepositRepository {
	return &DepositRepository{db: db}
}

func (r *DepositRepository) GetAll(ctx context.Context) ([]core.Deposit, error) {}

func (r *DepositRepository) GetByUser(ctx context.Context, idUser uint64) ([]core.Deposit, error) {}

func (r *DepositRepository) GetById(ctx context.Context, id uint64) (*core.Deposit, error) {}

func (r *DepositRepository) Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error) {}

func (r *DepositRepository) Repay(ctx context.Context, id uint64, amount int) error {}

func (r *DepositRepository) Delete(ctx context.Context, id uint64) error {}
