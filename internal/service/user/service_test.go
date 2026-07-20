package user

import "testing"

func TestUserServiceImplementsInterface(t *testing.T) {
	var service IUserService = &UserService{}
	if service == nil {
		t.Fatal("expected service implementation")
	}
}
