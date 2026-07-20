package currency

import (
	"context"

	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type ICurrencyService interface {
	GetAll(context.Context) ([]core.Currency, error)
	GetById(context.Context, int64) (*core.Currency, error)
	GetByIso(context.Context, string) (*core.Currency, error)
	GetBySymbol(context.Context, rune) (*core.Currency, error)
	Create(context.Context, *core.CurrencyCreateInput) (*core.Currency, error)
	Update(context.Context, int64, *core.CurrencyUpdateInput) (*core.Currency, error)
	Delete(context.Context, int64) error
}

type IExchangeRateService interface {
	GetAllRanking(context.Context, int64) ([]core.ExchangeRate, error)
	GetRelativeRanking(context.Context, int64, int64) (*core.ExchangeRate, error)
}

type IConvertService interface {
	Convert(context.Context, int64, int, int64) (float64, error)
}
