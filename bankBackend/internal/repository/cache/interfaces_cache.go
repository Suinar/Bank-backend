package cache

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type ICurrencyCache interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	GetById(ctx context.Context, id uint64) (*core.Currency, error)
	GetByIso(ctx context.Context, iso string) (*core.Currency, error)
	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	Set(ctx context.Context, currency *core.Currency) error
	SetAll(ctx context.Context, currencies []core.Currency) error
	UpdateById(ctx context.Context, currency *core.Currency) error
	Delete(ctx context.Context, id uint64) error
}
