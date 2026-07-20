package user

import (
	"context"
	"errors"

	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository userRepository.UserRepositoryClient
}

func NewUserService(
	userRepository userRepository.UserRepositoryClient) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]core.User, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return users, nil
}

func (s *UserService) GetById(ctx context.Context, id int64) (*core.User, error) {
	user, err := s.userRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*core.User, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return user, nil
}

func (s *UserService) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error) {
	user, err := s.userRepository.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return user, nil
}

func (s *UserService) Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error) {
	if input == nil || input.FirstName == "" || input.LastName == "" ||
		input.Email == "" || input.PasswordHash == "" || input.PhoneNumber == "" {
		return nil, coreErrors.BadRequest
	}

	if len(input.FirstName) == 0 || len(input.FirstName) > 50 {
		return nil, coreErrors.BadRequest
	}

	if input.MiddleName != nil {
		if len(*input.MiddleName) == 0 || len(*input.MiddleName) > 50 {
			return nil, coreErrors.BadRequest
		}
	}

	if len(input.LastName) == 0 || len(input.LastName) > 50 {
		return nil, coreErrors.BadRequest
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
		return nil, coreErrors.InternalServerError
	}

	return created, nil
}

func (s *UserService) Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error) {
	if input == nil {
		return nil, coreErrors.BadRequest
	}

	if input.FirstName != nil {
		if len(*input.FirstName) == 0 || len(*input.FirstName) > 50 {
			return nil, coreErrors.BadRequest
		}
	}

	if input.MiddleName != nil {
		if len(*input.MiddleName) == 0 || len(*input.MiddleName) > 50 {
			return nil, coreErrors.BadRequest
		}
	}

	if input.LastName != nil {
		if len(*input.LastName) == 0 || len(*input.LastName) > 50 {
			return nil, coreErrors.BadRequest
		}
	}

	updated, err := s.userRepository.Update(ctx, id, input)
	if err != nil {
		if errors.Is(err, coreErrors.BadRequest) {
			return nil, coreErrors.BadRequest
		}

		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return updated, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id int64, newPassword string) error {
	if len(newPassword) < 6 || len(newPassword) > 20 {
		return coreErrors.BadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return coreErrors.InternalServerError
	}

	err = s.userRepository.ChangePassword(ctx, id, string(hashedPassword))
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}
