package account

import (
	"net/http"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestAccountHandler_GetByUser_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewAccountSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/invalid/accounts", nil)
	ctx.AddParam("user_id", "invalid")

	sut.GetByUser(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestAccountHandler_Create_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewAccountSUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPost, "/accounts", "application/json", "{")

	sut.Create(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestAccountHandler_BlockingById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewAccountSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/invalid/block", nil)
	ctx.AddParam("id", "invalid")

	sut.BlockingById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestAccountHandler_CloseById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewAccountSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/invalid/close", nil)
	ctx.AddParam("id", "invalid")

	sut.CloseById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestAccountHandler_UpdateById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewAccountSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/accounts/invalid", fixture.AccountUpdateInputCore())
	ctx.AddParam("id", "invalid")

	sut.UpdateById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestAccountHandler_UpdateById_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewAccountSUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPatch, "/accounts/1", "application/json", "{")
	ctx.AddParam("id", "1")

	sut.UpdateById(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestAccountHandler_DeleteById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewAccountSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/accounts/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.DeleteById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
