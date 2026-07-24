package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestKafkaAuditPublishesRouteMetadata(t *testing.T) {
	publisher, router := NewSUT(t)
	router.POST("/api/v1/accounts/:id", func(ctx *gin.Context) {
		ctx.Status(http.StatusCreated)
	})

	request := httptest.NewRequest(http.MethodPost, "/api/v1/accounts/42", nil)
	request.Header.Set(correlationHeader, "request-123")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if len(publisher.Events) != 1 {
		t.Fatalf("event count = %d, want 1", len(publisher.Events))
	}
	event := publisher.Events[0]
	if event.Route != "/api/v1/accounts/:id" ||
		event.CorrelationID != "request-123" ||
		event.StatusCode != http.StatusCreated ||
		event.Outcome != "success" {
		t.Fatalf("unexpected audit event: %#v", event)
	}
	if response.Header().Get(correlationHeader) != "request-123" {
		t.Fatalf("missing response correlation header")
	}
}

func TestKafkaAuditSkipsHealth(t *testing.T) {
	publisher, router := NewSUT(t)
	router.GET("/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/health", nil))
	if len(publisher.Events) != 0 {
		t.Fatalf("health endpoint published %d events", len(publisher.Events))
	}
}

func NewSUT(t *testing.T) (*fixture.KafkaEventPublisherStub, *gin.Engine) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	publisher := &fixture.KafkaEventPublisherStub{}
	router := gin.New()
	router.Use(KafkaAudit(publisher))
	return publisher, router
}
