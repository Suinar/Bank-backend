package handler

import (
	accountHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/account"
	cardHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/card"
	creditHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/credit"
	currencyHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/currency"
	depositHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/deposit"
	userHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/user"
	service "github.com/kVinsom/Bank-backend/internal/service"
)

type Handlers struct {
	Account  *accountHandler.AccountHandler
	Card     *cardHandler.CardHandler
	Credit   *creditHandler.CreditHandler
	Currency *currencyHandler.CurrencyHandler
	Deposit  *depositHandler.DepositHandler
	User     *userHandler.UserHandler
}

func InitHandlers(services service.Services) *Handlers {
	return &Handlers{
		Account:  accountHandler.NewAccountHandler(services.Account),
		Card:     cardHandler.NewCardHandler(services.Card),
		Credit:   creditHandler.NewCreditHandler(services.Credit),
		Currency: currencyHandler.NewCurrencyHandler(services.Currency, services.Convert, services.ExchangeRate),
		Deposit:  depositHandler.NewDepositHandler(services.Deposit),
		User:     userHandler.NewUserHandler(services.User),
	}
}
