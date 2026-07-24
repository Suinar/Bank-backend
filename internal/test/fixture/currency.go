package fixture

import (
	"github.com/kVinsom/Bank-backend/internal/test"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CurrencyCore returns a valid currency fixture.
func CurrencyCore() core.Currency {
	return core.Currency{
		Id:         test.CurrencyId,
		Name:       test.CurrencyName,
		Symbol:     test.CurrencySymbol,
		IsoCode:    test.CurrencyISOCode,
		MinorUnits: test.CurrencyMinorUnits,
	}
}

func CurrencyListCore() []core.Currency { return []core.Currency{CurrencyCore()} }

func CurrencyProto() *currencyRepository.Currency {
	currency := CurrencyCore()
	return &currencyRepository.Currency{
		Id:         currency.Id,
		Name:       currency.Name,
		Symbol:     string(currency.Symbol),
		IsoCode:    currency.IsoCode,
		MinorUnits: int32(currency.MinorUnits),
	}
}

func CurrencyListProto(currencies ...*currencyRepository.Currency) *currencyRepository.CurrencyList {
	return &currencyRepository.CurrencyList{
		Currencies: currencies,
	}
}

// CurrencyCreateInputCore returns a valid currency creation fixture.
func CurrencyCreateInputCore() core.CurrencyCreateInput {
	return core.CurrencyCreateInput{
		Name:       test.CurrencyName,
		Symbol:     test.CurrencySymbol,
		IsoCode:    test.CurrencyISOCode,
		MinorUnits: test.CurrencyMinorUnits,
	}
}

// CurrencyUpdateInputCore returns a valid currency update fixture.
func CurrencyUpdateInputCore() core.CurrencyUpdateInput {
	return core.CurrencyUpdateInput{
		Name:       StringPointer(test.CurrencyName),
		Symbol:     RunePointer(test.CurrencySymbol),
		IsoCode:    StringPointer(test.CurrencyISOCode),
		MinorUnits: Int8Pointer(test.CurrencyMinorUnits),
	}
}
