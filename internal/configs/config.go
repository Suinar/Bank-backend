package configs

import (
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/kVinsom/Bank-backend/internal/logging/config"
	"github.com/spf13/viper"
)

const configFile = ".env"

// KafkaConfig controls the backend-owned Kafka producer and topic.
type KafkaConfig struct {
	Enabled           bool
	Brokers           []string
	ClientID          string
	AuditTopic        string
	Partitions        int
	ReplicationFactor int
}

// Config contains all runtime settings required by the application.
type Config struct {
	// AppConfig describes the current application environment.
	AppConfig struct {
		Environment string
	}

	// HTTPConfig defines the HTTP listener address.
	HTTPConfig struct {
		Host string
		Port int
	}

	// GRPCConfig contains addresses of upstream gRPC services.
	GRPCConfig struct {
		RepositoryURL   string
		ExchangeRateURL string
	}

	// CardConfig contains card issuing settings.
	CardConfig struct {
		BIN string
	}

	// KafkaConfig controls backend-owned audit event publishing.
	KafkaConfig KafkaConfig
}

// LoadConfig reads configuration values from .env and the process environment.
func LoadConfig() (*Config, error) {
	log.Loading(configFile)

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("env")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.LoadingFailed(configFile, err)
			return nil, fmt.Errorf("read config: %w", err)
		}
		log.FileNotFound(configFile)
	} else {
		log.FileLoaded(v.ConfigFileUsed())
	}

	cfg := &Config{}
	cfg.AppConfig.Environment = v.GetString("APP_ENV")
	cfg.HTTPConfig.Host = v.GetString("HTTP_HOST")
	cfg.HTTPConfig.Port = v.GetInt("HTTP_PORT")
	cfg.GRPCConfig.RepositoryURL = v.GetString("GRPC_REPOSITORY_URL")
	cfg.GRPCConfig.ExchangeRateURL = v.GetString("GRPC_EXCHANGE_RATE_URL")
	cfg.CardConfig.BIN = v.GetString("CARD_BIN")
	cfg.KafkaConfig.Enabled = v.GetBool("KAFKA_ENABLED")
	for _, broker := range strings.Split(v.GetString("KAFKA_BROKERS"), ",") {
		if broker = strings.TrimSpace(broker); broker != "" {
			cfg.KafkaConfig.Brokers = append(cfg.KafkaConfig.Brokers, broker)
		}
	}
	cfg.KafkaConfig.ClientID = valueOrDefault(v.GetString("KAFKA_CLIENT_ID"), "bank-backend")
	cfg.KafkaConfig.AuditTopic = valueOrDefault(v.GetString("KAFKA_AUDIT_TOPIC"), "bank.backend.audit.v1")
	cfg.KafkaConfig.Partitions = intOrDefault(v.GetInt("KAFKA_TOPIC_PARTITIONS"), 3)
	cfg.KafkaConfig.ReplicationFactor = intOrDefault(v.GetInt("KAFKA_TOPIC_REPLICATION_FACTOR"), 3)

	if cfg.KafkaConfig.Enabled && len(cfg.KafkaConfig.Brokers) == 0 {
		return nil, errors.New("KAFKA_BROKERS is required when Kafka is enabled")
	}

	log.Loaded(
		cfg.AppConfig.Environment,
		cfg.HTTPConfig.Host,
		cfg.HTTPConfig.Port,
		cfg.GRPCConfig.RepositoryURL,
		cfg.GRPCConfig.ExchangeRateURL,
	)

	return cfg, nil
}

func valueOrDefault(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

func intOrDefault(value, fallback int) int {
	if value <= 0 {
		return fallback
	}
	return value
}
