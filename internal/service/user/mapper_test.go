package user

import (
	"reflect"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
)

func TestUserMappers(t *testing.T) {
	t.Parallel()
	coreValue := fixture.UserCore()
	if got := UserToCore(fixture.UserProto()); !reflect.DeepEqual(got, &coreValue) {
		t.Fatalf("UserToCore: want %#v, got %#v", &coreValue, got)
	}
	if got := UserToProto(&coreValue); !reflect.DeepEqual(got, fixture.UserProto()) {
		t.Fatalf("UserToProto: want %#v, got %#v", fixture.UserProto(), got)
	}
	if got := UserToCore(nil); got != nil {
		t.Fatalf("UserToCore(nil): want nil, got %#v", got)
	}
	if got := UsersToCore([]*userRepository.User{fixture.UserProto(), nil}); !reflect.DeepEqual(got, fixture.UserListCore()) {
		t.Fatalf("UsersToCore: want %#v, got %#v", fixture.UserListCore(), got)
	}
}
