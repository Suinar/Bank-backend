package card

import (
	"net/http"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestCardHandler_GetByUser_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCardSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/invalid/cards", nil)
	ctx.AddParam("user_id", "invalid")

	sut.GetByUser(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestCardHandler_Create_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewCardSUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPost, "/cards", "application/json", "{")

	sut.Create(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestCardHandler_BlockingById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCardSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/cards/invalid/block", nil)
	ctx.AddParam("id", "invalid")

	sut.BlockingById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestCardHandler_DeleteById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCardSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/cards/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.DeleteById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
