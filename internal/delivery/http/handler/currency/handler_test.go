package currency

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kVinsom/Bank-backend/internal/mocks/services"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestCurrencyHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	expected := fixture.CurrencyListCore()

	service.EXPECT().GetAll(gomock.Any()).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCurrencyHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	expected := fixture.CurrencyCore()

	service.EXPECT().GetById(gomock.Any(), test.CurrencyId).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCurrencyHandler_GetByIso_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	expected := fixture.CurrencyCore()

	service.EXPECT().GetByIso(gomock.Any(), test.CurrencyISOCode).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies/iso/USD", nil)
	ctx.AddParam("iso_code", test.CurrencyISOCode)

	sut.GetByIso(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCurrencyHandler_GetBySymbol_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	expected := fixture.CurrencyCore()

	service.EXPECT().GetBySymbol(gomock.Any(), test.CurrencySymbol).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies/symbol/$", nil)
	ctx.AddParam("symbol", string(test.CurrencySymbol))

	sut.GetBySymbol(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCurrencyHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	input := fixture.CurrencyCreateInputCore()
	expected := fixture.CurrencyCore()

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/currencies", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCurrencyHandler_UpdateById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	input := fixture.CurrencyUpdateInputCore()
	expected := fixture.CurrencyCore()

	service.EXPECT().Update(gomock.Any(), test.CurrencyId, gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/currencies/1", input)
	ctx.AddParam("id", "1")

	sut.UpdateById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCurrencyHandler_DeleteById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	service.EXPECT().Delete(gomock.Any(), test.CurrencyId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/currencies/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func NewCurrencySUT(t *testing.T) (*mocks.MockICurrencyService, *CurrencyHandler) {
	t.Helper()
	return fixture.NewMockSUT(t, mocks.NewMockICurrencyService, func(service *mocks.MockICurrencyService) *CurrencyHandler {
		return NewCurrencyHandler(service)
	})
}
