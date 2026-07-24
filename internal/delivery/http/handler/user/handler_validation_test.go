package user

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestUserHandler_Create_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewUserSUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPost, "/users", "application/json", "{")

	sut.Create(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestUserHandler_UpdateById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewUserSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/users/invalid", fixture.UserUpdateInputCore())
	ctx.AddParam("id", "invalid")

	sut.UpdateById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestUserHandler_UpdateById_InvalidJson(t *testing.T) {
	t.Parallel()

	_, sut := NewUserSUT(t)

	ctx, recorder := fixture.NewRawHTTPContext(t, http.MethodPatch, "/users/1", "application/json", "{")
	ctx.AddParam("id", "1")

	sut.UpdateById(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("want %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestUserHandler_ChangePasswordById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewUserSUT(t)

	form := url.Values{"hash_password": {test.UserPasswordHash}}
	ctx, recorder := fixture.NewFormHTTPContext(t, http.MethodPatch, "/users/invalid/password", form)
	ctx.AddParam("id", "invalid")

	sut.ChangePasswordById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}

func TestUserHandler_DeleteById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewUserSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/users/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.DeleteById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
