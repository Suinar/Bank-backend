package deposit

import (
	"reflect"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	depositRepository "github.com/kVinsom/Bank-proto/repository/deposit"
)

func TestDepositMappers(t *testing.T) {
	t.Parallel()
	coreValue := fixture.DepositCore()
	if got := DepositToCore(fixture.DepositProto()); !reflect.DeepEqual(got, &coreValue) {
		t.Fatalf("DepositToCore: want %#v, got %#v", &coreValue, got)
	}
	if got := DepositToProto(&coreValue); !reflect.DeepEqual(got, fixture.DepositProto()) {
		t.Fatalf("DepositToProto: want %#v, got %#v", fixture.DepositProto(), got)
	}
	if got := DepositToCore(nil); got != nil {
		t.Fatalf("DepositToCore(nil): got %#v", got)
	}
	if got := DepositsToCore([]*depositRepository.Deposit{fixture.DepositProto(), nil}); !reflect.DeepEqual(got, fixture.DepositListCore()) {
		t.Fatalf("DepositsToCore: want %#v, got %#v", fixture.DepositListCore(), got)
	}
}
