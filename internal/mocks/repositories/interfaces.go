package repositories

import (
	accountPr "github.com/kVinsom/Bank-proto/repository/account"
	cardPr "github.com/kVinsom/Bank-proto/repository/card"
	creditPr "github.com/kVinsom/Bank-proto/repository/credit"
	currencyPr "github.com/kVinsom/Bank-proto/repository/currency"
	depositPr "github.com/kVinsom/Bank-proto/repository/deposit"
	userPr "github.com/kVinsom/Bank-proto/repository/user"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -destination=account.go -package=repositories github.com/kVinsom/Bank-backend/internal/mocks/repositories AccountRepositoryClient
//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -destination=card.go -package=repositories github.com/kVinsom/Bank-backend/internal/mocks/repositories CardRepositoryClient
//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -destination=credit.go -package=repositories github.com/kVinsom/Bank-backend/internal/mocks/repositories CreditRepositoryClient
//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -destination=currency.go -package=repositories github.com/kVinsom/Bank-backend/internal/mocks/repositories CurrencyRepositoryClient
//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -destination=deposit.go -package=repositories github.com/kVinsom/Bank-backend/internal/mocks/repositories DepositRepositoryClient
//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -destination=user.go -package=repositories github.com/kVinsom/Bank-backend/internal/mocks/repositories UserRepositoryClient

type AccountRepositoryClient interface {
	accountPr.AccountRepositoryClient
}

type CardRepositoryClient interface {
	cardPr.CardRepositoryClient
}

type CreditRepositoryClient interface {
	creditPr.CreditRepositoryClient
}

type CurrencyRepositoryClient interface {
	currencyPr.CurrencyRepositoryClient
}

type DepositRepositoryClient interface {
	depositPr.DepositRepositoryClient
}

type UserRepositoryClient interface {
	userPr.UserRepositoryClient
}
