package deposit

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestDepositHandler_GetById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewDepositSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/deposits/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.GetById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestDepositHandler_Create_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewDepositSUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPost, "/deposits", "application/json", "{")

	sut.Create(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestDepositHandler_ReplenishById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewDepositSUT(t)

	form := url.Values{"amount": {"1000"}}
	ctx, recorder := fixture.NewFormHTTPContext(t, http.MethodPatch, "/deposits/invalid/replenish", form)
	ctx.AddParam("id", "invalid")

	sut.ReplenishById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestDepositHandler_DeleteById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewDepositSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/deposits/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.DeleteById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
