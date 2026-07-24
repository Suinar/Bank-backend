package fixture

import (
	"time"

	core "github.com/Suinar/Bank-exhange-rate-service/pkg/core"
	"github.com/kVinsom/Bank-backend/internal/test"
	exchangeRateProto "github.com/kVinsom/Bank-proto/exchange_rate"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RankingCore() core.Ranking {
	return core.Ranking{
		CurrencyIsoFrom: test.ExchangeISOFrom,
		CurrencyIsoTo:   test.ExchangeISOTo,
		Date:            time.Date(2026, time.July, 21, 0, 0, 0, 0, time.UTC),
		RateSell:        1.1,
		RateBuy:         1.0,
		RateCross:       1.05,
	}
}

func RankingListCore() []core.Ranking { return []core.Ranking{RankingCore()} }

func RankingProto() *exchangeRateProto.Ranking {
	ranking := RankingCore()
	return &exchangeRateProto.Ranking{
		CurrencyIsoFrom: ranking.CurrencyIsoFrom,
		CurrencyIsoTo:   ranking.CurrencyIsoTo,
		Date:            timestamppb.New(ranking.Date),
		RateSell:        ranking.RateSell,
		RateBuy:         ranking.RateBuy,
		RateCross:       ranking.RateCross,
	}
}

func RankingListProto(rankings ...*exchangeRateProto.Ranking) *exchangeRateProto.RankingList {
	return &exchangeRateProto.RankingList{
		Rankings: rankings,
	}
}
