package credit

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

var creditRepositoryError = errors.New("repository failure")

func TestCreditService_Create_InvalidInput(t *testing.T) {
	t.Parallel()
	inputs := []*core.CreditCreateInput{nil, {UserId: 0, CurrencyId: 1, Amount: 1000, TermMonths: 1}, {UserId: 1, CurrencyId: 0, Amount: 1000, TermMonths: 1}, {UserId: 1, CurrencyId: 1, Amount: 999, TermMonths: 1}, {UserId: 1, CurrencyId: 1, Amount: 100001, TermMonths: 1}, {UserId: 1, CurrencyId: 1, Amount: 1000, TermMonths: 0}, {UserId: 1, CurrencyId: 1, Amount: 1000, TermMonths: 25}}
	for _, input := range inputs {
		got, err := (&CreditService{}).Create(context.Background(), input)
		assertCreditError(t, got, err, coreErrors.BadRequest)
	}
}

func TestCreditService_RepositoryErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.CreditRepositoryMocks)
		invoke func(*CreditService) (any, error)
		want   error
	}{
		{name: "get all", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, creditRepositoryError)
		}, invoke: func(s *CreditService) (any, error) { return s.GetAll(context.Background()) }, want: coreErrors.InternalServerError},
		{name: "get by user not found", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().GetByUser(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CreditService) (any, error) { return s.GetByUser(context.Background(), fixture.UserId) }, want: coreErrors.NotFound},
		{name: "get by id", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, creditRepositoryError)
		}, invoke: func(s *CreditService) (any, error) { return s.GetById(context.Background(), fixture.CreditId) }, want: coreErrors.InternalServerError},
		{name: "repay not found", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().Repay(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CreditService) (any, error) {
			return nil, s.Repay(context.Background(), fixture.CreditId, fixture.TransactionAmount)
		}, want: coreErrors.NotFound},
		{name: "delete", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, creditRepositoryError)
		}, invoke: func(s *CreditService) (any, error) { return nil, s.Delete(context.Background(), fixture.CreditId) }, want: coreErrors.InternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCreditServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertCreditError(t, got, err, tt.want)
		})
	}
}

func TestCreditService_Create_DependencyErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		setup func(*fixture.CreditRepositoryMocks)
		want  error
	}{
		{name: "user not found", setup: func(m *fixture.CreditRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, want: coreErrors.BadRequest},
		{name: "user failure", setup: func(m *fixture.CreditRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, creditRepositoryError)
		}, want: coreErrors.InternalServerError},
		{name: "currency not found", setup: func(m *fixture.CreditRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
			m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, want: coreErrors.BadRequest},
		{name: "create failure", setup: func(m *fixture.CreditRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
			m.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.CurrencyProto(), nil)
			m.Credit.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, creditRepositoryError)
		}, want: coreErrors.InternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCreditServiceSUT(t)
			tt.setup(mocks)
			input := fixture.CreditCreateInputCore()
			got, err := sut.Create(context.Background(), &input)
			assertCreditError(t, got, err, tt.want)
		})
	}
}

func assertCreditError(t *testing.T, got any, err, want error) {
	t.Helper()
	if got != nil && !reflect.ValueOf(got).IsNil() {
		t.Fatalf("result: want nil, got %#v", got)
	}
	if !errors.Is(err, want) {
		t.Fatalf("error: want %v, got %v", want, err)
	}
}

func TestCreditService_RepositoryErrorBranches(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.CreditRepositoryMocks)
		invoke func(*CreditService) (any, error)
		want   error
	}{
		{name: "get by user failure", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().GetByUser(gomock.Any(), gomock.Any()).Return(nil, creditRepositoryError)
		}, invoke: func(s *CreditService) (any, error) { return s.GetByUser(context.Background(), fixture.UserId) }, want: coreErrors.InternalServerError},
		{name: "get by id not found", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CreditService) (any, error) { return s.GetById(context.Background(), fixture.CreditId) }, want: coreErrors.NotFound},
		{name: "repay failure", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().Repay(gomock.Any(), gomock.Any()).Return(nil, creditRepositoryError)
		}, invoke: func(s *CreditService) (any, error) {
			return nil, s.Repay(context.Background(), fixture.CreditId, fixture.TransactionAmount)
		}, want: coreErrors.InternalServerError},
		{name: "delete not found", expect: func(m *fixture.CreditRepositoryMocks) {
			m.Credit.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CreditService) (any, error) { return nil, s.Delete(context.Background(), fixture.CreditId) }, want: coreErrors.NotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCreditServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertCreditError(t, got, err, tt.want)
		})
	}
}
