package fixture

import (
	"context"
	"time"

	kafkaBroker "github.com/kVinsom/Bank-backend/internal/brokers/kafka"
	kafkaGo "github.com/segmentio/kafka-go"
)

const (
	KafkaAuditTopic    = "bank.backend.audit.v1"
	KafkaProducer      = "bank-backend"
	KafkaCorrelationID = "correlation-1"
)

// KafkaWriterStub records messages written by the Kafka publisher.
type KafkaWriterStub struct {
	Messages []kafkaGo.Message
	Closed   bool
}

func (w *KafkaWriterStub) WriteMessages(_ context.Context, messages ...kafkaGo.Message) error {
	w.Messages = append(w.Messages, messages...)
	return nil
}

func (w *KafkaWriterStub) Close() error {
	w.Closed = true
	return nil
}

// KafkaEventPublisherStub records audit events produced by HTTP middleware.
type KafkaEventPublisherStub struct {
	Events []kafkaBroker.AuditEvent
}

func (p *KafkaEventPublisherStub) PublishAudit(
	_ context.Context,
	event kafkaBroker.AuditEvent,
) error {
	p.Events = append(p.Events, event)
	return nil
}

// KafkaAuditEvent returns a deterministic backend audit event.
func KafkaAuditEvent() kafkaBroker.AuditEvent {
	return kafkaBroker.AuditEvent{
		EventID:       "event-1",
		CorrelationID: KafkaCorrelationID,
		Method:        "POST",
		Route:         "/api/v1/accounts",
		StatusCode:    200,
		Outcome:       "success",
		DurationMS:    12,
		OccurredAt:    time.Date(2026, time.July, 24, 10, 0, 0, 0, time.UTC),
	}
}
