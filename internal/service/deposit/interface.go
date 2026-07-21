package deposit

import (
	"context"

	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/deposit.go -package=mocks

// IDepositService defines deposit operations required by HTTP delivery.
type IDepositService interface {
	GetAll(context.Context) ([]core.Deposit, error)
	GetByUser(context.Context, int64) ([]core.Deposit, error)
	GetById(context.Context, int64) (*core.Deposit, error)
	Create(context.Context, *core.DepositCreateInput) (*core.Deposit, error)
	Replenish(context.Context, int64, int) error
	Delete(context.Context, int64) error
}
