package card

import "testing"

func TestNewCardHandler(t *testing.T) {
	handler := NewCardHandler(nil)
	if handler == nil {
		t.Fatal("expected handler instance")
	}
}
