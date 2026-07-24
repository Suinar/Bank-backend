package currency

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	"github.com/kVinsom/Bank-proto/repository/common"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
)

func TestCurrencyService_Success(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.CurrencyRepositoryMocks)
		invoke func(*CurrencyService) (any, error)
		want   any
	}{
		{name: "get all", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetAll(gomock.Any(), &common.Empty{}).Return(fixture.CurrencyListProto(fixture.CurrencyProto()), nil)
		}, invoke: func(s *CurrencyService) (any, error) { return s.GetAll(context.Background()) }, want: fixture.CurrencyListCore()},
		{name: "get by id", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetById(gomock.Any(), &common.IdRequest{Id: test.CurrencyId}).Return(fixture.CurrencyProto(), nil)
		}, invoke: func(s *CurrencyService) (any, error) { return s.GetById(context.Background(), test.CurrencyId) }, want: currencyCorePointer()},
		{name: "get by iso", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetByIso(gomock.Any(), &currencyRepository.IsoCodeRequest{IsoCode: test.CurrencyISOCode}).Return(fixture.CurrencyProto(), nil)
		}, invoke: func(s *CurrencyService) (any, error) {
			return s.GetByIso(context.Background(), test.CurrencyISOCode)
		}, want: currencyCorePointer()},
		{name: "get by symbol", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().GetBySymbol(gomock.Any(), &currencyRepository.SymbolRequest{Symbol: string(test.CurrencySymbol)}).Return(fixture.CurrencyProto(), nil)
		}, invoke: func(s *CurrencyService) (any, error) {
			return s.GetBySymbol(context.Background(), test.CurrencySymbol)
		}, want: currencyCorePointer()},
		{name: "create", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fixture.CurrencyProto(), nil)
		}, invoke: func(s *CurrencyService) (any, error) {
			input := fixture.CurrencyCreateInputCore()
			return s.Create(context.Background(), &input)
		}, want: currencyCorePointer()},
		{name: "update", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().Update(gomock.Any(), gomock.Any()).Return(fixture.CurrencyProto(), nil)
		}, invoke: func(s *CurrencyService) (any, error) {
			input := fixture.CurrencyUpdateInputCore()
			return s.Update(context.Background(), test.CurrencyId, &input)
		}, want: currencyCorePointer()},
		{name: "delete", expect: func(m *fixture.CurrencyRepositoryMocks) {
			m.Currency.EXPECT().Delete(gomock.Any(), &common.IdRequest{Id: test.CurrencyId}).Return(&common.Empty{}, nil)
		}, invoke: func(s *CurrencyService) (any, error) { return nil, s.Delete(context.Background(), test.CurrencyId) }, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCurrencyServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("result: want %#v, got %#v", tt.want, got)
			}
		})
	}
}

func newCurrencyServiceSUT(t *testing.T) (*fixture.CurrencyRepositoryMocks, *CurrencyService) {
	t.Helper()
	mocks := fixture.NewCurrencyRepositoryMocks(t)
	return mocks, NewCurrencyService(mocks.Currency)
}
func currencyCorePointer() any { value := fixture.CurrencyCore(); return &value }
