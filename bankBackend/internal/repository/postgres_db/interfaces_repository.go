package core

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type IUserRepository interface {
	GetAll(ctx context.Context) ([]core.User, error)
	GetById(ctx context.Context, id int64) (*core.User, error)
	GetByEmail(ctx context.Context, email string) (*core.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error)
	Create(ctx context.Context, input *core.User) (*core.User, error)
	ChangePassword(ctx context.Context, id int64, newPassword string) error
	Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error)
	Delete(ctx context.Context, id int64) error
}

type IAccountRepository interface {
	GetAll(ctx context.Context) ([]core.Account, error)
	GetByUser(ctx context.Context, userId int64) ([]core.Account, error)
	GetById(ctx context.Context, id int64) (*core.Account, error)
	Create(ctx context.Context, input *core.Account) (*core.Account, error)
	Blocking(ctx context.Context, id int64) (core.Account, error)
	Close(ctx context.Context, id int64) (core.Account, error)
	Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type ICardRepository interface {
	GetAll(ctx context.Context) ([]core.Card, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Card, error)
	GetById(ctx context.Context, id int64) (*core.Card, error)
	GetByNumber(ctx context.Context, number string) (*core.Card, error)
	Blocking(ctx context.Context, id int64) (core.Card, error)
	Create(ctx context.Context, input *core.Card) (*core.Card, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type ICreditRepository interface {
	GetAll(ctx context.Context) ([]core.Credit, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Credit, error)
	GetById(ctx context.Context, id int64) (*core.Credit, error)
	Create(ctx context.Context, input *core.Credit) (*core.Credit, error)
	Repay(ctx context.Context, id int64, amount int) (*core.Credit, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type IDepositRepository interface {
	GetAll(ctx context.Context) ([]core.Deposit, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Deposit, error)
	GetById(ctx context.Context, id int64) (*core.Deposit, error)
	Create(ctx context.Context, input *core.Deposit) (*core.Deposit, error)
	Replenish(ctx context.Context, id int64, amount int) (*core.Deposit, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type ICurrencyRepository interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	GetById(ctx context.Context, id int64) (*core.Currency, error)
	GetByIso(ctx context.Context, isoCode string) (*core.Currency, error)
	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	Create(ctx context.Context, input *core.Currency) (*core.Currency, error)
	Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error)
	Delete(ctx context.Context, id int64) error
}
