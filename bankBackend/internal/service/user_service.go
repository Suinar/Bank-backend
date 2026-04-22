package core

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAll(ctx context.Context) ([]core.User, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	return users, nil
}

func (s *UserService) GetById(ctx context.Context, id int64) (*core.User, error) {
	user, err := s.userRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*core.User, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return user, nil
}

func (s *UserService) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error) {
	user, err := s.userRepository.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return user, nil
}

func (s *UserService) Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error) {
	if input == nil || input.FirstName == "" || input.LastName == "" ||
		input.Email == "" || input.PasswordHash == "" || input.PhoneNumber == "" {
		return nil, core.BadRequest
	}

	if len(input.FirstName) == 0 || len(input.FirstName) > 50 {
		return nil, core.BadRequest
	}

	if input.MiddleName != nil {
		if len(*input.MiddleName) == 0 || len(*input.MiddleName) > 50 {
			return nil, core.BadRequest
		}
	}

	if len(input.LastName) == 0 || len(input.LastName) > 50 {
		return nil, core.BadRequest
	}

	user := &core.User{
		FirstName:    input.FirstName,
		MiddleName:   input.MiddleName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		PasswordHash: input.PasswordHash,
	}

	created, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, core.InternalServerError
	}

	// todo: notification

	return created, nil
}

func (s *UserService) Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error) {
	if input == nil {
		return nil, core.BadRequest
	}

	if input.FirstName != nil {
		if len(*input.FirstName) == 0 || len(*input.FirstName) > 50 {
			return nil, core.BadRequest
		}
	}

	if input.MiddleName != nil {
		if len(*input.MiddleName) == 0 || len(*input.MiddleName) > 50 {
			return nil, core.BadRequest
		}
	}

	if input.LastName != nil {
		if len(*input.LastName) == 0 || len(*input.LastName) > 50 {
			return nil, core.BadRequest
		}
	}

	updated, err := s.userRepository.Update(ctx, id, input)
	if err != nil {
		if errors.Is(err, core.BadRequest) {
			return nil, core.BadRequest
		}

		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return updated, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id int64, newPassword string) error {
	if len(newPassword) < 6 || len(newPassword) > 20 {
		return core.BadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return core.InternalServerError
	}

	err = s.userRepository.ChangePassword(ctx, id, string(hashedPassword))
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// todo: notification

	return nil
}
