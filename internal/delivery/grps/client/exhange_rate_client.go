package client

import (
	"context"

	configs "github.com/kVinsom/Bank-backend/internal/configs"
	log "github.com/kVinsom/Bank-backend/internal/logging/client"
	exchanRatePr "github.com/kVinsom/Bank-proto/exchange_rate"
	"google.golang.org/grpc"
)

// NewExchangeRateClient connects to the exchange-rate repository.
func NewExchangeRateClient(
	ctx context.Context,
	cfg configs.Config,
) (exchanRatePr.RankingRepositoryClient, *grpc.ClientConn, error) {
	const clientName = "exchange_rate"
	log.Initializing(clientName, cfg.GRPCConfig.ExchangeRateURL)

	conn, err := Dial(ctx, cfg.GRPCConfig.ExchangeRateURL)
	if err != nil {
		log.InitializationFailed(clientName, cfg.GRPCConfig.ExchangeRateURL, err)
		return nil, nil, err
	}

	client := exchanRatePr.NewRankingRepositoryClient(conn)
	log.Initialized(clientName, cfg.GRPCConfig.ExchangeRateURL)

	return client, conn, nil
}
