package cache

import (
	"context"
	"errors"
	"strconv"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/redis/go-redis/v9"
)

const (
	currenciesSetKey  = "currencies"
	currencyKeyPrefix = "currency:"
)

type CurrencyCache struct {
	rdb *redis.Client
}

func NewCurrencyCache(rdb *redis.Client) *CurrencyCache {
	return &CurrencyCache{rdb: rdb}
}

func (c *CurrencyCache) GetAll(ctx context.Context) ([]core.Currency, error) {
	ids, err := c.rdb.SMembers(ctx, currenciesSetKey).Result()
	if err != nil {
		return nil, core.CacheGetError
	}

	if len(ids) == 0 {
		return []core.Currency{}, nil
	}

	pipe := c.rdb.Pipeline()
	cmds := make([]*redis.MapStringStringCmd, 0, len(ids))

	for _, id := range ids {
		parsedId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, core.CacheGetError
		}

		key := c.GetPrimaryKey(parsedId)
		cmds = append(cmds, pipe.HGetAll(ctx, key))
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, core.CacheGetError
	}

	currencies := make([]core.Currency, 0, len(ids))

	for _, cmd := range cmds {
		data := cmd.Val()

		if len(data) == 0 {
			continue
		}

		currency, err := c.MapToCurrency(data)
		if err != nil {
			continue
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (c *CurrencyCache) GetById(ctx context.Context, id uint64) (*core.Currency, error) {
	primaryKey := c.GetPrimaryKey(id)

	data, err := c.rdb.HGetAll(ctx, primaryKey).Result()
	if err != nil {
		return nil, core.CacheGetError
	}

	if len(data) == 0 {
		return nil, nil
	}

	currency, err := c.MapToCurrency(data)
	if err != nil {
		if errors.Is(err, core.CacheMapError) {
			return nil, core.CacheMapError
		}

		return nil, core.InternalServerError
	}

	return &currency, nil
}

func (c *CurrencyCache) GetByIso(ctx context.Context, iso string) (*core.Currency, error) {
	idStr, err := c.rdb.Get(ctx, c.GetIsoKey(iso)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, core.CacheGetError
		}

		return nil, core.InternalServerError
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, core.BadRequest
	}

	return c.GetById(ctx, id)
}

func (c *CurrencyCache) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	idStr, err := c.rdb.Get(ctx, c.GetSymbolKey(symbol)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, core.CacheGetError
		}

		return nil, core.InternalServerError
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, core.BadRequest
	}

	return c.GetById(ctx, id)
}
func (c *CurrencyCache) Set(ctx context.Context, currency *core.Currency) error {
	primaryKey := c.GetPrimaryKey(currency.Id)
	isoKey := c.GetIsoKey(currency.IsoCode)
	symbolKey := c.GetSymbolKey(currency.Symbol)
	minorUnitsKey := c.GetMinorUnitsKey(currency.MinorUnits)

	data := map[string]interface{}{
		"id":          currency.Id,
		"name":        currency.Name,
		"symbol":      string(currency.Symbol),
		"iso_code":    currency.IsoCode,
		"minor_units": currency.MinorUnits,
	}

	pipe := c.rdb.Pipeline()

	pipe.HSet(ctx, primaryKey, data)

	pipe.SAdd(ctx, currenciesSetKey, strconv.FormatUint(currency.Id, 10))

	pipe.Set(ctx, isoKey, currency.Id, 0)

	pipe.SAdd(ctx, symbolKey, currency.Id)
	pipe.SAdd(ctx, minorUnitsKey, currency.Id)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return core.CacheSetError
	}

	return nil
}

