package currency

import (
	"net/http"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestCurrencyHandler_Create_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewCurrencySUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPost, "/currencies", "application/json", "{")

	sut.Create(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestCurrencyHandler_UpdateById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCurrencySUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/currencies/invalid", fixture.CurrencyUpdateInputCore())
	ctx.AddParam("id", "invalid")

	sut.UpdateById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestCurrencyHandler_UpdateById_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewCurrencySUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPatch, "/currencies/1", "application/json", "{")
	ctx.AddParam("id", "1")

	sut.UpdateById(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestCurrencyHandler_DeleteById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCurrencySUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/currencies/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.DeleteById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestCurrencyHandler_GetBySymbol_InvalidSymbol(t *testing.T) {
	t.Parallel()

	for _, symbol := range []string{"", "USD"} {
		symbol := symbol
		t.Run(symbol, func(t *testing.T) {
			t.Parallel()

			_, sut := NewCurrencySUT(t)
			ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/currencies/symbol/"+symbol, nil)
			ctx.AddParam("symbol", symbol)

			sut.GetBySymbol(ctx)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
			}
		})
	}
}
