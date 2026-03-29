package core

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetAll(ctx context.Context) ([]core.User, error) {
	users, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	return users, nil
}

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

func (s *UserService) Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error) {
	return nil, nil
}

func (s *UserService) Update(ctx context.Context, id uint64, input *core.UserUpdateInput) (*core.User, error) {
	return nil, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id uint64, newPassword string) error {
	if len(newPassword) < 6 || len(newPassword) > 100 {
		return core.BadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return core.InternalServerError
	}

	err = s.repository.ChangePassword(ctx, id, string(hashedPassword))
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id uint64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// todo: notification

	return nil
}
