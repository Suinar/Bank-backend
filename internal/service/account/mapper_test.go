package account

import (
	"reflect"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
)

func TestAccountToCore(t *testing.T) {
	t.Parallel()

	want := fixture.AccountCore()
	got := AccountToCore(fixture.AccountProto())

	if !reflect.DeepEqual(got, &want) {
		t.Fatalf("result: want %#v, got %#v", &want, got)
	}
}

func TestAccountToCore_Nil(t *testing.T) {
	t.Parallel()

	if got := AccountToCore(nil); got != nil {
		t.Fatalf("result: want nil, got %#v", got)
	}
}

func TestAccountToProto(t *testing.T) {
	t.Parallel()

	input := fixture.AccountCore()
	want := fixture.AccountProto()
	got := AccountToProto(&input)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("result: want %#v, got %#v", want, got)
	}
}

func TestAccountsToCore(t *testing.T) {
	t.Parallel()

	second := fixture.AccountProto()
	second.Id++
	second.Name = "Second account"

	got := AccountsToCore([]*accountRepository.Account{
		fixture.AccountProto(),
		nil,
		second,
	})
	want := fixture.AccountListCore()
	want = append(want, *AccountToCore(second))

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("result: want %#v, got %#v", want, got)
	}
}

func TestAccountsToCore_Nil(t *testing.T) {
	t.Parallel()

	got := AccountsToCore(nil)
	if got == nil {
		t.Fatal("result: want empty non-nil slice, got nil")
	}
	if len(got) != 0 {
		t.Fatalf("length: want 0, got %d", len(got))
	}
}
