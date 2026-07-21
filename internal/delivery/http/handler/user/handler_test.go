package user

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kVinsom/Bank-backend/internal/mocks/services"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestUserHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	expected := fixture.UserListCore()

	service.EXPECT().GetAll(gomock.Any()).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestUserHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	expected := fixture.UserCore()

	service.EXPECT().GetById(gomock.Any(), fixture.UserId).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestUserHandler_GetByEmail_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	expected := fixture.UserCore()

	service.EXPECT().GetByEmail(gomock.Any(), fixture.UserEmail).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/email", nil)
	ctx.AddParam("email", fixture.UserEmail)

	sut.GetByEmail(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestUserHandler_GetByPhoneNumber_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	expected := fixture.UserCore()

	service.EXPECT().GetByPhoneNumber(gomock.Any(), fixture.UserPhoneNumber).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/phone", nil)
	ctx.AddParam("phone_number", fixture.UserPhoneNumber)

	sut.GetByPhoneNumber(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestUserHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	input := fixture.UserCreateInputCore()
	expected := fixture.UserCore()

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/users", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestUserHandler_UpdateById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	input := fixture.UserUpdateInputCore()
	expected := fixture.UserCore()

	service.EXPECT().Update(gomock.Any(), fixture.UserId, gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/users/1", input)
	ctx.AddParam("id", "1")

	sut.UpdateById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestUserHandler_ChangePasswordById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	service.EXPECT().ChangePassword(gomock.Any(), fixture.UserId, fixture.UserPasswordHash).Return(nil).Times(1)

	form := url.Values{"hash_password": {fixture.UserPasswordHash}}
	ctx, recorder := fixture.NewFormHTTPContext(t, http.MethodPatch, "/users/1/password", form)
	ctx.AddParam("id", "1")

	sut.ChangePasswordById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func TestUserHandler_DeleteById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	service.EXPECT().Delete(gomock.Any(), fixture.UserId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/users/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func NewUserSUT(t *testing.T) (*mocks.MockIUserService, *UserHandler) {
	t.Helper()
	return fixture.NewMockSUT(t, mocks.NewMockIUserService, func(service *mocks.MockIUserService) *UserHandler {
		return NewUserHandler(service)
	})
}
