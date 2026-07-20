package client

import (
	"context"

	"github.com/kVinsom/Bank-backend/configs"
	exchanRatePr "github.com/kVinsom/Bank-proto/exchange_rate"
	"google.golang.org/grpc"
)

func NewExchangeRateClient(
	ctx context.Context,
	cfg configs.Config,
) (exchanRatePr.RankingRepositoryClient, *grpc.ClientConn, error) {
	conn, err := dial(ctx, cfg.Grpc.ExchangeRateURL)
	if err != nil {
		return nil, nil, err
	}

	client := exchanRatePr.NewRankingRepositoryClient(conn)

	return client, conn, nil
}
