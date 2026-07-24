package account

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
	"github.com/kVinsom/Bank-proto/repository/common"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

var accountRepositoryError = errors.New("repository failure")

func TestAccountService_GetAll_Error(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	mocks.Account.EXPECT().
		GetAll(gomock.Any(), gomock.Eq(&common.Empty{})).
		Return(nil, accountRepositoryError).
		Times(1)

	got, err := sut.GetAll(context.Background())
	assertAccountServiceError(t, got, err, coreErrors.InternalServerError)
}

func TestAccountService_GetByUser_Errors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		repository error
		want       error
	}{
		{name: "not found", repository: coreErrors.NotFound, want: coreErrors.NotFound},
		{name: "wrapped not found", repository: fmt.Errorf("get by user: %w", coreErrors.NotFound), want: coreErrors.NotFound},
		{name: "bad request", repository: coreErrors.BadRequest, want: coreErrors.InternalServerError},
		{name: "repository failure", repository: accountRepositoryError, want: coreErrors.InternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks, sut := NewAccountServiceSUT(t)
			mocks.Account.EXPECT().
				GetByUser(gomock.Any(), gomock.Eq(&common.UserIdRequest{UserId: test.UserId})).
				Return(nil, tt.repository).
				Times(1)

			got, err := sut.GetByUser(context.Background(), test.UserId)
			assertAccountServiceError(t, got, err, tt.want)
		})
	}
}

func TestAccountService_GetById_Errors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		repository error
		want       error
	}{
		{name: "not found", repository: coreErrors.NotFound, want: coreErrors.NotFound},
		{name: "wrapped not found", repository: fmt.Errorf("get by id: %w", coreErrors.NotFound), want: coreErrors.NotFound},
		{name: "bad request", repository: coreErrors.BadRequest, want: coreErrors.InternalServerError},
		{name: "repository failure", repository: accountRepositoryError, want: coreErrors.InternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks, sut := NewAccountServiceSUT(t)
			mocks.Account.EXPECT().
				GetById(gomock.Any(), gomock.Eq(&common.IdRequest{Id: test.AccountId})).
				Return(nil, tt.repository).
				Times(1)

			got, err := sut.GetById(context.Background(), test.AccountId)
			assertAccountServiceError(t, got, err, tt.want)
		})
	}
}

func TestAccountService_Create_InvalidInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *core.AccountCreateInput
	}{
		{name: "nil input", input: nil},
		{name: "invalid user", input: &core.AccountCreateInput{UserId: 0, CurrencyId: 1, Name: "account"}},
		{name: "negative user", input: &core.AccountCreateInput{UserId: -1, CurrencyId: 1, Name: "account"}},
		{name: "invalid currency", input: &core.AccountCreateInput{UserId: 1, CurrencyId: 0, Name: "account"}},
		{name: "negative currency", input: &core.AccountCreateInput{UserId: 1, CurrencyId: -1, Name: "account"}},
		{name: "empty name", input: &core.AccountCreateInput{UserId: 1, CurrencyId: 1}},
		{name: "long name", input: &core.AccountCreateInput{UserId: 1, CurrencyId: 1, Name: strings.Repeat("a", 16)}},
		{name: "very long name", input: &core.AccountCreateInput{UserId: 1, CurrencyId: 1, Name: strings.Repeat("a", 100)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, sut := NewAccountServiceSUT(t)
			got, err := sut.Create(context.Background(), tt.input)
			assertAccountServiceError(t, got, err, coreErrors.BadRequest)
		})
	}
}

func TestAccountService_Create_DependencyErrors(t *testing.T) {
	t.Parallel()

	for _, tt := range accountCreateDependencyErrorCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks, sut := NewAccountServiceSUT(t)
			expectAccountCreateDependencies(mocks, tt)
			input := fixture.AccountCreateInputCore()

			got, err := sut.Create(context.Background(), &input)
			assertAccountServiceError(t, got, err, tt.want)
		})
	}
}

type accountCreateDependencyErrorCase struct {
	name        string
	userErr     error
	currencyErr error
	accountErr  error
	want        error
}

func accountCreateDependencyErrorCases() []accountCreateDependencyErrorCase {
	return []accountCreateDependencyErrorCase{
		{name: "user not found", userErr: coreErrors.NotFound, want: coreErrors.BadRequest},
		{name: "user wrapped not found", userErr: fmt.Errorf("user: %w", coreErrors.NotFound), want: coreErrors.BadRequest},
		{name: "user bad request", userErr: coreErrors.BadRequest, want: coreErrors.InternalServerError},
		{name: "user repository failure", userErr: accountRepositoryError, want: coreErrors.InternalServerError},
		{name: "currency not found", currencyErr: coreErrors.NotFound, want: coreErrors.BadRequest},
		{name: "currency wrapped not found", currencyErr: fmt.Errorf("currency: %w", coreErrors.NotFound), want: coreErrors.BadRequest},
		{name: "currency bad request", currencyErr: coreErrors.BadRequest, want: coreErrors.InternalServerError},
		{name: "currency repository failure", currencyErr: accountRepositoryError, want: coreErrors.InternalServerError},
		{name: "account not found", accountErr: coreErrors.NotFound, want: coreErrors.InternalServerError},
		{name: "account bad request", accountErr: coreErrors.BadRequest, want: coreErrors.InternalServerError},
		{name: "account repository failure", accountErr: accountRepositoryError, want: coreErrors.InternalServerError},
	}
}

