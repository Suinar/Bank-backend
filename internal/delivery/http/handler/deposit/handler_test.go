package deposit

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kVinsom/Bank-backend/internal/mocks/services"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestDepositHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	expected := fixture.DepositListCore()

	service.EXPECT().GetAll(gomock.Any()).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/deposits", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestDepositHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	expected := fixture.DepositListCore()

	service.EXPECT().GetByUser(gomock.Any(), test.UserId).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1/deposits", nil)
	ctx.AddParam("user_id", "1")

	sut.GetByUser(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestDepositHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	expected := fixture.DepositCore()

	service.EXPECT().GetById(gomock.Any(), test.DepositId).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/deposits/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestDepositHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	input := fixture.DepositCreateInputCore()
	expected := fixture.DepositCore()

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/deposits", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestDepositHandler_ReplenishById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	service.EXPECT().Replenish(gomock.Any(), test.DepositId, test.TransactionAmount).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/deposits/1/replenish/1000", nil)
	ctx.AddParam("id", "1")
	ctx.AddParam("amount", strconv.Itoa(test.TransactionAmount))

	sut.ReplenishById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func TestDepositHandler_DeleteById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewDepositSUT(t)

	service.EXPECT().Delete(gomock.Any(), test.DepositId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/deposits/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func NewDepositSUT(t *testing.T) (*mocks.MockIDepositService, *DepositHandler) {
	t.Helper()
	return fixture.NewMockSUT(t, mocks.NewMockIDepositService, func(service *mocks.MockIDepositService) *DepositHandler {
		return NewDepositHandler(service)
	})
}
