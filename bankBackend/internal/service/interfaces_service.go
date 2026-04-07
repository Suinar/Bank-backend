package core

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type IUserService interface {
	GetAll(ctx context.Context) ([]core.User, error)
	GetById(ctx context.Context, id uint64) (*core.User, error)
	GetByEmail(ctx context.Context, email string) (*core.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error)
	Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error)
	Update(ctx context.Context, id uint64, input *core.UserUpdateInput) (*core.User, error)
	ChangePassword(ctx context.Context, id uint64, newPassword string) error
	Delete(ctx context.Context, id uint64) error
}

type IAccountService interface {
	GetAll(ctx context.Context) ([]core.Account, error)
	GetByUser(ctx context.Context, userId uint64) ([]core.Account, error)
	GetById(ctx context.Context, id uint64) (*core.Account, error)
	Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error)
	Blocking(ctx context.Context, id uint64) error
	Close(ctx context.Context, id uint64) error
	Update(ctx context.Context, id uint64, input *core.AccountUpdateInput) (*core.Account, error)
	Delete(ctx context.Context, id uint64) error
}

type ICardService interface {
	GetAll(ctx context.Context) ([]core.Card, error)
	GetByUser(ctx context.Context, userId uint64) ([]core.Card, error)
	GetById(ctx context.Context, id uint64) (*core.Card, error)
	GetByNumber(ctx context.Context, number string) (*core.Card, error)
	Blocking(ctx context.Context, id uint64) error
	Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error)
	Delete(ctx context.Context, id uint64) error
}

type ICreditService interface {
	GetAll(ctx context.Context) ([]core.Credit, error)
	GetByUser(ctx context.Context, userId uint64) ([]core.Credit, error)
	GetById(ctx context.Context, id uint64) (*core.Credit, error)
	Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error)
	Repay(ctx context.Context, id uint64, amount int) error
	Delete(ctx context.Context, id uint64) error
}

type IDepositService interface {
	GetAll(ctx context.Context) ([]core.Deposit, error)
	GetByUser(ctx context.Context, userId uint64) ([]core.Deposit, error)
	GetById(ctx context.Context, id uint64) (*core.Deposit, error)
	Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error)
	Replenish(ctx context.Context, id uint64, amount int) error
	Delete(ctx context.Context, id uint64) error
}

type ICurrencyService interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	GetById(ctx context.Context, id uint64) (*core.Currency, error)
	GetByIsoCod(ctx context.Context, isoCode string) (*core.Currency, error)
	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	Convert(ctx context.Context, currencyIdFrom uint64, amount int, currencyIdTo uint64) (float64, error)
	Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error)
	Update(ctx context.Context, id uint64, input *core.CurrencyUpdateInput) (*core.Currency, error)
	Delete(ctx context.Context, id uint64) error
	GetAllRanking(ctx context.Context, currencyIdFrom uint64) ([]core.ExchangeRate, error)
	GetRelativeRanking(ctx context.Context, currencyIdFrom uint64, currencyIdTo uint64) (*core.ExchangeRate, error)
}
