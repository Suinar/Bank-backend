package credit

import (
	"reflect"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	creditRepository "github.com/kVinsom/Bank-proto/repository/credit"
)

func TestCreditMappers(t *testing.T) {
	t.Parallel()
	coreValue := fixture.CreditCore()
	if got := CreditToCore(fixture.CreditProto()); !reflect.DeepEqual(got, &coreValue) {
		t.Fatalf("CreditToCore: want %#v, got %#v", &coreValue, got)
	}
	if got := CreditToProto(&coreValue); !reflect.DeepEqual(got, fixture.CreditProto()) {
		t.Fatalf("CreditToProto: want %#v, got %#v", fixture.CreditProto(), got)
	}
	if got := CreditToCore(nil); got != nil {
		t.Fatalf("CreditToCore(nil): got %#v", got)
	}
	if got := CreditsToCore([]*creditRepository.Credit{fixture.CreditProto(), nil}); !reflect.DeepEqual(got, fixture.CreditListCore()) {
		t.Fatalf("CreditsToCore: want %#v, got %#v", fixture.CreditListCore(), got)
	}
}
