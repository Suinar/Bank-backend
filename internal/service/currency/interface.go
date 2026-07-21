package currency

import (
	"context"

	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/currency.go -package=mocks

// ICurrencyService defines currency operations required by HTTP delivery.
type ICurrencyService interface {
	GetAll(context.Context) ([]core.Currency, error)
	GetById(context.Context, int64) (*core.Currency, error)
	GetByIso(context.Context, string) (*core.Currency, error)
	GetBySymbol(context.Context, rune) (*core.Currency, error)
	Create(context.Context, *core.CurrencyCreateInput) (*core.Currency, error)
	Update(context.Context, int64, *core.CurrencyUpdateInput) (*core.Currency, error)
	Delete(context.Context, int64) error
}
