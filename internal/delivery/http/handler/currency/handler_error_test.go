package currency

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

var currencyHandlerTestError = errors.New("service failure")

func TestCurrencyHandler_GetAll_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	service.EXPECT().GetAll(gomock.Any()).Return(nil, currencyHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, currencyHandlerTestError.Error())
}

func TestCurrencyHandler_GetById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	service.EXPECT().GetById(gomock.Any(), test.CurrencyId).Return(nil, currencyHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, currencyHandlerTestError.Error())
}

func TestCurrencyHandler_GetByIso_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	service.EXPECT().GetByIso(gomock.Any(), test.CurrencyISOCode).Return(nil, currencyHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies/iso/USD", nil)
	ctx.AddParam("iso_code", test.CurrencyISOCode)

	sut.GetByIso(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, currencyHandlerTestError.Error())
}

func TestCurrencyHandler_GetBySymbol_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	service.EXPECT().GetBySymbol(gomock.Any(), test.CurrencySymbol).Return(nil, currencyHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies/symbol/$", nil)
	ctx.AddParam("symbol", string(test.CurrencySymbol))

	sut.GetBySymbol(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, currencyHandlerTestError.Error())
}

func TestCurrencyHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	input := core.CurrencyCreateInput{}

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(nil, currencyHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/currencies", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, currencyHandlerTestError.Error())
}

func TestCurrencyHandler_UpdateById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	input := core.CurrencyUpdateInput{}

	service.EXPECT().Update(gomock.Any(), test.CurrencyId, gomock.Eq(&input)).Return(nil, currencyHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/currencies/1", input)
	ctx.AddParam("id", "1")

	sut.UpdateById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, currencyHandlerTestError.Error())
}

func TestCurrencyHandler_DeleteById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCurrencySUT(t)

	service.EXPECT().Delete(gomock.Any(), test.CurrencyId).Return(currencyHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/currencies/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, currencyHandlerTestError.Error())
}

func TestCurrencyHandler_GetById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCurrencySUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.GetById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
