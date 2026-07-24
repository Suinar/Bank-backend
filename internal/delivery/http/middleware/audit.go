package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	kafkaBroker "github.com/kVinsom/Bank-backend/internal/brokers/kafka"
)

const correlationHeader = "X-Correlation-ID"

// KafkaAudit publishes metadata for completed API operations without request bodies.
func KafkaAudit(publisher kafkaBroker.EventPublisher) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/health" {
			ctx.Next()
			return
		}

		correlationID := ctx.GetHeader(correlationHeader)
		if correlationID == "" {
			correlationID = newEventID()
		}
		ctx.Header(correlationHeader, correlationID)

		started := time.Now().UTC()
		ctx.Next()

		route := ctx.FullPath()
		if route == "" {
			route = "unmatched"
		}
		status := ctx.Writer.Status()
		outcome := "success"
		if status >= http.StatusBadRequest {
			outcome = "failure"
		}

		event := kafkaBroker.AuditEvent{
			EventID:       newEventID(),
			CorrelationID: correlationID,
			Method:        ctx.Request.Method,
			Route:         route,
			StatusCode:    status,
			Outcome:       outcome,
			DurationMS:    time.Since(started).Milliseconds(),
			OccurredAt:    started,
		}
		if err := publisher.PublishAudit(context.Background(), event); err != nil {
			log.Printf("HTTP audit event was not queued: %v", err)
		}
	}
}

func newEventID() string {
	var value [16]byte
	if _, err := rand.Read(value[:]); err != nil {
		return hex.EncodeToString([]byte(time.Now().UTC().Format(time.RFC3339Nano)))
	}
	value[6] = (value[6] & 0x0f) | 0x40
	value[8] = (value[8] & 0x3f) | 0x80
	return hex.EncodeToString(value[:])
}
