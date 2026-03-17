package core

import (
	"context"
	"time"

	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetAll(ctx context.Context) ([]core.Account, error) {}

func (r *AccountRepository) GetByUser(ctx context.Context, idUser uint64) (*core.Account, error) {}

func (r *AccountRepository) GetById(ctx context.Context, id uint64) (*core.Account, error) {}

func (r *AccountRepository) GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Account, error) {
}

func (r *AccountRepository) Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error) {
}

func (r *AccountRepository) Update(ctx context.Context, input *core.AccountUpdateInput) (*core.Account, error) {
}

func (r *AccountRepository) Delete(ctx context.Context, id uint64) error {}
