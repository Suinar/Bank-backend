package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const connectionTimeout = 5 * time.Second

func dial(ctx context.Context, target string) (*grpc.ClientConn, error) {
	connectionContext, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	return grpc.DialContext(
		connectionContext,
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
}
