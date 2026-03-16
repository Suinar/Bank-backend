package core

import (
	"context"

	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]core.User, error) {}

func (r *UserRepository) GetById(ctx context.Context, id string) (*core.User, error) {}

func (r *UserRepository) GetByEmail(email string) (*core.User, error) {}

func (r *UserRepository) GetByPhoneNumber(phone string) (*core.User, error) {}

func (r *UserRepository) Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error) {}

func (r *UserRepository) ChangePassword(ctx context.Context, id string, newPassword string) (*core.User, error) {}

func (r *UserRepository) Update(ctx context.Context, input *core.UserUpdateInput) (*core.User, error) {}

func (r *UserRepository) Delete(ctx context.Context, id string) error {}