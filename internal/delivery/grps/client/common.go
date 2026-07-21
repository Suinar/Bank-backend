package client

import (
	"context"
	"time"

	log "github.com/kVinsom/Bank-backend/internal/logging/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const connectionTimeout = 5 * time.Second

// Dial opens a blocking gRPC connection bounded by the shared connection timeout.
func Dial(ctx context.Context, target string) (*grpc.ClientConn, error) {
	startedAt := time.Now()
	log.ConnectionStarting(target, connectionTimeout)

	connectionContext, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	conn, err := grpc.DialContext(
		connectionContext,
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.ConnectionFailed(target, time.Since(startedAt), err)
		return nil, err
	}

	log.ConnectionEstablished(target, time.Since(startedAt))
	return conn, nil
}
