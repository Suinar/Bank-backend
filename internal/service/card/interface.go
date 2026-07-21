package card

import (
	"context"

	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/card.go -package=mocks

// ICardService defines card operations required by HTTP delivery.
type ICardService interface {
	GetAll(context.Context) ([]core.Card, error)
	GetByUser(context.Context, int64) ([]core.Card, error)
	GetById(context.Context, int64) (*core.Card, error)
	GetByNumber(context.Context, string) (*core.Card, error)
	Blocking(context.Context, int64) error
	Create(context.Context, *core.CardCreateInput) (*core.Card, error)
	Delete(context.Context, int64) error
}
