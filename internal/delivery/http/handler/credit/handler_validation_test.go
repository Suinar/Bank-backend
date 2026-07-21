package credit

import (
	"net/http"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestCreditHandler_GetByUser_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCreditSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/invalid/credits", nil)
	ctx.AddParam("user_id", "invalid")

	sut.GetByUser(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestCreditHandler_GetById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCreditSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/credits/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.GetById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestCreditHandler_Create_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewCreditSUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPost, "/credits", "application/json", "{")

	sut.Create(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestCreditHandler_RepayById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCreditSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/credits/invalid/repay/1000", nil)
	ctx.AddParam("id", "invalid")
	ctx.AddParam("amount", "1000")

	sut.RepayById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestCreditHandler_DeleteById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCreditSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/credits/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.DeleteById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
