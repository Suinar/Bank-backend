package exhange_rate

import (
	"context"

	exhangeRate "github.com/Suinar/Bank-exhange-rate-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/exchange_rate.go -package=mocks

// IExchangeRateService defines exchange-rate queries required by HTTP delivery.
type IExchangeRateService interface {
	GetAllRanking(context.Context, int) ([]exhangeRate.Ranking, error)
	GetRelativeRanking(context.Context, int, int) (*exhangeRate.Ranking, error)
}
