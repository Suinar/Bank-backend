package configs

import "github.com/spf13/viper"

type Config struct {
	Grpc struct {
		RepositoryURL   string
		ExchangeRateURL string
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	cfg.Grpc.RepositoryURL = viper.GetString("GRPC_REPOSITORY_URL")
	cfg.Grpc.ExchangeRateURL = viper.GetString("GRPC_EXCHANGE_RATE_URL")

	return cfg, nil
}
