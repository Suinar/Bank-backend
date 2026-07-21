package exchange_rate

import exchangeRatePr "github.com/kVinsom/Bank-proto/exchange_rate"

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -destination=ranking.go -package=exchange_rate github.com/kVinsom/Bank-backend/internal/mocks/exchange_rate RankingRepositoryClient

type RankingRepositoryClient interface {
	exchangeRatePr.RankingRepositoryClient
}