func expectAccountCreateDependencies(
	mocks *fixture.AccountRepositoryMocks,
	tt accountCreateDependencyErrorCase,
) {
	mocks.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), tt.userErr)
	if tt.userErr != nil {
		return
	}

	mocks.Currency.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.CurrencyProto(), tt.currencyErr)
	if tt.currencyErr != nil {
		return
	}

	mocks.Account.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, tt.accountErr)
}

func TestAccountService_StateChangeErrors(t *testing.T) {
	t.Parallel()

	operations := []struct {
		name   string
		expect func(*fixture.AccountRepositoryMocks, error)
		invoke func(*AccountService) error
	}{
		{
			name: "blocking",
			expect: func(mocks *fixture.AccountRepositoryMocks, err error) {
				mocks.Account.EXPECT().Blocking(gomock.Any(), gomock.Any()).Return(nil, err)
			},
			invoke: func(sut *AccountService) error {
				return sut.Blocking(context.Background(), test.AccountId)
			},
		},
		{
			name: "close",
			expect: func(mocks *fixture.AccountRepositoryMocks, err error) {
				mocks.Account.EXPECT().Close(gomock.Any(), gomock.Any()).Return(nil, err)
			},
			invoke: func(sut *AccountService) error {
				return sut.Close(context.Background(), test.AccountId)
			},
		},
		{
			name: "delete",
			expect: func(mocks *fixture.AccountRepositoryMocks, err error) {
				mocks.Account.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, err)
			},
			invoke: func(sut *AccountService) error {
				return sut.Delete(context.Background(), test.AccountId)
			},
		},
	}
	errorsToCheck := []struct {
		name       string
		repository error
		want       error
	}{
		{name: "not found", repository: coreErrors.NotFound, want: coreErrors.NotFound},
		{name: "wrapped not found", repository: fmt.Errorf("state change: %w", coreErrors.NotFound), want: coreErrors.NotFound},
		{name: "bad request", repository: coreErrors.BadRequest, want: coreErrors.InternalServerError},
		{name: "repository failure", repository: accountRepositoryError, want: coreErrors.InternalServerError},
	}

	for _, operation := range operations {
		operation := operation
		t.Run(operation.name, func(t *testing.T) {
			for _, errorCase := range errorsToCheck {
				errorCase := errorCase
				t.Run(errorCase.name, func(t *testing.T) {
					t.Parallel()

					mocks, sut := NewAccountServiceSUT(t)
					operation.expect(mocks, errorCase.repository)

					err := operation.invoke(sut)
					assertAccountServiceError(t, nil, err, errorCase.want)
				})
			}
		})
	}
}

func TestAccountService_Update_InvalidInput(t *testing.T) {
	t.Parallel()

	emptyName := ""
	longName := strings.Repeat("a", 16)
	tests := []struct {
		name  string
		input *core.AccountUpdateInput
	}{
		{name: "nil input", input: nil},
		{name: "nil name", input: &core.AccountUpdateInput{}},
		{name: "empty name", input: &core.AccountUpdateInput{Name: &emptyName}},
		{name: "long name", input: &core.AccountUpdateInput{Name: &longName}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, sut := NewAccountServiceSUT(t)
			got, err := sut.Update(context.Background(), test.AccountId, tt.input)
			assertAccountServiceError(t, got, err, coreErrors.BadRequest)
		})
	}
}

func TestAccountService_Update_RepositoryErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		repository error
		want       error
	}{
		{name: "bad request", repository: coreErrors.BadRequest, want: coreErrors.BadRequest},
		{name: "wrapped bad request", repository: fmt.Errorf("update: %w", coreErrors.BadRequest), want: coreErrors.BadRequest},
		{name: "not found", repository: coreErrors.NotFound, want: coreErrors.NotFound},
		{name: "wrapped not found", repository: fmt.Errorf("update: %w", coreErrors.NotFound), want: coreErrors.NotFound},
		{name: "repository failure", repository: accountRepositoryError, want: coreErrors.InternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks, sut := NewAccountServiceSUT(t)
			mocks.Account.EXPECT().
				Update(gomock.Any(), gomock.Any()).
				Return((*accountRepository.Account)(nil), tt.repository).
				Times(1)
			input := fixture.AccountUpdateInputCore()

			got, err := sut.Update(context.Background(), test.AccountId, &input)
			assertAccountServiceError(t, got, err, tt.want)
		})
	}
}

func assertAccountServiceError(t *testing.T, got any, err, want error) {
	t.Helper()

	if got != nil && !reflect.ValueOf(got).IsNil() {
		t.Fatalf("result: want nil, got %#v", got)
	}
	if !errors.Is(err, want) {
		t.Fatalf("error: want %v, got %v", want, err)
	}
}
