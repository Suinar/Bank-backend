package exhange_rate

import (
	"reflect"
	"testing"
	"time"

	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestRankingToCore(t *testing.T) {
	t.Parallel()
	want := fixture.RankingCore()
	if got := RankingToCore(fixture.RankingProto()); !reflect.DeepEqual(got, &want) {
		t.Fatalf("want %#v, got %#v", &want, got)
	}
	if got := RankingToCore(nil); got != nil {
		t.Fatalf("nil input: want nil, got %#v", got)
	}
	withoutDate := fixture.RankingProto()
	withoutDate.Date = nil
	got := RankingToCore(withoutDate)
	if !got.Date.Equal(time.Time{}) {
		t.Fatalf("nil date: want zero time, got %v", got.Date)
	}
}
