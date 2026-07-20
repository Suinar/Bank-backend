package user

import "testing"

func TestNewUserHandler(t *testing.T) {
	handler := NewUserHandler(nil)
	if handler == nil {
		t.Fatal("expected handler instance")
	}
}
