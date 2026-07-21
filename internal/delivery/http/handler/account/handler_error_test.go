package account

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

var accountHandlerTestError = errors.New("service failure")

func TestAccountHandler_GetAll_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().GetAll(gomock.Any()).Return(nil, accountHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/accounts", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, accountHandlerTestError.Error())
}

func TestAccountHandler_GetByUser_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().GetByUser(gomock.Any(), fixture.UserId).Return(nil, accountHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1/accounts", nil)
	ctx.AddParam("user_id", "1")

	sut.GetByUser(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, accountHandlerTestError.Error())
}

func TestAccountHandler_GetById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().GetById(gomock.Any(), fixture.AccountId).Return(nil, accountHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/accounts/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, accountHandlerTestError.Error())
}

func TestAccountHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	input := core.AccountCreateInput{}

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(nil, accountHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/accounts", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, accountHandlerTestError.Error())
}

func TestAccountHandler_BlockingById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().Blocking(gomock.Any(), fixture.AccountId).Return(accountHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/1/block", nil)
	ctx.AddParam("id", "1")

	sut.BlockingById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, accountHandlerTestError.Error())
}

func TestAccountHandler_CloseById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().Close(gomock.Any(), fixture.AccountId).Return(accountHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/1/close", nil)
	ctx.AddParam("id", "1")

	sut.CloseById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, accountHandlerTestError.Error())
}

func TestAccountHandler_UpdateById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	input := core.AccountUpdateInput{}

	service.EXPECT().Update(gomock.Any(), fixture.AccountId, gomock.Eq(&input)).Return(nil, accountHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/1", input)
	ctx.AddParam("id", "1")

	sut.UpdateById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, accountHandlerTestError.Error())
}

func TestAccountHandler_DeleteById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewAccountSUT(t)

	service.EXPECT().Delete(gomock.Any(), fixture.AccountId).Return(accountHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/accounts/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, accountHandlerTestError.Error())
}

func TestAccountHandler_GetById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewAccountSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/accounts/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.GetById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
