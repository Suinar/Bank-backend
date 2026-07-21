package configs

import (
	"errors"
	"fmt"
	"os"

	log "github.com/kVinsom/Bank-backend/internal/logging/config"
	"github.com/spf13/viper"
)

const configFile = ".env"

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

	log.Loaded(
		cfg.AppConfig.Environment,
		cfg.HTTPConfig.Host,
		cfg.HTTPConfig.Port,
		cfg.GRPCConfig.RepositoryURL,
		cfg.GRPCConfig.ExchangeRateURL,
	)

	return cfg, nil
}
