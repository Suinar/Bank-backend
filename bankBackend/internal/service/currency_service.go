package core

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type CurrencyService struct {
	currencyRepository repository.ICurrencyRepository
}

func NewCurrencyService(repository repository.ICurrencyRepository) *CurrencyService {
	return &CurrencyService{currencyRepository: repository}
}

func (s *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {
	currency, err := s.currencyRepository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	return currency, nil
}

func (s *CurrencyService) GetById(ctx context.Context, id uint64) (*core.Currency, error) {
	currency, err := s.currencyRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (s *CurrencyService) GetByIsoCod(ctx context.Context, isoCode string) (*core.Currency, error) {
	currency, err := s.currencyRepository.GetByIsoCode(ctx, isoCode)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return currency, nil
}

func (s *CurrencyService) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	currency, err := s.currencyRepository.GetBySymbol(ctx, symbol)
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
	if input == nil || input.Name == "" || input.Symbol == '0' ||
		input.MinorUnits <= 0 || input.MinorUnits > 50 ||
		input.IsoCode == "" {
		return nil, core.BadRequest
	}

	currency := &core.Currency{
		Name:   input.Name,
		Symbol: input.Symbol,
	}

	create, err := s.currencyRepository.Create(ctx, currency)
	if err != nil {
		return nil, core.InternalServerError
	}

	// Maybe server todo: notification

	return create, nil
}

func (s *CurrencyService) Update(ctx context.Context, id uint64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
	if input == nil {
		return nil, core.BadRequest
	}

	if input.Name != nil {
		if len(*input.Name) == 0 {
			return nil, core.BadRequest
		}
	}

	if input.MinorUnits != nil {
		if *input.MinorUnits <= 0 {
			return nil, core.BadRequest
		}
	}

	if input.IsoCode != nil {
		if len(*input.IsoCode) <= 0 || len(*input.IsoCode) > 3 {
			return nil, core.BadRequest
		}
	}

	currency, err := s.currencyRepository.Update(ctx, id, input)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	// Maybe server todo: notification

	return currency, nil
}

func (s *CurrencyService) Delete(ctx context.Context, id uint64) error {
	err := s.currencyRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// Maybe server todo: notification

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
