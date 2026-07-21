package fixture

import (
	"testing"

	"github.com/golang/mock/gomock"
	exchangeRateMocks "github.com/kVinsom/Bank-backend/internal/mocks/exchange_rate"
	repositoryMocks "github.com/kVinsom/Bank-backend/internal/mocks/repositories"
)

type UserRepositoryMocks struct {
	User *repositoryMocks.MockUserRepositoryClient
}

func NewUserRepositoryMocks(t *testing.T) *UserRepositoryMocks {
	t.Helper()
	return &UserRepositoryMocks{
		User: repositoryMocks.NewMockUserRepositoryClient(gomock.NewController(t)),
	}
}

type CurrencyRepositoryMocks struct {
	Currency *repositoryMocks.MockCurrencyRepositoryClient
}

func NewCurrencyRepositoryMocks(t *testing.T) *CurrencyRepositoryMocks {
	t.Helper()
	return &CurrencyRepositoryMocks{
		Currency: repositoryMocks.NewMockCurrencyRepositoryClient(gomock.NewController(t)),
	}
}

type CardRepositoryMocks struct {
	Card    *repositoryMocks.MockCardRepositoryClient
	User    *repositoryMocks.MockUserRepositoryClient
	Account *repositoryMocks.MockAccountRepositoryClient
}

func NewCardRepositoryMocks(t *testing.T) *CardRepositoryMocks {
	t.Helper()
	ctrl := gomock.NewController(t)
	return &CardRepositoryMocks{
		Card:    repositoryMocks.NewMockCardRepositoryClient(ctrl),
		User:    repositoryMocks.NewMockUserRepositoryClient(ctrl),
		Account: repositoryMocks.NewMockAccountRepositoryClient(ctrl),
	}
}

type CreditRepositoryMocks struct {
	Credit   *repositoryMocks.MockCreditRepositoryClient
	User     *repositoryMocks.MockUserRepositoryClient
	Currency *repositoryMocks.MockCurrencyRepositoryClient
}

func NewCreditRepositoryMocks(t *testing.T) *CreditRepositoryMocks {
	t.Helper()
	ctrl := gomock.NewController(t)
	return &CreditRepositoryMocks{
		Credit:   repositoryMocks.NewMockCreditRepositoryClient(ctrl),
		User:     repositoryMocks.NewMockUserRepositoryClient(ctrl),
		Currency: repositoryMocks.NewMockCurrencyRepositoryClient(ctrl),
	}
}

type DepositRepositoryMocks struct {
	Deposit  *repositoryMocks.MockDepositRepositoryClient
	User     *repositoryMocks.MockUserRepositoryClient
	Currency *repositoryMocks.MockCurrencyRepositoryClient
}

func NewDepositRepositoryMocks(t *testing.T) *DepositRepositoryMocks {
	t.Helper()
	ctrl := gomock.NewController(t)
	return &DepositRepositoryMocks{
		Deposit:  repositoryMocks.NewMockDepositRepositoryClient(ctrl),
		User:     repositoryMocks.NewMockUserRepositoryClient(ctrl),
		Currency: repositoryMocks.NewMockCurrencyRepositoryClient(ctrl),
	}
}

type ExchangeRateRepositoryMocks struct {
	Ranking *exchangeRateMocks.MockRankingRepositoryClient
}

func NewExchangeRateRepositoryMocks(t *testing.T) *ExchangeRateRepositoryMocks {
	t.Helper()
	return &ExchangeRateRepositoryMocks{
		Ranking: exchangeRateMocks.NewMockRankingRepositoryClient(gomock.NewController(t)),
	}
}
