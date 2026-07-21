package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	serviceContainer "github.com/kVinsom/Bank-backend/internal/service"
)

func TestInitHandlers(t *testing.T) {
	router := InitHandlers(&serviceContainer.Services{})

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("health status: want %d, got %d", http.StatusOK, recorder.Code)
	}
	if len(router.Routes()) != 46 {
		t.Fatalf("registered routes: want 46, got %d", len(router.Routes()))
	}
}
