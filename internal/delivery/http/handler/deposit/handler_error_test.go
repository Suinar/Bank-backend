package deposit

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

var depositHandlerTestError = errors.New("service failure")

func TestDepositHandler_GetAll_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	service.EXPECT().GetAll(gomock.Any()).Return(nil, depositHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/deposits", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, depositHandlerTestError.Error())
}

func TestDepositHandler_GetByUser_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	service.EXPECT().GetByUser(gomock.Any(), fixture.UserId).Return(nil, depositHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1/deposits", nil)
	ctx.AddParam("user_id", "1")

	sut.GetByUser(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, depositHandlerTestError.Error())
}

func TestDepositHandler_GetById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	service.EXPECT().GetById(gomock.Any(), fixture.DepositId).Return(nil, depositHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/deposits/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, depositHandlerTestError.Error())
}

func TestDepositHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	input := core.DepositCreateInput{}

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(nil, depositHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/deposits", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, depositHandlerTestError.Error())
}

func TestDepositHandler_ReplenishById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	service.EXPECT().Replenish(gomock.Any(), fixture.DepositId, fixture.TransactionAmount).Return(depositHandlerTestError).Times(1)

	form := url.Values{"amount": {strconv.Itoa(fixture.TransactionAmount)}}
	ctx, recorder := fixture.NewFormHTTPContext(t, http.MethodPatch, "/deposits/1/replenish", form)
	ctx.AddParam("id", "1")

	sut.ReplenishById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, depositHandlerTestError.Error())
}

func TestDepositHandler_DeleteById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	service.EXPECT().Delete(gomock.Any(), fixture.DepositId).Return(depositHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/deposits/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, depositHandlerTestError.Error())
}

func TestDepositHandler_ReplenishById_InvalidAmount(t *testing.T) {
	t.Parallel()

	_, sut := NewDepositSUT(t)

	form := url.Values{"amount": {"invalid"}}
	ctx, recorder := fixture.NewFormHTTPContext(t, http.MethodPatch, "/deposits/1/replenish", form)
	ctx.AddParam("id", "1")

	sut.ReplenishById(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}
