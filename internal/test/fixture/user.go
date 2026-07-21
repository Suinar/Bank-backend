package fixture

import (
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// UserCore returns a valid user fixture.
func UserCore() core.User {
	return core.User{
		Id:           UserId,
		FirstName:    UserFirstName,
		MiddleName:   StringPointer(UserMiddleName),
		LastName:     UserLastName,
		Email:        UserEmail,
		PhoneNumber:  UserPhoneNumber,
		PasswordHash: UserPasswordHash,
	}
}

func UserListCore() []core.User { return []core.User{UserCore()} }

func UserProto() *userRepository.User {
	user := UserCore()
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

func UserListProto(users ...*userRepository.User) *userRepository.UserList {
	return &userRepository.UserList{
		Users: users,
	}
}

// UserCreateInputCore returns a valid user creation fixture.
func UserCreateInputCore() core.UserCreateInput {
	return core.UserCreateInput{
		FirstName:    UserFirstName,
		MiddleName:   StringPointer(UserMiddleName),
		LastName:     UserLastName,
		Email:        UserEmail,
		PhoneNumber:  UserPhoneNumber,
		PasswordHash: UserPasswordHash,
	}
}

// UserUpdateInputCore returns a valid user update fixture.
func UserUpdateInputCore() core.UserUpdateInput {
	return core.UserUpdateInput{
		FirstName:  StringPointer(UserFirstName),
		MiddleName: StringPointer(UserMiddleName),
		LastName:   StringPointer(UserLastName),
	}
}
