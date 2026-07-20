package currency

import "testing"

func TestNewCurrencyHandler(t *testing.T) {
	handler := NewCurrencyHandler(nil, nil, nil)
	if handler == nil {
		t.Fatal("expected handler instance")
	}
}
