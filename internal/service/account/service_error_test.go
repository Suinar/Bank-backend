package account

import (
	"context"
	"errors"
	"testing"

	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
)

func TestAccountService_Create_NilInputError(t *testing.T) {
	service := &AccountService{}
	result, err := service.Create(context.Background(), nil)

	if result != nil {
		t.Fatalf("expected nil result, got %#v", result)
	}
	if !errors.Is(err, coreErrors.BadRequest) {
		t.Fatalf("expected BadRequest, got %v", err)
	}
}
