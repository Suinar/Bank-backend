package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kVinsom/Bank-backend/internal/configs"
	kafkaGo "github.com/segmentio/kafka-go"
)

// AuditEvent is a versioned, non-sensitive record of a completed HTTP operation.
type AuditEvent struct {
	EventID       string    `json:"event_id"`
	EventType     string    `json:"event_type"`
	EventVersion  int       `json:"event_version"`
	Producer      string    `json:"producer"`
	CorrelationID string    `json:"correlation_id"`
	Method        string    `json:"method"`
	Route         string    `json:"route"`
	StatusCode    int       `json:"status_code"`
	Outcome       string    `json:"outcome"`
	DurationMS    int64     `json:"duration_ms"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// Publisher publishes backend-owned events and owns the Kafka writer lifecycle.
type Publisher struct {
	writer   MessageWriter
	topic    string
	producer string
}

// NewPublisher creates a publisher over an injected Kafka writer.
func NewPublisher(writer MessageWriter, topic, producer string) *Publisher {
	return &Publisher{writer: writer, topic: topic, producer: producer}
}

// Open ensures the backend-owned topic exists and creates an asynchronous writer.
func Open(ctx context.Context, cfg configs.KafkaConfig) (*Publisher, error) {
	if len(cfg.Brokers) == 0 {
		return nil, errors.New("at least one Kafka broker is required")
	}
	if cfg.AuditTopic == "" {
		return nil, errors.New("Kafka audit topic is required")
	}
	if cfg.Partitions <= 0 || cfg.ReplicationFactor <= 0 {
		return nil, errors.New("Kafka topic partitions and replication factor must be positive")
	}

	client := &kafkaGo.Client{Addr: kafkaGo.TCP(cfg.Brokers...)}
	var ensureErr error
	for {
		ensureErr = ensureTopic(ctx, client, cfg)
		if ensureErr == nil {
			break
		}
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("prepare Kafka audit topic: %w", ensureErr)
		case <-time.After(time.Second):
		}
	}

	writer := &kafkaGo.Writer{
		Addr:         kafkaGo.TCP(cfg.Brokers...),
		Balancer:     &kafkaGo.Hash{},
		RequiredAcks: kafkaGo.RequireAll,
		Async:        true,
		BatchTimeout: 100 * time.Millisecond,
		Completion: func(_ []kafkaGo.Message, err error) {
			if err != nil {
				log.Printf("kafka audit publisher: asynchronous publish failed: %v", err)
			}
		},
	}
	return NewPublisher(writer, cfg.AuditTopic, cfg.ClientID), nil
}

func ensureTopic(ctx context.Context, client *kafkaGo.Client, cfg configs.KafkaConfig) error {
	response, err := client.CreateTopics(ctx, &kafkaGo.CreateTopicsRequest{
		Topics: []kafkaGo.TopicConfig{{
			Topic:             cfg.AuditTopic,
			NumPartitions:     cfg.Partitions,
			ReplicationFactor: cfg.ReplicationFactor,
		}},
	})
	if err != nil {
		return fmt.Errorf("create Kafka audit topic: %w", err)
	}
	if response == nil {
		return errors.New("create Kafka audit topic: empty response")
	}
	if topicErr := response.Errors[cfg.AuditTopic]; topicErr != nil &&
		!errors.Is(topicErr, kafkaGo.TopicAlreadyExists) {
		return fmt.Errorf("create Kafka audit topic %q: %w", cfg.AuditTopic, topicErr)
	}
	return nil
}

// PublishAudit queues one audit event without putting Kafka on the API response path.
func (p *Publisher) PublishAudit(ctx context.Context, event AuditEvent) error {
	if p == nil || p.writer == nil {
		return errors.New("Kafka publisher is not initialized")
	}
	event.EventType = "backend.http.operation.completed"
	event.EventVersion = 1
	event.Producer = p.producer

	value, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal Kafka audit event: %w", err)
	}
	return p.writer.WriteMessages(ctx, kafkaGo.Message{
		Topic: p.topic,
		Key:   []byte(event.CorrelationID),
		Value: value,
		Time:  event.OccurredAt,
	})
}

// Close flushes pending events and releases writer resources.
func (p *Publisher) Close() error {
	if p == nil || p.writer == nil {
		return nil
	}
	return p.writer.Close()
}
