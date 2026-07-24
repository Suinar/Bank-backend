package user

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

var userHandlerTestError = errors.New("service failure")

func TestUserHandler_GetAll_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	service.EXPECT().GetAll(gomock.Any()).Return(nil, userHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, userHandlerTestError.Error())
}

func TestUserHandler_GetById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	service.EXPECT().GetById(gomock.Any(), test.UserId).Return(nil, userHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, userHandlerTestError.Error())
}

func TestUserHandler_GetByEmail_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	service.EXPECT().GetByEmail(gomock.Any(), test.UserEmail).Return(nil, userHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/email", nil)
	ctx.AddParam("email", test.UserEmail)

	sut.GetByEmail(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, userHandlerTestError.Error())
}

func TestUserHandler_GetByPhoneNumber_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	service.EXPECT().GetByPhoneNumber(gomock.Any(), test.UserPhoneNumber).Return(nil, userHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/phone", nil)
	ctx.AddParam("phone_number", test.UserPhoneNumber)

	sut.GetByPhoneNumber(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, userHandlerTestError.Error())
}

func TestUserHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	input := core.UserCreateInput{}

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(nil, userHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/users", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, userHandlerTestError.Error())
}

func TestUserHandler_UpdateById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	input := core.UserUpdateInput{}

	service.EXPECT().Update(gomock.Any(), test.UserId, gomock.Eq(&input)).Return(nil, userHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/users/1", input)
	ctx.AddParam("id", "1")

	sut.UpdateById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, userHandlerTestError.Error())
}

func TestUserHandler_ChangePasswordById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	service.EXPECT().ChangePassword(gomock.Any(), test.UserId, test.UserPasswordHash).Return(userHandlerTestError).Times(1)

	form := url.Values{"hash_password": {test.UserPasswordHash}}
	ctx, recorder := fixture.NewFormHTTPContext(t, http.MethodPatch, "/users/1/password", form)
	ctx.AddParam("id", "1")

	sut.ChangePasswordById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, userHandlerTestError.Error())
}

func TestUserHandler_DeleteById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewUserSUT(t)

	service.EXPECT().Delete(gomock.Any(), test.UserId).Return(userHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/users/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, userHandlerTestError.Error())
}

func TestUserHandler_GetById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewUserSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.GetById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
