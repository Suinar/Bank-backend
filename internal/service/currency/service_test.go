package currency

import "testing"

func TestCurrencyServiceImplementsInterface(t *testing.T) {
	var service ICurrencyService = &CurrencyService{}
	if service == nil {
		t.Fatal("expected service implementation")
	}
}
