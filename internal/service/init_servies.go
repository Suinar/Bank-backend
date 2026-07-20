package service

import (
	accountservice "github.com/kVinsom/Bank-backend/internal/service/account"
	cardservice "github.com/kVinsom/Bank-backend/internal/service/card"
	creditservice "github.com/kVinsom/Bank-backend/internal/service/credit"
	currencyservice "github.com/kVinsom/Bank-backend/internal/service/currency"
	depositservice "github.com/kVinsom/Bank-backend/internal/service/deposit"
	userservice "github.com/kVinsom/Bank-backend/internal/service/user"
)

type Services struct {
	Account      accountservice.IAccountService
	Card         cardservice.ICardService
	Credit       creditservice.ICreditService
	Currency     currencyservice.ICurrencyService
	Convert      currencyservice.IConvertService
	ExchangeRate currencyservice.IExchangeRateService
	Deposit      depositservice.IDepositService
	User         userservice.IUserService
}

func InitServices() *Services {
	return &Services{}
}
