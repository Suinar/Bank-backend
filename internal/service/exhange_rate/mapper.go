package exhange_rate

import (
	log "github.com/kVinsom/Bank-backend/internal/logging/service/exhange_rate"
	"time"

	exchangeRate "github.com/Suinar/Bank-exhange-rate-service/pkg/core"
	exchangeRatePr "github.com/kVinsom/Bank-proto/exchange_rate"
)

// RankingToCore maps an exchange-rate response into the domain model.
func RankingToCore(ranking *exchangeRatePr.Ranking) *exchangeRate.Ranking {
	const operation = "ranking_to_core"
	defer log.MappingStarted(operation)()
	if ranking == nil {
		log.NilInput(operation)
		return nil
	}

	date := time.Time{}
	if ranking.Date != nil {
		date = ranking.Date.AsTime()
	}

	return &exchangeRate.Ranking{
		CurrencyIsoFrom: ranking.CurrencyIsoFrom,
		CurrencyIsoTo:   ranking.CurrencyIsoTo,
		Date:            date,
		RateSell:        ranking.RateSell,
		RateBuy:         ranking.RateBuy,
		RateCross:       ranking.RateCross,
	}
}
