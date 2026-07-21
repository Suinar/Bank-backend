package credit

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kVinsom/Bank-backend/internal/mocks/services"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestCreditHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	expected := fixture.CreditListCore()

	service.EXPECT().GetAll(gomock.Any()).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/credits", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCreditHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	expected := fixture.CreditListCore()

	service.EXPECT().GetByUser(gomock.Any(), fixture.UserId).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1/credits", nil)
	ctx.AddParam("user_id", "1")

	sut.GetByUser(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCreditHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	expected := fixture.CreditCore()

	service.EXPECT().GetById(gomock.Any(), fixture.CreditId).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/credits/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCreditHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	input := fixture.CreditCreateInputCore()
	expected := fixture.CreditCore()

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/credits", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCreditHandler_RepayById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	service.EXPECT().Repay(gomock.Any(), fixture.CreditId, fixture.TransactionAmount).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/credits/1/repay/1000", nil)
	ctx.AddParam("id", "1")
	ctx.AddParam("amount", strconv.Itoa(fixture.TransactionAmount))

	sut.RepayById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func TestCreditHandler_DeleteById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCreditSUT(t)

	service.EXPECT().Delete(gomock.Any(), fixture.CreditId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/credits/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func NewCreditSUT(t *testing.T) (*mocks.MockICreditService, *CreditHandler) {
	t.Helper()
	return fixture.NewMockSUT(t, mocks.NewMockICreditService, func(service *mocks.MockICreditService) *CreditHandler {
		return NewCreditHandler(service)
	})
}
