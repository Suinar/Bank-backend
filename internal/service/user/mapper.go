package user

import (
	log "github.com/kVinsom/Bank-backend/internal/logging/service/user"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// UserToCore maps a repository user into the domain model.
func UserToCore(user *userRepository.User) *core.User {
	const operation = "user_to_core"
	defer log.MappingStarted(operation)()
	if user == nil {
		log.NilInput(operation)
		return nil
	}
	return &core.User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		MiddleName:   user.MiddleName,
		LastName:     user.LastName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		PasswordHash: user.PasswordHash,
	}
}

// UserToProto maps a domain user into the repository contract.
func UserToProto(user *core.User) *userRepository.User {
	const operation = "user_to_proto"
	defer log.MappingStarted(operation)()
	return &userRepository.User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		MiddleName:   user.MiddleName,
		LastName:     user.LastName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		PasswordHash: user.PasswordHash,
	}
}

// UsersToCore maps repository users while preserving their order.
func UsersToCore(users []*userRepository.User) []core.User {
	const operation = "users_to_core"
	defer log.MappingStarted(operation)()
	result := make([]core.User, 0, len(users))
	for _, user := range users {
		if converted := UserToCore(user); converted != nil {
			result = append(result, *converted)
		}
	}
	return result
}
