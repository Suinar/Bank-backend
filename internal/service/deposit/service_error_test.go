package deposit

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

var depositRepositoryError = errors.New("repository failure")

func TestDepositService_Create_InvalidInput(t *testing.T) {
	t.Parallel()
	inputs := []*core.DepositCreateInput{nil, {UserId: 0, CurrencyId: 1, Amount: 1000, TermMonths: 1}, {UserId: 1, CurrencyId: 0, Amount: 1000, TermMonths: 1}, {UserId: 1, CurrencyId: 1, Amount: 999, TermMonths: 1}, {UserId: 1, CurrencyId: 1, Amount: 10001, TermMonths: 1}, {UserId: 1, CurrencyId: 1, Amount: 1000, TermMonths: 0}, {UserId: 1, CurrencyId: 1, Amount: 1000, TermMonths: 25}}
	for _, input := range inputs {
		got, err := (&DepositService{}).Create(context.Background(), input)
		assertDepositError(t, got, err, coreErrors.BadRequest)
	}
}

func TestDepositService_RepositoryErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.DepositRepositoryMocks)
		invoke func(*DepositService) (any, error)
		want   error
	}{
		{name: "get all", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, depositRepositoryError)
		}, invoke: func(s *DepositService) (any, error) { return s.GetAll(context.Background()) }, want: coreErrors.InternalServerError},
		{name: "get by user not found", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().GetByUser(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *DepositService) (any, error) { return s.GetByUser(context.Background(), fixture.UserId) }, want: coreErrors.NotFound},
		{name: "get by id", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, depositRepositoryError)
		}, invoke: func(s *DepositService) (any, error) { return s.GetById(context.Background(), fixture.DepositId) }, want: coreErrors.InternalServerError},
		{name: "replenish not found", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().Replenish(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *DepositService) (any, error) {
			return nil, s.Replenish(context.Background(), fixture.DepositId, fixture.TransactionAmount)
		}, want: coreErrors.NotFound},
		{name: "delete", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, depositRepositoryError)
		}, invoke: func(s *DepositService) (any, error) { return nil, s.Delete(context.Background(), fixture.DepositId) }, want: coreErrors.InternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newDepositServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertDepositError(t, got, err, tt.want)
		})
	}
}

func TestDepositService_Create_DependencyErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		setup func(*fixture.DepositRepositoryMocks)
		want  error
	}{
		{name: "user not found", setup: func(m *fixture.DepositRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, want: coreErrors.BadRequest},
		{name: "user failure", setup: func(m *fixture.DepositRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, depositRepositoryError)
		}, want: coreErrors.InternalServerError},
		{name: "currency not found", setup: func(m *fixture.DepositRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
			m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, want: coreErrors.NotFound},
		{name: "currency failure", setup: func(m *fixture.DepositRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
			m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, depositRepositoryError)
		}, want: coreErrors.InternalServerError},
		{name: "create failure", setup: func(m *fixture.DepositRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
			m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.CurrencyProto(), nil)
			m.Deposit.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, depositRepositoryError)
		}, want: coreErrors.InternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newDepositServiceSUT(t)
			tt.setup(mocks)
			input := fixture.DepositCreateInputCore()
			got, err := sut.Create(context.Background(), &input)
			assertDepositError(t, got, err, tt.want)
		})
	}
}

func assertDepositError(t *testing.T, got any, err, want error) {
	t.Helper()
	if got != nil && !reflect.ValueOf(got).IsNil() {
		t.Fatalf("result: want nil, got %#v", got)
	}
	if !errors.Is(err, want) {
		t.Fatalf("error: want %v, got %v", want, err)
	}
}

func TestDepositService_RepositoryErrorBranches(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.DepositRepositoryMocks)
		invoke func(*DepositService) (any, error)
		want   error
	}{
		{name: "get by user failure", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().GetByUser(gomock.Any(), gomock.Any()).Return(nil, depositRepositoryError)
		}, invoke: func(s *DepositService) (any, error) { return s.GetByUser(context.Background(), fixture.UserId) }, want: coreErrors.InternalServerError},
		{name: "get by id not found", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *DepositService) (any, error) { return s.GetById(context.Background(), fixture.DepositId) }, want: coreErrors.NotFound},
		{name: "replenish failure", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().Replenish(gomock.Any(), gomock.Any()).Return(nil, depositRepositoryError)
		}, invoke: func(s *DepositService) (any, error) {
			return nil, s.Replenish(context.Background(), fixture.DepositId, fixture.TransactionAmount)
		}, want: coreErrors.InternalServerError},
		{name: "delete not found", expect: func(m *fixture.DepositRepositoryMocks) {
			m.Deposit.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *DepositService) (any, error) { return nil, s.Delete(context.Background(), fixture.DepositId) }, want: coreErrors.NotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newDepositServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertDepositError(t, got, err, tt.want)
		})
	}
}
