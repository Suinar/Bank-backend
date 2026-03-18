package core

import (
	"context"

	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type CurrencyRepository struct {
	db *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

func (r *CurrencyRepository) GetAll(ctx context.Context) ([]core.Currency, error) {}

func (r *CurrencyRepository) GetByUser(ctx context.Context, idUser string) (*core.Currency, error) {}

func (r *CurrencyRepository) GetById(ctx context.Context, id string) (*core.Currency, error) {}

func (r *CurrencyRepository) GetByIsoCode(ctx context.Context, isoCode string) (*core.Currency, error) {}

func (r *CurrencyRepository) GetByNumberCode(ctx context.Context, numberCode string) (*core.Currency, error) {}

func (r *CurrencyRepository) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {}

func (r *CurrencyRepository) Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error) {}

func (r *CurrencyRepository) Update(ctx context.Context, input *core.CurrencyUpdateInput) (*core.Currency, error) {}

func (r *CurrencyRepository) Delete(ctx context.Context, id string) error {}
