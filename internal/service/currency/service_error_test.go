package currency

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

var currencyRepositoryError = errors.New("repository failure")

func TestCurrencyService_InvalidInput(t *testing.T) {
	t.Parallel()
	invalidCreates := []*core.CurrencyCreateInput{nil, {Name: "", Symbol: '$', IsoCode: "USD", MinorUnits: 2}, {Name: "Dollar", Symbol: '0', IsoCode: "USD", MinorUnits: 2}, {Name: "Dollar", Symbol: '$', IsoCode: "", MinorUnits: 2}, {Name: "Dollar", Symbol: '$', IsoCode: "USD", MinorUnits: 0}, {Name: "Dollar", Symbol: '$', IsoCode: "USD", MinorUnits: 51}}
	for _, input := range invalidCreates {
		got, err := (&CurrencyService{}).Create(context.Background(), input)
		assertCurrencyError(t, got, err, coreErrors.BadRequest)
	}
	empty := ""
	badISO := "EURO"
	zero := int8(0)
	invalidUpdates := []*core.CurrencyUpdateInput{nil, {Name: &empty}, {IsoCode: &empty}, {IsoCode: &badISO}, {MinorUnits: &zero}}
	for _, input := range invalidUpdates {
		got, err := (&CurrencyService{}).Update(context.Background(), fixture.CurrencyId, input)
		assertCurrencyError(t, got, err, coreErrors.BadRequest)
	}
}

func TestCurrencyService_RepositoryErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.CurrencyRepositoryMocks)
		invoke func(*CurrencyService) (any, error)
		want   error
	}{
		{name: "get all", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, currencyRepositoryError)
		}, invoke: func(s *CurrencyService) (any, error) { return s.GetAll(context.Background()) }, want: coreErrors.InternalServerError},
		{name: "get by id not found", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CurrencyService) (any, error) { return s.GetById(context.Background(), fixture.CurrencyId) }, want: coreErrors.NotFound},
		{name: "get by iso", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetByIso(gomock.Any(), gomock.Any()).Return(nil, currencyRepositoryError)
		}, invoke: func(s *CurrencyService) (any, error) {
			return s.GetByIso(context.Background(), fixture.CurrencyISOCode)
		}, want: coreErrors.InternalServerError},
		{name: "get by symbol", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetBySymbol(gomock.Any(), gomock.Any()).Return(nil, currencyRepositoryError)
		}, invoke: func(s *CurrencyService) (any, error) {
			return s.GetBySymbol(context.Background(), fixture.CurrencySymbol)
		}, want: coreErrors.InternalServerError},
		{name: "create", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, currencyRepositoryError)
		}, invoke: func(s *CurrencyService) (any, error) {
			input := fixture.CurrencyCreateInputCore()
			return s.Create(context.Background(), &input)
		}, want: coreErrors.InternalServerError},
		{name: "update not found", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CurrencyService) (any, error) {
			input := fixture.CurrencyUpdateInputCore()
			return s.Update(context.Background(), fixture.CurrencyId, &input)
		}, want: coreErrors.NotFound},
		{name: "delete", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, currencyRepositoryError)
		}, invoke: func(s *CurrencyService) (any, error) { return nil, s.Delete(context.Background(), fixture.CurrencyId) }, want: coreErrors.InternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCurrencyServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertCurrencyError(t, got, err, tt.want)
		})
	}
}

func assertCurrencyError(t *testing.T, got any, err, want error) {
	t.Helper()
	if got != nil && !reflect.ValueOf(got).IsNil() {
		t.Fatalf("result: want nil, got %#v", got)
	}
	if !errors.Is(err, want) {
		t.Fatalf("error: want %v, got %v", want, err)
	}
}

func TestCurrencyService_RepositoryErrorBranches(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.CurrencyRepositoryMocks)
		invoke func(*CurrencyService) (any, error)
		want   error
	}{
		{name: "get by id failure", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, currencyRepositoryError)
		}, invoke: func(s *CurrencyService) (any, error) { return s.GetById(context.Background(), fixture.CurrencyId) }, want: coreErrors.InternalServerError},
		{name: "get by iso not found", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetByIso(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CurrencyService) (any, error) {
			return s.GetByIso(context.Background(), fixture.CurrencyISOCode)
		}, want: coreErrors.NotFound},
		{name: "get by symbol not found", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetBySymbol(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CurrencyService) (any, error) {
			return s.GetBySymbol(context.Background(), fixture.CurrencySymbol)
		}, want: coreErrors.NotFound},
		{name: "update failure", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, currencyRepositoryError)
		}, invoke: func(s *CurrencyService) (any, error) {
			input := fixture.CurrencyUpdateInputCore()
			return s.Update(context.Background(), fixture.CurrencyId, &input)
		}, want: coreErrors.InternalServerError},
		{name: "delete not found", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CurrencyService) (any, error) { return nil, s.Delete(context.Background(), fixture.CurrencyId) }, want: coreErrors.NotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCurrencyServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertCurrencyError(t, got, err, tt.want)
		})
	}
}
