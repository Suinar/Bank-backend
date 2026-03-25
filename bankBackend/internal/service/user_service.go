package core

import (
	"context"
	"errors"

	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	"github.com/Suinar/Bank-backend/bankBackend/internal/core"
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetAll(ctx context.Context) ([]core.User, error) {}

func (s *UserService) GetById(ctx context.Context, id uint64) (*core.User, error) {
	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*core.User, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return user, nil
}

func (s *UserService) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error) {
	user, err := s.repository.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return user, nil
}

func (s *UserService) Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error) {}

func (s *UserService) Update(ctx context.Context, id uint64, input *core.UserUpdateInput) (*core.User, error) {
}

func (s *UserService) ChangePassword(ctx context.Context, id uint64, newPassword string) (*core.User, error) {
}

func (s *UserService) Delete(ctx context.Context, id uint64) error {}