func (c *CurrencyCache) SetAll(ctx context.Context, currencies []core.Currency) error {
	if len(currencies) == 0 {
		return nil
	}

	keys, err := c.rdb.Keys(ctx, currencyKeyPrefix+"*").Result()
	if err == nil && len(keys) > 0 {
		_ = c.rdb.Del(ctx, keys...).Err()
	}

	pipe := c.rdb.Pipeline()

	for _, currency := range currencies {
		primaryKey := c.GetPrimaryKey(currency.Id)
		isoKey := c.GetIsoKey(currency.IsoCode)
		symbolKey := c.GetSymbolKey(currency.Symbol)
		minorUnitsKey := c.GetMinorUnitsKey(currency.MinorUnits)

		data := map[string]interface{}{
			"id":          currency.Id,
			"name":        currency.Name,
			"symbol":      string(currency.Symbol),
			"iso_code":    currency.IsoCode,
			"minor_units": currency.MinorUnits,
		}

		pipe.HSet(ctx, primaryKey, data)
		pipe.SAdd(ctx, currenciesSetKey, strconv.FormatUint(currency.Id, 10))
		pipe.Set(ctx, isoKey, currency.Id, 0)
		pipe.SAdd(ctx, symbolKey, currency.Id)
		pipe.SAdd(ctx, minorUnitsKey, currency.Id)
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		return core.CacheSetError
	}

	return nil
}

func (c *CurrencyCache) UpdateById(ctx context.Context, currency *core.Currency) error {
	existing, err := c.GetById(ctx, currency.Id)
	if err != nil {
		if errors.Is(err, core.CacheGetError) {
			return core.CacheGetError
		}

		return core.InternalServerError
	}

	if existing == nil {
		return c.Set(ctx, currency)
	}

	if err := c.Delete(ctx, currency.Id); err != nil {
		if errors.Is(err, core.CacheDeleteError) {
			return core.CacheDeleteError
		}

		return core.InternalServerError
	}

	if err := c.Set(ctx, currency); err != nil {
		if errors.Is(err, core.CacheSetError) {
			return core.CacheSetError
		}

		return core.InternalServerError
	}

	return nil
}

func (c *CurrencyCache) Delete(ctx context.Context, id uint64) error {
	currency, err := c.GetById(ctx, id)
	if err != nil {
		return err
	}

	if currency == nil {
		return nil
	}

	primaryKey := c.GetPrimaryKey(id)
	isoKey := c.GetIsoKey(currency.IsoCode)
	symbolKey := c.GetSymbolKey(currency.Symbol)
	minorUnitsKey := c.GetMinorUnitsKey(currency.MinorUnits)

	pipe := c.rdb.Pipeline()

	pipe.Del(ctx, primaryKey)
	pipe.Del(ctx, isoKey)
	pipe.SRem(ctx, symbolKey, id)
	pipe.SRem(ctx, minorUnitsKey, id)
	pipe.SRem(ctx, currenciesSetKey, strconv.FormatUint(id, 10))

	_, err = pipe.Exec(ctx)
	if err != nil {
		return core.CacheDeleteError
	}

	return nil
}

func (c *CurrencyCache) MapToCurrency(data map[string]string) (core.Currency, error) {
	id, err := strconv.ParseUint(data["id"], 10, 64)
	if err != nil {
		return core.Currency{}, core.BadRequest
	}

	minorUnits, err := strconv.Atoi(data["minor_units"])
	if err != nil {
		return core.Currency{}, core.BadRequest
	}

	symbolStr := data["symbol"]
	var symbol rune
	if len(symbolStr) > 0 {
		runes := []rune(symbolStr)
		symbol = runes[0]
	}

	return core.Currency{
		Id:         id,
		Name:       data["name"],
		Symbol:     symbol,
		IsoCode:    data["iso_code"],
		MinorUnits: minorUnits,
	}, nil
}

func (c *CurrencyCache) GetPrimaryKey(id uint64) string {
	return currencyKeyPrefix + strconv.FormatUint(id, 10)
}

func (c *CurrencyCache) GetIsoKey(iso string) string {
	return currencyKeyPrefix + "iso:" + iso
}

func (c *CurrencyCache) GetSymbolKey(symbol rune) string {
	return currencyKeyPrefix + "symbol:" + string(symbol)
}

func (c *CurrencyCache) GetMinorUnitsKey(minorUnits int) string {
	return currencyKeyPrefix + "minor_units:" + strconv.Itoa(minorUnits)
}
