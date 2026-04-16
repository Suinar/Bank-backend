package client

import (
	"context"
	"time"

	"github.com/Suinar/Bank-backend/bankBackend/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	ranking "github.com/Suinar/Bank-backend/bankBackend/proto"
)

func NewCurrencyRankingClient(ctx context.Context, cfg configs.Config) (ranking.RankingServiceClient, *grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		cfg.Grps.CurrencyRankingUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, nil, err
	}

	client := ranking.NewRankingServiceClient(conn)

	return client, conn, nil
}
