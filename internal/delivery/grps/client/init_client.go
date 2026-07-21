package client

import (
	"context"
	"errors"
	"fmt"

	exchangeRateProto "github.com/kVinsom/Bank-proto/exchange_rate"
	"google.golang.org/grpc"

	"github.com/kVinsom/Bank-backend/internal/configs"
	log "github.com/kVinsom/Bank-backend/internal/logging/client"
)

// Clients owns all outbound gRPC clients and their underlying connections.
type Clients struct {
	Repositories *RepositoryClients
	ExchangeRate exchangeRateProto.RankingRepositoryClient

	repositoryConn   *grpc.ClientConn
	exchangeRateConn *grpc.ClientConn
}

// InitClients connects all upstream dependencies and rolls back partial setup on failure.
func InitClients(ctx context.Context, cfg configs.Config) (*Clients, error) {
	log.ContainerInitializing()
	repositories, repositoryConn, err := NewRepositoryClient(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("connect to repository service: %w", err)
	}

	exchangeRate, exchangeRateConn, err := NewExchangeRateClient(ctx, cfg)
	if err != nil {
		log.RollbackStarting("repository")
		if closeErr := repositoryConn.Close(); closeErr != nil {
			log.RollbackFailed("repository", closeErr)
		}
		return nil, fmt.Errorf("connect to exchange-rate service: %w", err)
	}

	log.ContainerInitialized()
	return &Clients{
		Repositories:     repositories,
		ExchangeRate:     exchangeRate,
		repositoryConn:   repositoryConn,
		exchangeRateConn: exchangeRateConn,
	}, nil
}

// Close releases every connection owned by the client container.
func (c *Clients) Close() error {
	if c == nil {
		return nil
	}

	var closeErrors []error
	if c.exchangeRateConn != nil {
		log.Closing("exchange_rate")
		if err := c.exchangeRateConn.Close(); err != nil {
			log.CloseFailed("exchange_rate", err)
			closeErrors = append(closeErrors, err)
		} else {
			log.Closed("exchange_rate")
		}
	}
	if c.repositoryConn != nil {
		log.Closing("repository")
		if err := c.repositoryConn.Close(); err != nil {
			log.CloseFailed("repository", err)
			closeErrors = append(closeErrors, err)
		} else {
			log.Closed("repository")
		}
	}

	closeErr := errors.Join(closeErrors...)
	if closeErr == nil {
		log.AllClosed()
	}
	return closeErr
}
