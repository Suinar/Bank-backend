package common

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
)

func TestParseHandlers(t *testing.T) {
	t.Parallel()

	if got, err := IdParseHandler(nil, "42"); err != nil || got != 42 {
		t.Fatalf("IdParseHandler valid: got %d, %v", got, err)
	}
	if _, err := IdParseHandler(nil, "invalid"); err == nil {
		t.Fatal("IdParseHandler invalid: expected error")
	}
	if got, err := IsoParseHandler(nil, "840"); err != nil || got != 840 {
		t.Fatalf("IsoParseHandler valid: got %d, %v", got, err)
	}
	if _, err := IsoParseHandler(nil, "invalid"); !errors.Is(err, coreErrors.BadRequest) {
		t.Fatalf("IsoParseHandler invalid: want BadRequest, got %v", err)
	}
}

func TestErrorHandler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    error
		status int
	}{
		{name: "not found", err: coreErrors.NotFound, status: http.StatusNotFound},
		{name: "bad request", err: coreErrors.BadRequest, status: http.StatusBadRequest},
		{name: "internal", err: errors.New("failure"), status: http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder, ctx := newCommonContext()
			ErrorHandler(ctx, tt.err)
			if recorder.Code != tt.status {
				t.Fatalf("status: want %d, got %d", tt.status, recorder.Code)
			}
		})
	}
}

func TestResponseHelpers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		invoke func(*gin.Context)
		status int
	}{
		{name: "success", invoke: func(ctx *gin.Context) { ResponseSuccess(ctx, gin.H{"ok": true}) }, status: http.StatusOK},
		{name: "bad request", invoke: func(ctx *gin.Context) { ResponseBadRequest(ctx, coreErrors.BadRequest) }, status: http.StatusBadRequest},
		{name: "not found", invoke: func(ctx *gin.Context) { ResponseNotFound(ctx, coreErrors.NotFound) }, status: http.StatusNotFound},
		{name: "internal", invoke: func(ctx *gin.Context) { ResponseInternalServerError(ctx, errors.New("failure")) }, status: http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder, ctx := newCommonContext()
			tt.invoke(ctx)
			if recorder.Code != tt.status {
				t.Fatalf("status: want %d, got %d", tt.status, recorder.Code)
			}
		})
	}
}

func newCommonContext() (*httptest.ResponseRecorder, *gin.Context) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	return recorder, ctx
}
