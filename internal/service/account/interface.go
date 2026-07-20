package account

import (
	"context"

	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type IAccountService interface {
	GetAll(context.Context) ([]core.Account, error)
	GetByUser(context.Context, int64) ([]core.Account, error)
	GetById(context.Context, int64) (*core.Account, error)
	Create(context.Context, *core.AccountCreateInput) (*core.Account, error)
	Blocking(context.Context, int64) error
	Close(context.Context, int64) error
	Update(context.Context, int64, *core.AccountUpdateInput) (*core.Account, error)
	Delete(context.Context, int64) error
}
