package deposit

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	"github.com/kVinsom/Bank-proto/repository/common"
	depositRepository "github.com/kVinsom/Bank-proto/repository/deposit"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

func TestDepositService_Success(t *testing.T) {
	t.Parallel()
	closed := fixture.DepositProto()
	closed.Status = depositRepository.DepositStatus(core.DepositStatusClosed)
	tests := []struct {
		name   string
		expect func(*fixture.DepositRepositoryMocks)
		invoke func(*DepositService) (any, error)
		want   any
	}{
		{name: "get all", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(fixture.DepositListProto(fixture.DepositProto(), closed), nil)
		}, invoke: func(s *DepositService) (any, error) { return s.GetAll(context.Background()) }, want: fixture.DepositListCore()},
		{name: "get by user", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().GetByUser(gomock.Any(), gomock.Any()).Return(fixture.DepositListProto(fixture.DepositProto()), nil)
		}, invoke: func(s *DepositService) (any, error) { return s.GetByUser(context.Background(), fixture.UserId) }, want: fixture.DepositListCore()},
		{name: "get by id", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.DepositProto(), nil)
		}, invoke: func(s *DepositService) (any, error) { return s.GetById(context.Background(), fixture.DepositId) }, want: depositCorePointer()},
		{name: "create", expect: expectDepositCreateSuccess, invoke: func(s *DepositService) (any, error) {
			input := fixture.DepositCreateInputCore()
			return s.Create(context.Background(), &input)
		}, want: depositCorePointer()},
		{name: "replenish", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().Replenish(gomock.Any(), &common.AmountRequest{Id: fixture.DepositId, Amount: fixture.TransactionAmount}).Return(fixture.DepositProto(), nil)
		}, invoke: func(s *DepositService) (any, error) {
			return nil, s.Replenish(context.Background(), fixture.DepositId, fixture.TransactionAmount)
		}, want: nil},
		{name: "delete", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&common.Empty{}, nil)
		}, invoke: func(s *DepositService) (any, error) { return nil, s.Delete(context.Background(), fixture.DepositId) }, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newDepositServiceSUT(t)
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

func expectDepositCreateSuccess(m *fixture.DepositRepositoryMocks) {
	m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
	m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.CurrencyProto(), nil)
	m.Deposit.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fixture.DepositProto(), nil)
}
func newDepositServiceSUT(t *testing.T) (*fixture.DepositRepositoryMocks, *DepositService) {
	t.Helper()
	mocks := fixture.NewDepositRepositoryMocks(t)
	return mocks, NewDepositService(mocks.Deposit, mocks.User, mocks.Currency)
}
func depositCorePointer() any { value := fixture.DepositCore(); return &value }
