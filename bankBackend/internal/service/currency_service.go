package core

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	cache "github.com/Suinar/Bank-backend/bankBackend/internal/repository/cache"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	notificationService "github.com/Suinar/Bank-backend/bankBackend/proto/notification"
	rankingService "github.com/Suinar/Bank-exhange-rate-service/ranking"
)

type CurrencyService struct {
	currencyRepository  repository.ICurrencyRepository
	currencyCache       cache.ICurrencyCache
	rankingService      rankingService.RankingServiceClient
	notificationService notificationService.NotificationServiceClient
}

func NewCurrencyService(
	repository repository.ICurrencyRepository,
	currencyCache cache.ICurrencyCache,
	rankingService rankingService.RankingServiceClient,
	notificationService notificationService.NotificationServiceClient) *CurrencyService {
	return &CurrencyService{
		currencyRepository:  repository,
		currencyCache:       currencyCache,
		rankingService:      rankingService,
		notificationService: notificationService,
	}
}

func (s *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {
	currencies, err := s.currencyCache.GetAll(ctx)
	if err == nil && currencies != nil {
		return currencies, nil
	}

	currencies, err = s.currencyRepository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
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
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}
		return nil, core.InternalServerError
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
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
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
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	_ = s.currencyCache.Set(ctx, currency)

	return currency, nil
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

	err = s.currencyCache.Set(ctx, create)
	if err != nil {
		// todo: logging
	}

	s.notificationService.SendEvent(ctx, &notificationService.NotificationEventRequest{
		Entity:   notificationService.EntityType_CURRENCY,
		Action:   notificationService.ActionType_CREATE,
		EntityId: create.Id,
		UserId:   0,
	})

	return create, nil
}

func (s *CurrencyService) Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
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

	err = s.currencyCache.Update(ctx, currency)
	if err != nil {
		// todo: logging
	}

	return currency, nil
}

func (s *CurrencyService) Delete(ctx context.Context, id int64) error {
	err := s.currencyRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	err = s.currencyCache.Delete(ctx, id)
	if err != nil {
		// todo: logging
	}

	s.notificationService.SendEvent(ctx, &notificationService.NotificationEventRequest{
		Entity:   notificationService.EntityType_CURRENCY,
		Action:   notificationService.ActionType_DELETE,
		EntityId: id,
		UserId:   0,
	})

	return nil
}

func (s *CurrencyService) Convert(ctx context.Context, currencyIdFrom int64, amount int, currencyIdTo int64) (float64, error) {
	if currencyIdFrom == currencyIdTo {
		return float64(amount), nil
	}

	rate, err := s.GetRelativeRanking(ctx, currencyIdFrom, currencyIdTo)
	if err != nil && rate == nil {
		if errors.Is(err, core.NotFound) {
			return 0, core.NotFound
		}

		return 0, core.InternalServerError
	}

	return float64(amount) * float64(rate.RateCross), nil
}

func (s *CurrencyService) GetAllRanking(ctx context.Context, currencyIdFrom int64) ([]core.ExchangeRate, error) {
	resp, err := s.rankingService.GetAllRanking(ctx, &rankingService.GetAllRankingRequest{
		CurrencyIdFrom: currencyIdFrom,
	})
	if err != nil {
		if errors.Is(err, core.BadRequest) {
			return nil, core.BadRequest
		}

		return nil, core.InternalServerError
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
	resp, err := s.rankingService.GetRelativeRanking(ctx, &rankingService.GetRelativeRankingRequest{
		CurrencyIdFrom: currencyIdFrom,
		CurrencyIdTo:   currencyIdTo,
	})
	if err != nil {
		return nil, core.InternalServerError
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
