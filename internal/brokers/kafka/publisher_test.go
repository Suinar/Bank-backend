package kafka_test

import (
	"context"
	"encoding/json"
	"testing"

	kafkaBroker "github.com/kVinsom/Bank-backend/internal/brokers/kafka"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestPublisherPublishAudit(t *testing.T) {
	writer, publisher := NewSUT(t)

	err := publisher.PublishAudit(context.Background(), fixture.KafkaAuditEvent())
	if err != nil {
		t.Fatalf("PublishAudit() error = %v", err)
	}
	if len(writer.Messages) != 1 {
		t.Fatalf("message count = %d, want 1", len(writer.Messages))
	}

	message := writer.Messages[0]
	if message.Topic != fixture.KafkaAuditTopic || string(message.Key) != fixture.KafkaCorrelationID {
		t.Fatalf("unexpected Kafka message routing: %#v", message)
	}
	var event kafkaBroker.AuditEvent
	if err := json.Unmarshal(message.Value, &event); err != nil {
		t.Fatalf("unmarshal event: %v", err)
	}
	if event.EventType != "backend.http.operation.completed" ||
		event.EventVersion != 1 ||
		event.Producer != "bank-backend" {
		t.Fatalf("unexpected envelope: %#v", event)
	}

	if err := publisher.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if !writer.Closed {
		t.Fatal("writer was not closed")
	}
}

func NewSUT(t *testing.T) (*fixture.KafkaWriterStub, *kafkaBroker.Publisher) {
	t.Helper()
	writer := &fixture.KafkaWriterStub{}
	return writer, kafkaBroker.NewPublisher(
		writer,
		fixture.KafkaAuditTopic,
		fixture.KafkaProducer,
	)
}
