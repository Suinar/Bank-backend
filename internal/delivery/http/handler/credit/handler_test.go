package credit

import "testing"

func TestNewCreditHandler(t *testing.T) {
	handler := NewCreditHandler(nil)
	if handler == nil {
		t.Fatal("expected handler instance")
	}
}
