package user

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

var userRepositoryError = errors.New("repository failure")

func TestUserService_Create_InvalidInput(t *testing.T) {
	t.Parallel()
	valid := fixture.UserCreateInputCore()
	tests := []struct {
		name     string
		mutate   func(*core.UserCreateInput)
		nilInput bool
	}{
		{name: "nil", nilInput: true},
		{name: "empty first name", mutate: func(v *core.UserCreateInput) { v.FirstName = "" }},
		{name: "long first name", mutate: func(v *core.UserCreateInput) { v.FirstName = strings.Repeat("a", 51) }},
		{name: "empty middle name", mutate: func(v *core.UserCreateInput) { v.MiddleName = fixture.StringPointer("") }},
		{name: "long middle name", mutate: func(v *core.UserCreateInput) { v.MiddleName = fixture.StringPointer(strings.Repeat("a", 51)) }},
		{name: "empty last name", mutate: func(v *core.UserCreateInput) { v.LastName = "" }},
		{name: "long last name", mutate: func(v *core.UserCreateInput) { v.LastName = strings.Repeat("a", 51) }},
		{name: "empty email", mutate: func(v *core.UserCreateInput) { v.Email = "" }},
		{name: "empty phone", mutate: func(v *core.UserCreateInput) { v.PhoneNumber = "" }},
		{name: "empty password", mutate: func(v *core.UserCreateInput) { v.PasswordHash = "" }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := valid
			if tt.nilInput {
				assertUserError(t, func() (any, error) { return (&UserService{}).Create(context.Background(), nil) }, coreErrors.BadRequest)
				return
			}
			tt.mutate(&input)
			assertUserError(t, func() (any, error) { return (&UserService{}).Create(context.Background(), &input) }, coreErrors.BadRequest)
		})
	}
}

func TestUserService_RepositoryErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.UserRepositoryMocks)
		invoke func(*UserService) (any, error)
		want   error
	}{
		{name: "get all", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, userRepositoryError)
		}, invoke: func(s *UserService) (any, error) { return s.GetAll(context.Background()) }, want: coreErrors.InternalServerError},
		{name: "get by id not found", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *UserService) (any, error) { return s.GetById(context.Background(), test.UserId) }, want: coreErrors.NotFound},
		{name: "get by id failure", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, userRepositoryError)
		}, invoke: func(s *UserService) (any, error) { return s.GetById(context.Background(), test.UserId) }, want: coreErrors.InternalServerError},
		{name: "get by email", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(nil, userRepositoryError)
		}, invoke: func(s *UserService) (any, error) { return s.GetByEmail(context.Background(), test.UserEmail) }, want: coreErrors.InternalServerError},
		{name: "get by email not found", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *UserService) (any, error) { return s.GetByEmail(context.Background(), test.UserEmail) }, want: coreErrors.NotFound},
		{name: "get by phone not found", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetByPhoneNumber(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *UserService) (any, error) {
			return s.GetByPhoneNumber(context.Background(), test.UserPhoneNumber)
		}, want: coreErrors.NotFound},
		{name: "get by phone failure", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetByPhoneNumber(gomock.Any(), gomock.Any()).Return(nil, userRepositoryError)
		}, invoke: func(s *UserService) (any, error) {
			return s.GetByPhoneNumber(context.Background(), test.UserPhoneNumber)
		}, want: coreErrors.InternalServerError},
		{name: "create", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, userRepositoryError)
		}, invoke: func(s *UserService) (any, error) {
			input := fixture.UserCreateInputCore()
			return s.Create(context.Background(), &input)
		}, want: coreErrors.InternalServerError},
		{name: "update not found", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *UserService) (any, error) {
			input := fixture.UserUpdateInputCore()
			return s.Update(context.Background(), test.UserId, &input)
		}, want: coreErrors.NotFound},
		{name: "update bad request", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, coreErrors.BadRequest)
		}, invoke: func(s *UserService) (any, error) {
			input := fixture.UserUpdateInputCore()
			return s.Update(context.Background(), test.UserId, &input)
		}, want: coreErrors.BadRequest},
		{name: "update failure", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, userRepositoryError)
		}, invoke: func(s *UserService) (any, error) {
			input := fixture.UserUpdateInputCore()
			return s.Update(context.Background(), test.UserId, &input)
		}, want: coreErrors.InternalServerError},
		{name: "delete", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, userRepositoryError)
		}, invoke: func(s *UserService) (any, error) { return nil, s.Delete(context.Background(), test.UserId) }, want: coreErrors.InternalServerError},
		{name: "delete not found", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *UserService) (any, error) { return nil, s.Delete(context.Background(), test.UserId) }, want: coreErrors.NotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newUserServiceSUT(t)
			tt.expect(mocks)
			assertUserError(t, func() (any, error) { return tt.invoke(sut) }, tt.want)
		})
	}
}

func TestUserService_UpdateAndPassword_InvalidInput(t *testing.T) {
	t.Parallel()
	assertUserError(t, func() (any, error) { return (&UserService{}).Update(context.Background(), test.UserId, nil) }, coreErrors.BadRequest)
	empty := ""
	long := strings.Repeat("a", 51)
	for _, input := range []*core.UserUpdateInput{
		{FirstName: &empty}, {FirstName: &long},
		{MiddleName: &empty}, {MiddleName: &long},
		{LastName: &empty}, {LastName: &long},
	} {
		input := input
		assertUserError(t, func() (any, error) {
			return (&UserService{}).Update(context.Background(), test.UserId, input)
		}, coreErrors.BadRequest)
	}
	for _, password := range []string{"", "12345", strings.Repeat("a", 21)} {
		password := password
		t.Run("password length", func(t *testing.T) {
			assertUserError(t, func() (any, error) {
				return nil, (&UserService{}).ChangePassword(context.Background(), test.UserId, password)
			}, coreErrors.BadRequest)
		})
	}
}

func TestUserService_ChangePassword_Errors(t *testing.T) {
	t.Parallel()

	assertUserError(t, func() (any, error) {
		return nil, NewUserService(nil).ChangePassword(context.Background(), test.UserId, "secret12")
	}, coreErrors.InternalServerError)

	for _, tt := range []struct {
		name       string
		repository error
		want       error
	}{
		{name: "not found", repository: coreErrors.NotFound, want: coreErrors.NotFound},
		{name: "failure", repository: userRepositoryError, want: coreErrors.InternalServerError},
	} {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newUserServiceSUT(t)
			mocks.User.EXPECT().ChangePassword(gomock.Any(), gomock.Any()).Return(nil, tt.repository)
			assertUserError(t, func() (any, error) {
				return nil, sut.ChangePassword(context.Background(), test.UserId, "secret12")
			}, tt.want)
		})
	}
}

func assertUserError(t *testing.T, invoke func() (any, error), want error) {
	t.Helper()
	got, err := invoke()
	if got != nil && !reflect.ValueOf(got).IsNil() {
		t.Fatalf("result: want nil, got %#v", got)
	}
	if !errors.Is(err, want) {
		t.Fatalf("error: want %v, got %v", want, err)
	}
}
