package user

import (
	"context"
	log "github.com/kVinsom/Bank-backend/internal/logging/service/user"

	serviceCommon "github.com/kVinsom/Bank-backend/internal/service/common"
	"github.com/kVinsom/Bank-proto/repository/common"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"golang.org/x/crypto/bcrypt"
)

// UserService implements user management use cases.
type UserService struct {
	userRepository userRepository.UserRepositoryClient
}

// NewUserService creates a user service backed by the user repository.
func NewUserService(
	userRepository userRepository.UserRepositoryClient) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]core.User, error) {
	const operation = "get_all"
	defer log.OperationStarted(operation)()
	response, err := s.userRepository.GetAll(ctx, &common.Empty{})
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return UsersToCore(response.Users), nil
}

func (s *UserService) GetById(ctx context.Context, id int64) (*core.User, error) {
	const operation = "get_by_id"
	defer log.OperationStarted(operation)()
	user, err := s.userRepository.GetById(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if serviceCommon.IsError(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return UserToCore(user), nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*core.User, error) {
	const operation = "get_by_email"
	defer log.OperationStarted(operation)()
	user, err := s.userRepository.GetByEmail(ctx, &userRepository.EmailRequest{Email: email})
	if err != nil {
		if serviceCommon.IsError(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return UserToCore(user), nil
}

func (s *UserService) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error) {
	const operation = "get_by_phone_number"
	defer log.OperationStarted(operation)()
	user, err := s.userRepository.GetByPhoneNumber(ctx, &userRepository.PhoneNumberRequest{PhoneNumber: phoneNumber})
	if err != nil {
		if serviceCommon.IsError(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return UserToCore(user), nil
}

func (s *UserService) Create(ctx context.Context, input *core.UserCreateInput) (*core.User, error) {
	const operation = "create"
	defer log.OperationStarted(operation)()
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

	created, err := s.userRepository.Create(ctx, UserToProto(user))
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return UserToCore(created), nil
}

func (s *UserService) Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error) {
	const operation = "update"
	defer log.OperationStarted(operation)()
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

	updated, err := s.userRepository.Update(ctx, &userRepository.UpdateUserRequest{
		Id: id,
		Input: &userRepository.UserUpdateInput{
			FirstName:  input.FirstName,
			MiddleName: input.MiddleName,
			LastName:   input.LastName,
		},
	})
	if err != nil {
		if serviceCommon.IsError(err, coreErrors.BadRequest) {
			return nil, coreErrors.BadRequest
		}

		if serviceCommon.IsError(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return UserToCore(updated), nil
}

func (s *UserService) ChangePassword(ctx context.Context, id int64, newPassword string) error {
	const operation = "change_password"
	defer log.OperationStarted(operation)()
	if len(newPassword) < 6 || len(newPassword) > 20 {
		return coreErrors.BadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return coreErrors.InternalServerError
	}

	if s.userRepository == nil {
		return coreErrors.InternalServerError
	}

	_, err = s.userRepository.ChangePassword(ctx, &userRepository.ChangePasswordRequest{
		Id:          id,
		NewPassword: string(hashedPassword),
	})
	if err != nil {
		if serviceCommon.IsError(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	const operation = "delete"
	defer log.OperationStarted(operation)()
	_, err := s.userRepository.Delete(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if serviceCommon.IsError(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}
