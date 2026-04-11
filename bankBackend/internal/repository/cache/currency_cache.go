package cache

import (
	"context"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/redis/go-redis/v9"
)

type CurrencyCache struct {
	rdb *redis.Client
}

func NewCurrencyCache(rdb *redis.Client) *CurrencyCache {
	return &CurrencyCache{rdb: rdb}
}

func (c *CurrencyCache) GetAll(ctx context.Context) ([]core.Currency, error) {
	ids, err := c.rdb.SMembers(ctx, "currencies").Result()
	if err != nil {
		return nil, core.InternalServerError
	}

	if len(ids) <= 0 {
		return []core.Currency{}, nil
	}

	var currencies []core.Currency

	for _, id := range ids {
		key := "currency:" + id

		data, err := c.rdb.HGetAll(ctx, key).Result()
		if err != nil {
			return nil, core.InternalServerError
		}

		if len(data) == 0 {
			continue
		}

		currency, err := c.MapToCurrency(data)
		if err != nil {
			if errors.Is(err, core.BadRequest) {
				return nil, core.BadRequest
			}

			return nil, core.InternalServerError
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (c *CurrencyCache) GetById(ctx context.Context, id uint64) (*core.Currency, error) {
	return nil, nil
}

func (c *CurrencyCache) GetByIso(ctx context.Context, iso string) (*core.Currency, error) {
	return nil, nil
}

func (c *CurrencyCache) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	return nil, nil
}

func (c *CurrencyCache) Delete(ctx context.Context, id uint64) error {
	return nil
}

func (c *CurrencyCache) MapToCurrency(data map[string]string) (core.Currency, error) {
	return core.Currency{}, nil
}

func (c *CurrencyCache) Set(ctx context.Context, currency *core.Currency) error {
	return nil
}

func (c *CurrencyCache) SetAll(ctx context.Context, currencies []core.Currency) error {
	return nil
}

func (c *CurrencyCache) Update(ctx context.Context, id uint64, currency *core.Currency) error {
	return nil
}

func (c *CurrencyCache) UpdateAll(ctx context.Context, currencies []core.Currency) error {
	return nil
}
