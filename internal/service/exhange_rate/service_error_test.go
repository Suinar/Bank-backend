package exhange_rate

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	exchangeRateProto "github.com/kVinsom/Bank-proto/exchange_rate"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
)

var exchangeRateRepositoryError = errors.New("repository failure")

func TestExchangeRateService_InvalidDependenciesAndInput(t *testing.T) {
	t.Parallel()
	missing := NewExchangeRateService(nil)
	assertExchangeRateError(t, func() (any, error) { return missing.GetAllRanking(context.Background(), test.ExchangeISOFrom) }, coreErrors.InternalServerError)
	assertExchangeRateError(t, func() (any, error) {
		return missing.GetRelativeRanking(context.Background(), test.ExchangeISOFrom, test.ExchangeISOTo)
	}, coreErrors.InternalServerError)
	_, sut := newExchangeRateServiceSUT(t)
	for _, iso := range []int{-1, 0, 1000} {
		iso := iso
		assertExchangeRateError(t, func() (any, error) { return sut.GetAllRanking(context.Background(), iso) }, coreErrors.BadRequest)
	}
	assertExchangeRateError(t, func() (any, error) { return sut.GetRelativeRanking(context.Background(), 0, test.ExchangeISOTo) }, coreErrors.BadRequest)
	assertExchangeRateError(t, func() (any, error) {
		return sut.GetRelativeRanking(context.Background(), test.ExchangeISOFrom, 1000)
	}, coreErrors.BadRequest)
}

func TestExchangeRateService_RepositoryErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.ExchangeRateRepositoryMocks)
		invoke func(*ExchangeRateService) (any, error)
		want   error
	}{
		{name: "all bad request", expect: func(m *fixture.ExchangeRateRepositoryMocks) {
			m.Ranking.EXPECT().GetAllRanking(gomock.Any(), gomock.Any()).Return(nil, coreErrors.BadRequest)
		}, invoke: func(s *ExchangeRateService) (any, error) {
			return s.GetAllRanking(context.Background(), test.ExchangeISOFrom)
		}, want: coreErrors.BadRequest},
		{name: "all failure", expect: func(m *fixture.ExchangeRateRepositoryMocks) {
			m.Ranking.EXPECT().GetAllRanking(gomock.Any(), gomock.Any()).Return(nil, exchangeRateRepositoryError)
		}, invoke: func(s *ExchangeRateService) (any, error) {
			return s.GetAllRanking(context.Background(), test.ExchangeISOFrom)
		}, want: coreErrors.InternalServerError},
		{name: "all nil response", expect: func(m *fixture.ExchangeRateRepositoryMocks) {
			m.Ranking.EXPECT().GetAllRanking(gomock.Any(), gomock.Any()).Return((*exchangeRateProto.RankingList)(nil), nil)
		}, invoke: func(s *ExchangeRateService) (any, error) {
			return s.GetAllRanking(context.Background(), test.ExchangeISOFrom)
		}, want: coreErrors.InternalServerError},
		{name: "relative bad request", expect: func(m *fixture.ExchangeRateRepositoryMocks) {
			m.Ranking.EXPECT().GetRelativeRanking(gomock.Any(), gomock.Any()).Return(nil, coreErrors.BadRequest)
		}, invoke: func(s *ExchangeRateService) (any, error) {
			return s.GetRelativeRanking(context.Background(), test.ExchangeISOFrom, test.ExchangeISOTo)
		}, want: coreErrors.BadRequest},
		{name: "relative not found", expect: func(m *fixture.ExchangeRateRepositoryMocks) {
			m.Ranking.EXPECT().GetRelativeRanking(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *ExchangeRateService) (any, error) {
			return s.GetRelativeRanking(context.Background(), test.ExchangeISOFrom, test.ExchangeISOTo)
		}, want: coreErrors.NotFound},
		{name: "relative failure", expect: func(m *fixture.ExchangeRateRepositoryMocks) {
			m.Ranking.EXPECT().GetRelativeRanking(gomock.Any(), gomock.Any()).Return(nil, exchangeRateRepositoryError)
		}, invoke: func(s *ExchangeRateService) (any, error) {
			return s.GetRelativeRanking(context.Background(), test.ExchangeISOFrom, test.ExchangeISOTo)
		}, want: coreErrors.InternalServerError},
		{name: "relative nil response", expect: func(m *fixture.ExchangeRateRepositoryMocks) {
			m.Ranking.EXPECT().GetRelativeRanking(gomock.Any(), gomock.Any()).Return(nil, nil)
		}, invoke: func(s *ExchangeRateService) (any, error) {
			return s.GetRelativeRanking(context.Background(), test.ExchangeISOFrom, test.ExchangeISOTo)
		}, want: coreErrors.NotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newExchangeRateServiceSUT(t)
			tt.expect(mocks)
			assertExchangeRateError(t, func() (any, error) { return tt.invoke(sut) }, tt.want)
		})
	}
}

func assertExchangeRateError(t *testing.T, invoke func() (any, error), want error) {
	t.Helper()
	got, err := invoke()
	if got != nil && !reflect.ValueOf(got).IsNil() {
		t.Fatalf("result: want nil, got %#v", got)
	}
	if !errors.Is(err, want) {
		t.Fatalf("error: want %v, got %v", want, err)
	}
}
