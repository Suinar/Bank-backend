package credit

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

var creditHandlerTestError = errors.New("service failure")

func TestCreditHandler_GetAll_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	service.EXPECT().GetAll(gomock.Any()).Return(nil, creditHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/credits", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, creditHandlerTestError.Error())
}

func TestCreditHandler_GetByUser_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	service.EXPECT().GetByUser(gomock.Any(), fixture.UserId).Return(nil, creditHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1/credits", nil)
	ctx.AddParam("user_id", "1")

	sut.GetByUser(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, creditHandlerTestError.Error())
}

func TestCreditHandler_GetById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	service.EXPECT().GetById(gomock.Any(), fixture.CreditId).Return(nil, creditHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/credits/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, creditHandlerTestError.Error())
}

func TestCreditHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	input := core.CreditCreateInput{}

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(nil, creditHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/credits", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, creditHandlerTestError.Error())
}

func TestCreditHandler_RepayById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	service.EXPECT().Repay(gomock.Any(), fixture.CreditId, fixture.TransactionAmount).Return(creditHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/credits/1/repay/1000", nil)
	ctx.AddParam("id", "1")
	ctx.AddParam("amount", strconv.Itoa(fixture.TransactionAmount))

	sut.RepayById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, creditHandlerTestError.Error())
}

func TestCreditHandler_DeleteById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	service.EXPECT().Delete(gomock.Any(), fixture.CreditId).Return(creditHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/credits/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, creditHandlerTestError.Error())
}

func TestCreditHandler_RepayById_InvalidAmount(t *testing.T) {
	t.Parallel()

	_, sut := NewCreditSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/credits/1/repay/invalid", nil)
	ctx.AddParam("id", "1")
	ctx.AddParam("amount", "invalid")

	sut.RepayById(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}
