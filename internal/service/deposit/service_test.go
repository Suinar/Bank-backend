package deposit

import "testing"

func TestDepositServiceImplementsInterface(t *testing.T) {
	var service IDepositService = &DepositService{}
	if service == nil {
		t.Fatal("expected service implementation")
	}
}
