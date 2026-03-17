package core

import (
	"context"
	"time"

	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type IUserRepository interface {
	GetAll(ctx context.Context) ([]core.User, error)
	GetById(ctx context.Context, id uint64) (*core.User, error)
	GetByEmail(ctx context.Context, email string) (*core.User, error)
	GetByPhoneNumber(ctx context.Context, phone string) (*core.User, error)
	Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error)
	ChangePassword(ctx context.Context, id uint64, newPassword string) (*core.User, error)
	Update(ctx context.Context, input *core.UserUpdateInput) (*core.User, error)
	Delete(ctx context.Context, id uint64) error
}

type IAccountRepository interface {
	GetAll(ctx context.Context) ([]core.Account, error)
	GetByUser(ctx context.Context, idUser uint64) (*core.Account, error)
	GetById(ctx context.Context, id uint64) (*core.Account, error)
	GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Account, error)
	Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error)
	Update(ctx context.Context, input *core.AccountUpdateInput) (*core.Account, error)
	Delete(ctx context.Context, id uint64) error
}

type ICardRepository interface {
	GetAll(ctx context.Context) ([]core.Card, error)
	GetByUser(ctx context.Context, idUser uint64) (*core.Card, error)
	GetById(ctx context.Context, id uint64) (*core.Card, error)
	GetByNumber(ctx context.Context, number string) (*core.Account, error)
	GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Account, error)
	Blocking(ctx context.Context, id uint64) error
	Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error)
	Delete(ctx context.Context, id uint64) error
}

type ICreditRepository interface {
	GetAll(ctx context.Context) ([]core.Credit, error)
	GetByUser(ctx context.Context, idUser uint64) (*core.Credit, error)
	GetById(ctx context.Context, id uint64) (*core.Credit, error)
	GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Credit, error)
	GetByRepayTime(ctx context.Context, repayTime time.Time) (*core.Credit, error)
	Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error)
	Replay(ctx context.Context, id uint64) (*core.Credit, error)
	Delete(ctx context.Context, id uint64) error
}

type IDepositRepository interface {
	GetAll(ctx context.Context) ([]core.Deposit, error)
	GetByUser(ctx context.Context, idUser uint64) (*core.Deposit, error)
	GetById(ctx context.Context, id uint64) (*core.Deposit, error)
	GetByCreateTime(ctx context.Context, createTime time.Time) (*core.Deposit, error)
	GetByCompletionTime(ctx context.Context, completeTime time.Time) (*core.Deposit, error)
	Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error)
	Repay(ctx context.Context, id uint64, amount int) error
	Delete(ctx context.Context, id uint64) error
}

type ICurrencyRepository interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	GetByUser(ctx context.Context, idUser uint64) (*core.Currency, error)
	GetById(ctx context.Context, id uint64) (*core.Currency, error)
	GetByIsoCode(ctx context.Context, isoCode string) (*core.Currency, error)
	GetByNumberCode(ctx context.Context, numberCode string) (*core.Currency, error)
	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error)
	Update(ctx context.Context, input *core.CurrencyUpdateInput) (*core.Currency, error)
	Delete(ctx context.Context, id uint64) error
}
