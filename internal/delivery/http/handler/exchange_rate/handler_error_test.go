package exchange_rate

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

var exchangeRateHandlerTestError = errors.New("service failure")

func TestExchangeRateHandler_GetAllRanking_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewExchangeRateSUT(t)

	service.EXPECT().GetAllRanking(gomock.Any(), fixture.ExchangeISOFrom).Return(nil, exchangeRateHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/exchange-rates/840", nil)
	ctx.AddParam("currency_iso_from", strconv.Itoa(fixture.ExchangeISOFrom))

	sut.GetAllRanking(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, exchangeRateHandlerTestError.Error())
}

func TestExchangeRateHandler_GetRelativeRanking_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewExchangeRateSUT(t)

	service.EXPECT().GetRelativeRanking(gomock.Any(), fixture.ExchangeISOFrom, fixture.ExchangeISOTo).Return(nil, exchangeRateHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/exchange-rates/840/978", nil)
	ctx.AddParam("currency_iso_from", strconv.Itoa(fixture.ExchangeISOFrom))
	ctx.AddParam("currency_iso_to", strconv.Itoa(fixture.ExchangeISOTo))

	sut.GetRelativeRanking(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, exchangeRateHandlerTestError.Error())
}

func TestExchangeRateHandler_GetAllRanking_InvalidISO(t *testing.T) {
	t.Parallel()

	_, sut := NewExchangeRateSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/exchange-rates/invalid", nil)
	ctx.AddParam("currency_iso_from", "invalid")

	sut.GetAllRanking(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}
