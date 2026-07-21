package exhange_rate

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	exchangeRateProto "github.com/kVinsom/Bank-proto/exchange_rate"
)

func TestExchangeRateService_Success(t *testing.T) {
	t.Parallel()

	t.Run("get all", func(t *testing.T) {
		mocks, sut := newExchangeRateServiceSUT(t)
		mocks.Ranking.EXPECT().GetAllRanking(gomock.Any(), &exchangeRateProto.GetAllRankingRequest{CurrencyIsoFrom: fixture.ExchangeISOFrom}).Return(fixture.RankingListProto(fixture.RankingProto()), nil)
		got, err := sut.GetAllRanking(context.Background(), fixture.ExchangeISOFrom)
		if err != nil || !reflect.DeepEqual(got, fixture.RankingListCore()) {
			t.Fatalf("want %#v and nil error, got %#v and %v", fixture.RankingListCore(), got, err)
		}
	})

	t.Run("get relative", func(t *testing.T) {
		mocks, sut := newExchangeRateServiceSUT(t)
		mocks.Ranking.EXPECT().GetRelativeRanking(gomock.Any(), &exchangeRateProto.GetRelativeRankingRequest{CurrencyIsoFrom: fixture.ExchangeISOFrom, CurrencyIsoTo: fixture.ExchangeISOTo}).Return(fixture.RankingProto(), nil)
		got, err := sut.GetRelativeRanking(context.Background(), fixture.ExchangeISOFrom, fixture.ExchangeISOTo)
		want := fixture.RankingCore()
		if err != nil || !reflect.DeepEqual(got, &want) {
			t.Fatalf("want %#v and nil error, got %#v and %v", &want, got, err)
		}
	})
}

func TestValidCurrencyISO(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		iso  int
		want bool
	}{{"minimum", 1, true}, {"common code", 840, true}, {"maximum", 999, true}, {"negative", -1, false}, {"zero", 0, false}, {"above maximum", 1000, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validCurrencyIso(tt.iso); got != tt.want {
				t.Fatalf("validCurrencyIso(%d): want %v, got %v", tt.iso, tt.want, got)
			}
		})
	}
}

func newExchangeRateServiceSUT(t *testing.T) (*fixture.ExchangeRateRepositoryMocks, *ExchangeRateService) {
	t.Helper()
	mocks := fixture.NewExchangeRateRepositoryMocks(t)
	return mocks, NewExchangeRateService(mocks.Ranking)
}
