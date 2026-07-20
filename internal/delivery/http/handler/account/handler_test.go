package account

import "testing"

func TestNewAccountHandler(t *testing.T) {
	handler := NewAccountHandler(nil)
	if handler == nil {
		t.Fatal("expected handler instance")
	}
}
