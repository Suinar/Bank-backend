package credit

import (
	"context"

	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/credit.go -package=mocks

// ICreditService defines credit operations required by HTTP delivery.
type ICreditService interface {
	GetAll(context.Context) ([]core.Credit, error)
	GetByUser(context.Context, int64) ([]core.Credit, error)
	GetById(context.Context, int64) (*core.Credit, error)
	Create(context.Context, *core.CreditCreateInput) (*core.Credit, error)
	Repay(context.Context, int64, int) error
	Delete(context.Context, int64) error
}
