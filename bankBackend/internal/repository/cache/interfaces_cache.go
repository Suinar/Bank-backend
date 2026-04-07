package cache

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type ICurrencyCache interface {
	GetByID(ctx context.Context, id uint64) (*core.Currency, error)
	SetByID(ctx context.Context, currency *core.Currency) error

	GetByISO(ctx context.Context, iso string) (*core.Currency, error)
	SetByISO(ctx context.Context, currency *core.Currency) error

	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	SetBySymbol(ctx context.Context, currency *core.Currency) error

	Delete(ctx context.Context, id uint64) error
}
