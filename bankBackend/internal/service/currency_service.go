package core

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type CurrencyService struct {
	repository repository.ICurrencyRepository
}

func NewCurrencyService(repository repository.ICurrencyRepository) *CurrencyService {
	return &CurrencyService{repository: repository}
}

func (service *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {
	currency, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	return currency, nil
}

func (service *CurrencyService) GetById(ctx context.Context, id uint64) (*core.Currency, error) {
	currency, err := service.repository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (service *CurrencyService) GetByIsoCod(ctx context.Context, isoCode string) (*core.Currency, error) {
	currency, err := service.repository.GetByIsoCode(ctx, isoCode)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (service *CurrencyService) GetByNumberCod(ctx context.Context, numberCode string) (*core.Currency, error) {
	currency, err := service.repository.GetByNumberCode(ctx, numberCode)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (service *CurrencyService) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	currency, err := service.repository.GetBySymbol(ctx, symbol)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (service *CurrencyService) Convert(ctx context.Context, currencyIdFrom uint64, amount int, currencyIdTo uint64) (float64, error) {}

func (service *CurrencyService) Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error) {}

func (service *CurrencyService) Update(ctx context.Context, id uint64, input *core.CurrencyUpdateInput) (*core.Currency, error) {}

func (service *CurrencyService) Delete(ctx context.Context, id uint64) error {}

func (service *CurrencyService) GetAllRanking(ctx context.Context, currencyIdFrom uint64) ([]core.ExchangeRate, error) {}

func (service *CurrencyService) GetRelativeRanking(ctx context.Context, currencyIdFrom uint64, currencyIdTo uint64) (*core.ExchangeRate, error) {}
