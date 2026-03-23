package core

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type CreditRepository struct {
	db *sqlx.DB
}

func NewCreditRepository(db *sqlx.DB) *CreditRepository {
	return &CreditRepository{db: db}
}

func (r *CreditRepository) GetAll(ctx context.Context) ([]core.Credit, error) {}

func (r *CreditRepository) GetByUser(ctx context.Context, idUser uint64) (*core.Credit, error) {}

func (r *CreditRepository) GetById(ctx context.Context, id uint64) (*core.Credit, error) {}

func (r *CreditRepository) Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error) {
}

func (r *CreditRepository) Replay(ctx context.Context, id uint64) (*core.Credit, error) {}

func (r *CreditRepository) Delete(ctx context.Context, id uint64) error {}
