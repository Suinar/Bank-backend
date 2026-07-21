package account

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kVinsom/Bank-backend/internal/mocks/services"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestAccountHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	expected := fixture.AccountListCore()

	service.EXPECT().GetAll(gomock.Any()).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/accounts", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestAccountHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	expected := fixture.AccountListCore()

	service.EXPECT().GetByUser(gomock.Any(), fixture.UserId).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1/accounts", nil)
	ctx.AddParam("user_id", strconv.FormatInt(fixture.UserId, 10))

	sut.GetByUser(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestAccountHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	expected := fixture.AccountCore()

	service.EXPECT().GetById(gomock.Any(), fixture.AccountId).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/accounts/1", nil)
	ctx.AddParam("id", strconv.FormatInt(fixture.AccountId, 10))

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestAccountHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	input := fixture.AccountCreateInputCore()
	expected := fixture.AccountCore()

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/accounts", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestAccountHandler_BlockingById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().Blocking(gomock.Any(), fixture.AccountId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/1/block", nil)
	ctx.AddParam("id", strconv.FormatInt(fixture.AccountId, 10))

	sut.BlockingById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func TestAccountHandler_CloseById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().Close(gomock.Any(), fixture.AccountId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/1/close", nil)
	ctx.AddParam("id", strconv.FormatInt(fixture.AccountId, 10))

	sut.CloseById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func TestAccountHandler_UpdateById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	input := fixture.AccountUpdateInputCore()
	expected := fixture.AccountCore()

	service.EXPECT().Update(gomock.Any(), fixture.AccountId, gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/1", input)
	ctx.AddParam("id", strconv.FormatInt(fixture.AccountId, 10))

	sut.UpdateById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestAccountHandler_DeleteById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().Delete(gomock.Any(), fixture.AccountId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/accounts/1", nil)
	ctx.AddParam("id", strconv.FormatInt(fixture.AccountId, 10))

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func NewAccountSUT(t *testing.T) (*mocks.MockIAccountService, *AccountHandler) {
	t.Helper()
	return fixture.NewMockSUT(t, mocks.NewMockIAccountService, func(service *mocks.MockIAccountService) *AccountHandler {
		return NewAccountHandler(service)
	})
}
