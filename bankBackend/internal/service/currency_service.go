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

func (s *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {
	currency, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	return currency, nil
}

func (s *CurrencyService) GetById(ctx context.Context, id uint64) (*core.Currency, error) {
	currency, err := s.repository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (s *CurrencyService) GetByIsoCod(ctx context.Context, isoCode string) (*core.Currency, error) {
	currency, err := s.repository.GetByIsoCode(ctx, isoCode)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (s *CurrencyService) GetByNumberCod(ctx context.Context, numberCode string) (*core.Currency, error) {
	currency, err := s.repository.GetByNumberCode(ctx, numberCode)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (s *CurrencyService) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	currency, err := s.repository.GetBySymbol(ctx, symbol)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (s *CurrencyService) Convert(ctx context.Context, currencyIdFrom uint64, amount int, currencyIdTo uint64) (float64, error) {
	// todo: convert service

	return 0, nil
}

func (s *CurrencyService) Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error) {
	return nil, nil
}

func (s *CurrencyService) Update(ctx context.Context, id uint64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
	return nil, nil
}

func (s *CurrencyService) Delete(ctx context.Context, id uint64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// Maybe todo: notification

	return nil
}

func (s *CurrencyService) GetAllRanking(ctx context.Context, currencyIdFrom uint64) ([]core.ExchangeRate, error) {
	// todo: currency ranking service

	return nil, nil
}

func (s *CurrencyService) GetRelativeRanking(ctx context.Context, currencyIdFrom uint64, currencyIdTo uint64) (*core.ExchangeRate, error) {
	// todo: currency ranking service

	return nil, nil
}
