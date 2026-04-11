package cache

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type ICurrencyCache interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	SetAll(ctx context.Context, currencies []core.Currency) error
	GetById(ctx context.Context, id uint64) (*core.Currency, error)
	SetById(ctx context.Context, currency *core.Currency) error
	GetByIso(ctx context.Context, iso string) (*core.Currency, error)
	SetByIso(ctx context.Context, currency *core.Currency) error
	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	SetBySymbol(ctx context.Context, currency *core.Currency) error
	DeleteById(ctx context.Context, id uint64) error
	DeleteAll(ctx context.Context) error

	MapToCurrency(data map[string]string) (*core.Currency, error)
}
