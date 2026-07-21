package currency

import (
	log "github.com/kVinsom/Bank-backend/internal/logging/service/currency"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CurrencyToCore maps a repository currency into the domain model.
func CurrencyToCore(currency *currencyRepository.Currency) *core.Currency {
	const operation = "currency_to_core"
	defer log.MappingStarted(operation)()
	if currency == nil {
		log.NilInput(operation)
		return nil
	}
	symbol := []rune(currency.Symbol)
	var value rune
	if len(symbol) > 0 {
		value = symbol[0]
	}
	return &core.Currency{
		Id:         currency.Id,
		Name:       currency.Name,
		Symbol:     value,
		IsoCode:    currency.IsoCode,
		MinorUnits: int8(currency.MinorUnits),
	}
}

// CurrencyToProto maps a domain currency into the repository contract.
func CurrencyToProto(currency *core.Currency) *currencyRepository.Currency {
	const operation = "currency_to_proto"
	defer log.MappingStarted(operation)()
	return &currencyRepository.Currency{
		Id:         currency.Id,
		Name:       currency.Name,
		Symbol:     string(currency.Symbol),
		IsoCode:    currency.IsoCode,
		MinorUnits: int32(currency.MinorUnits),
	}
}

// CurrenciesToCore maps repository currencies while preserving their order.
func CurrenciesToCore(currencies []*currencyRepository.Currency) []core.Currency {
	const operation = "currencies_to_core"
	defer log.MappingStarted(operation)()
	result := make([]core.Currency, 0, len(currencies))
	for _, currency := range currencies {
		if converted := CurrencyToCore(currency); converted != nil {
			result = append(result, *converted)
		}
	}
	return result
}
