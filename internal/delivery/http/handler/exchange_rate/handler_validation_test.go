package exchange_rate

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestExchangeRateHandler_GetRelativeRanking_InvalidIsoFrom(t *testing.T) {
	t.Parallel()

	_, sut := NewExchangeRateSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/exchange-rates/invalid/978", nil)
	ctx.AddParam("currency_iso_from", "invalid")
	ctx.AddParam("currency_iso_to", strconv.Itoa(test.ExchangeISOTo))

	sut.GetRelativeRanking(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestExchangeRateHandler_GetRelativeRanking_InvalidIsoTo(t *testing.T) {
	t.Parallel()

	_, sut := NewExchangeRateSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/exchange-rates/840/invalid", nil)
	ctx.AddParam("currency_iso_from", strconv.Itoa(test.ExchangeISOFrom))
	ctx.AddParam("currency_iso_to", "invalid")

	sut.GetRelativeRanking(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}
