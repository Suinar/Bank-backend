package kafka

import (
	"context"

	kafkaGo "github.com/segmentio/kafka-go"
)

// EventPublisher is the Kafka operation required by the HTTP audit middleware.
type EventPublisher interface {
	PublishAudit(context.Context, AuditEvent) error
}

// MessageWriter contains the Kafka writer operations used by Publisher.
type MessageWriter interface {
	WriteMessages(context.Context, ...kafkaGo.Message) error
	Close() error
}
