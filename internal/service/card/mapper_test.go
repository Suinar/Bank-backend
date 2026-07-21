package card

import (
	"reflect"
	"testing"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	cardRepository "github.com/kVinsom/Bank-proto/repository/card"
)

func TestCardMappers(t *testing.T) {
	t.Parallel()
	coreValue := fixture.CardCore()
	if got := CardToCore(fixture.CardProto()); !reflect.DeepEqual(got, &coreValue) {
		t.Fatalf("CardToCore: want %#v, got %#v", &coreValue, got)
	}
	if got := CardToProto(&coreValue); !reflect.DeepEqual(got, fixture.CardProto()) {
		t.Fatalf("CardToProto: want %#v, got %#v", fixture.CardProto(), got)
	}
	if got := CardToCore(nil); got != nil {
		t.Fatalf("CardToCore(nil): got %#v", got)
	}
	if got := CardsToCore([]*cardRepository.Card{fixture.CardProto(), nil}); !reflect.DeepEqual(got, fixture.CardListCore()) {
		t.Fatalf("CardsToCore: want %#v, got %#v", fixture.CardListCore(), got)
	}
}
