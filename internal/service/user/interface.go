package user

import (
	"context"

	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/user.go -package=mocks

// IUserService defines user operations required by HTTP delivery.
type IUserService interface {
	GetAll(context.Context) ([]core.User, error)
	GetById(context.Context, int64) (*core.User, error)
	GetByEmail(context.Context, string) (*core.User, error)
	GetByPhoneNumber(context.Context, string) (*core.User, error)
	Create(context.Context, *core.UserCreateInput) (*core.User, error)
	Update(context.Context, int64, *core.UserUpdateInput) (*core.User, error)
	ChangePassword(context.Context, int64, string) error
	Delete(context.Context, int64) error
}
