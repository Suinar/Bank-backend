package client

import (
	"context"

	"github.com/kVinsom/Bank-backend/configs"
	accountPr "github.com/kVinsom/Bank-proto/repository/account"
	cardPr "github.com/kVinsom/Bank-proto/repository/card"
	creditPr "github.com/kVinsom/Bank-proto/repository/credit"
	currencyPr "github.com/kVinsom/Bank-proto/repository/currency"
	depositPr "github.com/kVinsom/Bank-proto/repository/deposit"
	userPr "github.com/kVinsom/Bank-proto/repository/user"
	"google.golang.org/grpc"
)

type RepositoryClients struct {
	Account  accountPr.AccountRepositoryClient
	Card     cardPr.CardRepositoryClient
	Credit   creditPr.CreditRepositoryClient
	Currency currencyPr.CurrencyRepositoryClient
	Deposit  depositPr.DepositRepositoryClient
	User     userPr.UserRepositoryClient
}

func NewRepositoryClient(
	ctx context.Context,
	cfg configs.Config,
) (*RepositoryClients, *grpc.ClientConn, error) {
	conn, err := dial(ctx, cfg.Grpc.RepositoryURL)
	if err != nil {
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

	return clients, conn, nil
}
