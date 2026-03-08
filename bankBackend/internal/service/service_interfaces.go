package core

import (
	"context"
	"time"

	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type IUserService interface {
	GetAll(ctx context.Context) ([]core.User, error)
	GetById(ctx context.Context, id string) (*core.User, error)
	GetByEmail(ctx context.Context, email string) (*core.User, error)
	GetByPhoneNumber(ctx context.Context, phone string) (*core.User, error)
	GetMe(ctx context.Context, id string) (core.User, error)
	Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error)
	Update(ctx context.Context, id string, input *core.UserUpdateInput) (*core.User, error)
	ChangePassword(ctx context.Context, id string, newPassword string) (*core.User, error)
	Delete(ctx context.Context, id string) error
}

type IAccountService interface {
	GetAll(ctx context.Context) ([]core.Account, error)
	GetByUser(ctx context.Context, userId string) ([]core.Account, error)
	GetById(ctx context.Context, id string) (*core.Account, error)
	GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Account, error)
	Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error)
	Update(ctx context.Context, id string, input *core.AccountUpdateInput) (*core.Account, error)
	Delete(ctx context.Context, id string) error
}

type ICardService interface {
	GetAll(ctx context.Context) ([]core.Card, error)
	GetByUser(ctx context.Context, userId string) ([]core.Card, error)
	GetById(ctx context.Context, id string) (*core.Card, error)
	GetByNumber(ctx context.Context, number string) (*core.Card, error)
	GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Card, error)
	Blocked(ctx context.Context, id string) error
	Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error)
	Delete(ctx context.Context, id string) error
}

type ICreditService interface {
	GetAll(ctx context.Context) ([]core.Credit, error)
	GetByUser(ctx context.Context, userId string) ([]core.Credit, error)
	GetById(ctx context.Context, id string) (*core.Credit, error)
	GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Credit, error)
	GetByRepayTime(ctx context.Context, repayTime time.Time) (*core.Credit, error)
	Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error)
	Repay(ctx context.Context, id string, amount int) error
	Delete(ctx context.Context, id string) error
}

type IDepositService interface {
	GetAll(ctx context.Context) ([]core.Deposit, error)
	GetByUser(ctx context.Context, userID string) ([]core.Deposit, error)
	GetById(ctx context.Context, depositId string) (*core.Deposit, error)
	GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Deposit, error)
	GetByCompletionTime(ctx context.Context, createTime time.Time) (*core.Deposit, error)
	Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error)
	Repay(ctx context.Context, id string, amount int) error
	Delete(ctx context.Context, id string) error
}

type ICurrencyService interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	GetById(ctx context.Context, id string) (*core.Currency, error)
	GetByIsoCod(ctx context.Context, isoCode string) (*core.Currency, error)
	GetByNumberCod(ctx context.Context, numberCode string) (*core.Currency, error)
	GetBySymbol(ctx context.Context, symbol string) (*core.Currency, error)
	Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error)
	Update(ctx context.Context, id string, input *core.CurrencyUpdateInput) (*core.Currency, error)
	Delete(ctx context.Context, id string) error
}

type IExchangeRateService interface {
	GetAll(ctx context.Context) ([]core.ExchangeRate, error)
	GetForCurrencies(ctx context.Context, currencyIdFrom string, currencyIdTo string) (*core.ExchangeRate, error)
}
