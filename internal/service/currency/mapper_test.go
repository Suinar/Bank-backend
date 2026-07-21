package currency

import (
	"reflect"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
)

func TestCurrencyMappers(t *testing.T) {
	t.Parallel()
	coreValue := fixture.CurrencyCore()
	if got := CurrencyToCore(fixture.CurrencyProto()); !reflect.DeepEqual(got, &coreValue) {
		t.Fatalf("CurrencyToCore: want %#v, got %#v", &coreValue, got)
	}
	if got := CurrencyToProto(&coreValue); !reflect.DeepEqual(got, fixture.CurrencyProto()) {
		t.Fatalf("CurrencyToProto: want %#v, got %#v", fixture.CurrencyProto(), got)
	}
	if got := CurrencyToCore(nil); got != nil {
		t.Fatalf("CurrencyToCore(nil): want nil, got %#v", got)
	}
	if got := CurrenciesToCore([]*currencyRepository.Currency{fixture.CurrencyProto(), nil}); !reflect.DeepEqual(got, fixture.CurrencyListCore()) {
		t.Fatalf("CurrenciesToCore: want %#v, got %#v", fixture.CurrencyListCore(), got)
	}
}
