package currency

import (
	"context"
	"errors"

	exchangeRate "github.com/kVinsom/Bank-proto/exchange_rate"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type CurrencyService struct {
	currencyRepository currencyRepository.CurrencyRepositoryClient
	exchangeRate       exchangeRate.RankingRepositoryClient
}

func NewCurrencyService(
	repository currencyRepository.CurrencyRepositoryClient,
	exchangeRate exchangeRate.RankingRepositoryClient) *CurrencyService {
	return &CurrencyService{
		currencyRepository: repository,
		exchangeRate:       exchangeRate,
	}
}

func (s *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {
	currencies, err := s.currencyCache.GetAll(ctx)
	if err == nil && currencies != nil {
		return currencies, nil
	}

	currencies, err = s.currencyRepository.GetAll(ctx)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	_ = s.currencyCache.SetAll(ctx, currencies)

	return currencies, nil
}

func (s *CurrencyService) GetById(ctx context.Context, id int64) (*core.Currency, error) {
	currency, err := s.currencyCache.GetById(ctx, id)
	if err == nil && currency != nil {
		return currency, nil
	}

	currency, err = s.currencyRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}
		return nil, coreErrors.InternalServerError
	}

	_ = s.currencyCache.Set(ctx, currency)

	return currency, nil
}

func (s *CurrencyService) GetByIso(ctx context.Context, isoCode string) (*core.Currency, error) {
	currency, err := s.currencyCache.GetByIso(ctx, isoCode)
	if err == nil && currency != nil {
		return currency, nil
	}

	currency, err = s.currencyRepository.GetByIso(ctx, isoCode)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	_ = s.currencyCache.Set(ctx, currency)

	return currency, nil
}

func (s *CurrencyService) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	currency, err := s.currencyCache.GetBySymbol(ctx, symbol)
	if err == nil && currency != nil {
		return currency, nil
	}

	currency, err = s.currencyCache.GetBySymbol(ctx, symbol)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	_ = s.currencyCache.Set(ctx, currency)

	return currency, nil
}

func (s *CurrencyService) Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error) {
	if input == nil || input.Name == "" || input.Symbol == '0' ||
		input.MinorUnits <= 0 || input.MinorUnits > 50 ||
		input.IsoCode == "" {
		return nil, coreErrors.BadRequest
	}

	currency := &core.Currency{
		Name:   input.Name,
		Symbol: input.Symbol,
	}

	create, err := s.currencyRepository.Create(ctx, currency)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	err = s.currencyCache.Set(ctx, create)
	if err != nil {
		// todo: logging
	}

	return create, nil
}

func (s *CurrencyService) Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
	if input == nil {
		return nil, coreErrors.BadRequest
	}

	if input.Name != nil {
		if len(*input.Name) == 0 {
			return nil, coreErrors.BadRequest
		}
	}

	if input.MinorUnits != nil {
		if *input.MinorUnits <= 0 {
			return nil, coreErrors.BadRequest
		}
	}

	if input.IsoCode != nil {
		if len(*input.IsoCode) <= 0 || len(*input.IsoCode) > 3 {
			return nil, coreErrors.BadRequest
		}
	}

	currency, err := s.currencyRepository.Update(ctx, id, input)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	err = s.currencyCache.Update(ctx, currency)
	if err != nil {
		// todo: logging
	}

	return currency, nil
}

func (s *CurrencyService) Delete(ctx context.Context, id int64) error {
	err := s.currencyRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	err = s.currencyCache.Delete(ctx, id)
	if err != nil {
		// todo: logging
	}

	return nil
}

func (s *CurrencyService) Convert(ctx context.Context, currencyIdFrom int64, amount int, currencyIdTo int64) (float64, error) {
	if currencyIdFrom == currencyIdTo {
		return float64(amount), nil
	}

	rate, err := s.GetRelativeRanking(ctx, currencyIdFrom, currencyIdTo)
	if err != nil && rate == nil {
		if errors.Is(err, coreErrors.NotFound) {
			return 0, coreErrors.NotFound
		}

		return 0, coreErrors.InternalServerError
	}

	return float64(amount) * float64(rate.RateCross), nil
}

func (s *CurrencyService) GetAllRanking(ctx context.Context, currencyIdFrom int64) ([]core.ExchangeRate, error) {
	resp, err := s.exchangeRate.GetAllRanking(ctx, &exchangeRate.GetAllRankingRequest{
		CurrencyIdFrom: currencyIdFrom,
	})
	if err != nil {
		if errors.Is(err, coreErrors.BadRequest) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}

	var rates []core.ExchangeRate

	for _, r := range resp.Rates {
		rates = append(rates, core.ExchangeRate{
			CurrencyIdFrom: r.CurrencyIdFrom,
			CurrencyIdTo:   r.CurrencyIdTo,
			RateSell:       r.RateSell,
			RateBuy:        r.RateBuy,
			RateCross:      r.RateCross,
		})
	}

	return rates, nil
}

func (s *CurrencyService) GetRelativeRanking(ctx context.Context, currencyIdFrom int64, currencyIdTo int64) (*core.ExchangeRate, error) {
	resp, err := s.exchangeRate.GetRelativeRanking(ctx, &exchangeRate.GetRelativeRankingRequest{
		CurrencyIdFrom: currencyIdFrom,
		CurrencyIdTo:   currencyIdTo,
	})
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	rate := &core.ExchangeRate{
		CurrencyIdFrom: resp.Rate.CurrencyIdFrom,
		CurrencyIdTo:   resp.Rate.CurrencyIdTo,
		RateSell:       resp.Rate.RateSell,
		RateBuy:        resp.Rate.RateBuy,
		RateCross:      resp.Rate.RateCross,
	}

	return rate, nil
}
