package cache

import "github.com/redis/go-redis/v9"

type CurrencyCache struct {
	rdb *redis.Client
}
