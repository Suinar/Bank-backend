package credit

import (
	"context"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type ICreditService interface {
	GetAll(context.Context) ([]core.Credit, error)
	GetByUser(context.Context, int64) ([]core.Credit, error)
	GetById(context.Context, int64) (*core.Credit, error)
	Create(context.Context, *core.CreditCreateInput) (*core.Credit, error)
	Repay(context.Context, int64, int) error
	Delete(context.Context, int64) error
}
