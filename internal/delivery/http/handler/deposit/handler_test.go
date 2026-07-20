package deposit

import "testing"

func TestNewDepositHandler(t *testing.T) {
	handler := NewDepositHandler(nil)
	if handler == nil {
		t.Fatal("expected handler instance")
	}
}
