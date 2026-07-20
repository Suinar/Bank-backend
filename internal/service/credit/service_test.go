package credit

import "testing"

func TestCreditServiceImplementsInterface(t *testing.T) {
	var service ICreditService = &CreditService{}
	if service == nil {
		t.Fatal("expected service implementation")
	}
}
