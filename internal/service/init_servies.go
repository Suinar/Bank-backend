package service

import (
	exchangeRateProto "github.com/kVinsom/Bank-proto/exchange_rate"

	"github.com/kVinsom/Bank-backend/internal/configs"
	grpcClient "github.com/kVinsom/Bank-backend/internal/delivery/grps/client"
	accountService "github.com/kVinsom/Bank-backend/internal/service/account"
	cardService "github.com/kVinsom/Bank-backend/internal/service/card"
	creditService "github.com/kVinsom/Bank-backend/internal/service/credit"
	currencyService "github.com/kVinsom/Bank-backend/internal/service/currency"
	depositService "github.com/kVinsom/Bank-backend/internal/service/deposit"
	exchangeRateService "github.com/kVinsom/Bank-backend/internal/service/exhange_rate"
	userService "github.com/kVinsom/Bank-backend/internal/service/user"
)

// Services is the application service container consumed by the delivery layer.
type Services struct {
	Account      *accountService.AccountService
	Card         *cardService.CardService
	Credit       *creditService.CreditService
	Currency     *currencyService.CurrencyService
	ExchangeRate *exchangeRateService.ExchangeRateService
	Deposit      *depositService.DepositService
	User         *userService.UserService
}

// InitServices wires concrete services to the provided upstream clients.
func InitServices(
	cfg *configs.Config,
	repositories *grpcClient.RepositoryClients,
	exchangeRateClient exchangeRateProto.RankingRepositoryClient,
) *Services {
	return &Services{
		Account: accountService.NewAccountService(
			repositories.Account,
			repositories.User,
			repositories.Currency,
		),
		Card: cardService.NewCardService(
			repositories.Card,
			repositories.User,
			repositories.Account,
			cfg.CardConfig.BIN,
		),
		Credit: creditService.NewCreditService(
			repositories.Credit,
			repositories.User,
			repositories.Currency,
		),
		Currency: currencyService.NewCurrencyService(repositories.Currency),
		ExchangeRate: exchangeRateService.NewExchangeRateService(
			exchangeRateClient,
		),
		Deposit: depositService.NewDepositService(
			repositories.Deposit,
			repositories.User,
			repositories.Currency,
		),
		User: userService.NewUserService(repositories.User),
	}
}
