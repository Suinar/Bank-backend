package exhange_rate

import (
	"context"
	"errors"
	log "github.com/kVinsom/Bank-backend/internal/logging/service/exhange_rate"

	exchangeRate "github.com/Suinar/Bank-exhange-rate-service/pkg/core"
	exchangeRatePr "github.com/kVinsom/Bank-proto/exchange_rate"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
)

// ExchangeRateService exposes exchange-rate queries to the delivery layer.
type ExchangeRateService struct {
	exchangeRate exchangeRatePr.RankingRepositoryClient
}

// NewExchangeRateService creates a service backed by the ranking repository.
func NewExchangeRateService(
	exchangeRate exchangeRatePr.RankingRepositoryClient) *ExchangeRateService {
	return &ExchangeRateService{
		exchangeRate: exchangeRate,
	}
}

func (s *ExchangeRateService) GetAllRanking(ctx context.Context, currencyIsoFrom int) ([]exchangeRate.Ranking, error) {
	const operation = "get_all_ranking"
	defer log.OperationStarted(operation)()
	if s.exchangeRate == nil {
		return nil, coreErrors.InternalServerError
	}
	if !validCurrencyIso(currencyIsoFrom) {
		return nil, coreErrors.BadRequest
	}

	resp, err := s.exchangeRate.GetAllRanking(ctx, &exchangeRatePr.GetAllRankingRequest{
		CurrencyIsoFrom: int32(currencyIsoFrom),
	})
	if err != nil {
		if errors.Is(err, coreErrors.BadRequest) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}
	if resp == nil {
		return nil, coreErrors.InternalServerError
	}

	rates := make([]exchangeRate.Ranking, 0, len(resp.Rankings))

	for _, ranking := range resp.Rankings {
		if mapped := RankingToCore(ranking); mapped != nil {
			rates = append(rates, *mapped)
		}
	}

	return rates, nil
}

func (s *ExchangeRateService) GetRelativeRanking(ctx context.Context, currencyIsoFrom int, currencyIsoTo int) (*exchangeRate.Ranking, error) {
	const operation = "get_relative_ranking"
	defer log.OperationStarted(operation)()
	if s.exchangeRate == nil {
		return nil, coreErrors.InternalServerError
	}
	if !validCurrencyIso(currencyIsoFrom) || !validCurrencyIso(currencyIsoTo) {
		return nil, coreErrors.BadRequest
	}

	resp, err := s.exchangeRate.GetRelativeRanking(ctx, &exchangeRatePr.GetRelativeRankingRequest{
		CurrencyIsoFrom: int32(currencyIsoFrom),
		CurrencyIsoTo:   int32(currencyIsoTo),
	})
	if err != nil {
		if errors.Is(err, coreErrors.BadRequest) {
			return nil, coreErrors.BadRequest
		}
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}
	if resp == nil {
		return nil, coreErrors.NotFound
	}

	return RankingToCore(resp), nil
}

func validCurrencyIso(currencyISO int) bool {
	return currencyISO > 0 && currencyISO <= 999
}
