package account

import "testing"

func TestAccountServiceImplementsInterface(t *testing.T) {
	var service IAccountService = &AccountService{}
	if service == nil {
		t.Fatal("expected service implementation")
	}
}
