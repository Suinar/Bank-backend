package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	"github.com/kVinsom/Bank-proto/repository/common"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
)

func TestUserService_Read_Success(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		expect func(*fixture.UserRepositoryMocks)
		invoke func(*UserService) (any, error)
		want   any
	}{
		{name: "get all", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetAll(gomock.Any(), &common.Empty{}).Return(fixture.UserListProto(fixture.UserProto()), nil)
		}, invoke: func(s *UserService) (any, error) { return s.GetAll(context.Background()) }, want: fixture.UserListCore()},
		{name: "get by id", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), &common.IdRequest{Id: fixture.UserId}).Return(fixture.UserProto(), nil)
		}, invoke: func(s *UserService) (any, error) { return s.GetById(context.Background(), fixture.UserId) }, want: userCorePointer()},
		{name: "get by email", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetByEmail(gomock.Any(), &userRepository.EmailRequest{Email: fixture.UserEmail}).Return(fixture.UserProto(), nil)
		}, invoke: func(s *UserService) (any, error) { return s.GetByEmail(context.Background(), fixture.UserEmail) }, want: userCorePointer()},
		{name: "get by phone", expect: func(m *fixture.UserRepositoryMocks) {
			m.User.EXPECT().GetByPhoneNumber(gomock.Any(), &userRepository.PhoneNumberRequest{PhoneNumber: fixture.UserPhoneNumber}).Return(fixture.UserProto(), nil)
		}, invoke: func(s *UserService) (any, error) {
			return s.GetByPhoneNumber(context.Background(), fixture.UserPhoneNumber)
		}, want: userCorePointer()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks, sut := newUserServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertUserResult(t, got, tt.want, err)
		})
	}
}

func TestUserService_Write_Success(t *testing.T) {
	t.Parallel()

	t.Run("create", func(t *testing.T) {
		mocks, sut := newUserServiceSUT(t)
		input := fixture.UserCreateInputCore()
		mocks.User.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
		got, err := sut.Create(context.Background(), &input)
		assertUserResult(t, got, userCorePointer(), err)
	})

	t.Run("update", func(t *testing.T) {
		mocks, sut := newUserServiceSUT(t)
		input := fixture.UserUpdateInputCore()
		mocks.User.EXPECT().Update(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
		got, err := sut.Update(context.Background(), fixture.UserId, &input)
		assertUserResult(t, got, userCorePointer(), err)
	})

	t.Run("change password", func(t *testing.T) {
		mocks, sut := newUserServiceSUT(t)
		mocks.User.EXPECT().ChangePassword(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
		if err := sut.ChangePassword(context.Background(), fixture.UserId, "secret12"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("delete", func(t *testing.T) {
		mocks, sut := newUserServiceSUT(t)
		mocks.User.EXPECT().Delete(gomock.Any(), &common.IdRequest{Id: fixture.UserId}).Return(&common.Empty{}, nil)
		if err := sut.Delete(context.Background(), fixture.UserId); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func newUserServiceSUT(t *testing.T) (*fixture.UserRepositoryMocks, *UserService) {
	t.Helper()
	mocks := fixture.NewUserRepositoryMocks(t)
	return mocks, NewUserService(mocks.User)
}

func userCorePointer() any { value := fixture.UserCore(); return &value }

func assertUserResult(t *testing.T, got, want any, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("result: want %#v, got %#v", want, got)
	}
}
