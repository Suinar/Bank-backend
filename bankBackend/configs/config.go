package configs

import "github.com/spf13/viper"

type Config struct {
	DbUrl string

	Redis struct {
		Address      string
		Password     string
		Db           int
		PoolSize     int
		MinIdleConns int
	}

	Postgres struct {
		MaxOpenConns int
		MaxIdleConns int
	}
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		DbUrl: viper.GetString("DB_URL"),
	}

	cfg.Redis.Address = viper.GetString("REDIS_ADDR")
	cfg.Redis.Password = viper.GetString("REDIS_PASSWORD")
	cfg.Redis.Db = viper.GetInt("REDIS_DB")
	cfg.Redis.PoolSize = viper.GetInt("REDIS_POOL_SIZE")
	cfg.Redis.MinIdleConns = viper.GetInt("REDIS_MIN_IDLE")

	return cfg, nil
}
