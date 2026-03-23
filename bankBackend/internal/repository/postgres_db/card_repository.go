package core

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) GetAll(ctx context.Context) ([]core.Card, error) {}

func (r *CardRepository) GetByUser(ctx context.Context, idUser uint64) (*core.Card, error) {}

func (r *CardRepository) GetById(ctx context.Context, id uint64) (*core.Card, error) {}

func (r *CardRepository) GetByNumber(ctx context.Context, number string) (*core.Account, error) {}

func (r *CardRepository) Blocking(ctx context.Context, id uint64) error {}

func (r *CardRepository) Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error) {
}

func (r *CardRepository) Delete(ctx context.Context, id uint64) error {}
