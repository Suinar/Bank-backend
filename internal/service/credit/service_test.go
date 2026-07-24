package credit

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	"github.com/kVinsom/Bank-proto/repository/common"
	creditRepository "github.com/kVinsom/Bank-proto/repository/credit"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

func TestCreditService_Success(t *testing.T) {
	t.Parallel()
	closed := fixture.CreditProto()
	closed.Status = creditRepository.CreditStatus(core.CreditStatusClosed)
	tests := []struct {
		name   string
		expect func(*fixture.CreditRepositoryMocks)
		invoke func(*CreditService) (any, error)
		want   any
	}{
		{name: "get all", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(fixture.CreditListProto(fixture.CreditProto(), closed), nil)
		}, invoke: func(s *CreditService) (any, error) { return s.GetAll(context.Background()) }, want: fixture.CreditListCore()},
		{name: "get by user", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().GetByUser(gomock.Any(), gomock.Any()).Return(fixture.CreditListProto(fixture.CreditProto()), nil)
		}, invoke: func(s *CreditService) (any, error) { return s.GetByUser(context.Background(), test.UserId) }, want: fixture.CreditListCore()},
		{name: "get by id", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.CreditProto(), nil)
		}, invoke: func(s *CreditService) (any, error) { return s.GetById(context.Background(), test.CreditId) }, want: creditCorePointer()},
		{name: "create", expect: expectCreditCreateSuccess, invoke: func(s *CreditService) (any, error) {
			input := fixture.CreditCreateInputCore()
			return s.Create(context.Background(), &input)
		}, want: creditCorePointer()},
		{name: "repay", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().Repay(gomock.Any(), &common.AmountRequest{Id: test.CreditId, Amount: test.TransactionAmount}).Return(fixture.CreditProto(), nil)
		}, invoke: func(s *CreditService) (any, error) {
			return nil, s.Repay(context.Background(), test.CreditId, test.TransactionAmount)
		}, want: nil},
		{name: "delete", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&common.Empty{}, nil)
		}, invoke: func(s *CreditService) (any, error) { return nil, s.Delete(context.Background(), test.CreditId) }, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCreditServiceSUT(t)
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

func expectCreditCreateSuccess(m *fixture.CreditRepositoryMocks) {
	m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
	m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.CurrencyProto(), nil)
	m.Credit.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fixture.CreditProto(), nil)
}
func newCreditServiceSUT(t *testing.T) (*fixture.CreditRepositoryMocks, *CreditService) {
	t.Helper()
	mocks := fixture.NewCreditRepositoryMocks(t)
	return mocks, NewCreditService(mocks.Credit, mocks.User, mocks.Currency)
}
func creditCorePointer() any { value := fixture.CreditCore(); return &value }
