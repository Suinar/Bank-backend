package card

import (
	"context"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type ICardService interface {
	GetAll(context.Context) ([]core.Card, error)
	GetByUser(context.Context, int64) ([]core.Card, error)
	GetById(context.Context, int64) (*core.Card, error)
	GetByNumber(context.Context, string) (*core.Card, error)
	Blocking(context.Context, int64) error
	Create(context.Context, *core.CardCreateInput) (*core.Card, error)
	Delete(context.Context, int64) error
}
