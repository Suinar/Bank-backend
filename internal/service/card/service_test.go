package card

import "testing"

func TestCardServiceImplementsInterface(t *testing.T) {
	var service ICardService = &CardService{}
	if service == nil {
		t.Fatal("expected service implementation")
	}
}
