package client

import (
	"context"

	configs "github.com/kVinsom/Bank-backend/internal/configs"
	log "github.com/kVinsom/Bank-backend/internal/logging/client"
	accountPr "github.com/kVinsom/Bank-proto/repository/account"
	cardPr "github.com/kVinsom/Bank-proto/repository/card"
	creditPr "github.com/kVinsom/Bank-proto/repository/credit"
	currencyPr "github.com/kVinsom/Bank-proto/repository/currency"
	depositPr "github.com/kVinsom/Bank-proto/repository/deposit"
	userPr "github.com/kVinsom/Bank-proto/repository/user"
	"google.golang.org/grpc"
)

// RepositoryClients groups the typed clients exposed by the repository service.
type RepositoryClients struct {
	Account  accountPr.AccountRepositoryClient
	Card     cardPr.CardRepositoryClient
	Credit   creditPr.CreditRepositoryClient
	Currency currencyPr.CurrencyRepositoryClient
	Deposit  depositPr.DepositRepositoryClient
	User     userPr.UserRepositoryClient
}

// NewRepositoryClient establishes one shared connection for all repository clients.
func NewRepositoryClient(
	ctx context.Context,
	cfg configs.Config,
) (*RepositoryClients, *grpc.ClientConn, error) {
	const clientName = "repository"
	log.Initializing(clientName, cfg.GRPCConfig.RepositoryURL)

	conn, err := Dial(ctx, cfg.GRPCConfig.RepositoryURL)
	if err != nil {
		log.InitializationFailed(clientName, cfg.GRPCConfig.RepositoryURL, err)
		return nil, nil, err
	}

	clients := &RepositoryClients{
		Account:  accountPr.NewAccountRepositoryClient(conn),
		Card:     cardPr.NewCardRepositoryClient(conn),
		Credit:   creditPr.NewCreditRepositoryClient(conn),
		Currency: currencyPr.NewCurrencyRepositoryClient(conn),
		Deposit:  depositPr.NewDepositRepositoryClient(conn),
		User:     userPr.NewUserRepositoryClient(conn),
	}
	log.Initialized(clientName, cfg.GRPCConfig.RepositoryURL)

	return clients, conn, nil
}
