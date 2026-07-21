package fixture

import (
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CurrencyCore returns a valid currency fixture.
func CurrencyCore() core.Currency {
	return core.Currency{
		Id:         CurrencyId,
		Name:       CurrencyName,
		Symbol:     CurrencySymbol,
		IsoCode:    CurrencyISOCode,
		MinorUnits: CurrencyMinorUnits,
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
		Name:       CurrencyName,
		Symbol:     CurrencySymbol,
		IsoCode:    CurrencyISOCode,
		MinorUnits: CurrencyMinorUnits,
	}
}

// CurrencyUpdateInputCore returns a valid currency update fixture.
func CurrencyUpdateInputCore() core.CurrencyUpdateInput {
	return core.CurrencyUpdateInput{
		Name:       StringPointer(CurrencyName),
		Symbol:     RunePointer(CurrencySymbol),
		IsoCode:    StringPointer(CurrencyISOCode),
		MinorUnits: Int8Pointer(CurrencyMinorUnits),
	}
}
