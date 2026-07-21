package account

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
	"github.com/kVinsom/Bank-proto/repository/common"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

func TestAccountService_GetAll_Success(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	active := fixture.AccountProto()
	closed := fixture.AccountProto()
	closed.Id++
	closed.Status = accountRepository.AccountStatus(core.AccountStatusClosed)

	mocks.Account.EXPECT().
		GetAll(gomock.Any(), gomock.Eq(&common.Empty{})).
		Return(fixture.AccountListProto(active, closed), nil).
		Times(1)

	got, err := sut.GetAll(context.Background())
	assertAccountServiceResult(t, got, fixture.AccountListCore(), err)
}

func TestAccountService_GetByUser_Success(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	mocks.Account.EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(&common.UserIdRequest{UserId: fixture.UserId})).
		Return(fixture.AccountListProto(fixture.AccountProto()), nil).
		Times(1)

	got, err := sut.GetByUser(context.Background(), fixture.UserId)
	assertAccountServiceResult(t, got, fixture.AccountListCore(), err)
}

func TestAccountService_GetById_Success(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	mocks.Account.EXPECT().
		GetById(gomock.Any(), gomock.Eq(&common.IdRequest{Id: fixture.AccountId})).
		Return(fixture.AccountProto(), nil).
		Times(1)

	got, err := sut.GetById(context.Background(), fixture.AccountId)
	want := fixture.AccountCore()
	assertAccountServiceResult(t, got, &want, err)
}

func TestAccountService_Create_Success(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	input := fixture.AccountCreateInputCore()

	mocks.User.EXPECT().
		GetById(gomock.Any(), gomock.Eq(&common.IdRequest{Id: fixture.UserId})).
		Return(fixture.UserProto(), nil).
		Times(1)
	mocks.Currency.EXPECT().
		GetById(gomock.Any(), gomock.Eq(&common.IdRequest{Id: fixture.CurrencyId})).
		Return(fixture.CurrencyProto(), nil).
		Times(1)
	mocks.Account.EXPECT().
		Create(gomock.Any(), gomock.Eq(&accountRepository.Account{
			UserId:     fixture.UserId,
			CurrencyId: fixture.CurrencyId,
			Name:       fixture.AccountName,
			Status:     accountRepository.AccountStatus(core.AccountStatusActive),
		})).
		Return(fixture.AccountProto(), nil).
		Times(1)

	got, err := sut.Create(context.Background(), &input)
	want := fixture.AccountCore()
	assertAccountServiceResult(t, got, &want, err)
}

func TestAccountService_Blocking_Success(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	mocks.Account.EXPECT().
		Blocking(gomock.Any(), gomock.Eq(&common.IdRequest{Id: fixture.AccountId})).
		Return(fixture.AccountProto(), nil).
		Times(1)

	if err := sut.Blocking(context.Background(), fixture.AccountId); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAccountService_Close_Success(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	mocks.Account.EXPECT().
		Close(gomock.Any(), gomock.Eq(&common.IdRequest{Id: fixture.AccountId})).
		Return(fixture.AccountProto(), nil).
		Times(1)

	if err := sut.Close(context.Background(), fixture.AccountId); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAccountService_Update_Success(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	input := fixture.AccountUpdateInputCore()
	mocks.Account.EXPECT().
		Update(gomock.Any(), gomock.Eq(&accountRepository.UpdateAccountRequest{
			Id: fixture.AccountId,
			Input: &accountRepository.AccountUpdateInput{
				Name: input.Name,
			},
		})).
		Return(fixture.AccountProto(), nil).
		Times(1)

	got, err := sut.Update(context.Background(), fixture.AccountId, &input)
	want := fixture.AccountCore()
	assertAccountServiceResult(t, got, &want, err)
}

func TestAccountService_Delete_Success(t *testing.T) {
	t.Parallel()

	mocks, sut := NewAccountServiceSUT(t)
	mocks.Account.EXPECT().
		Delete(gomock.Any(), gomock.Eq(&common.IdRequest{Id: fixture.AccountId})).
		Return(&common.Empty{}, nil).
		Times(1)

	if err := sut.Delete(context.Background(), fixture.AccountId); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func NewAccountServiceSUT(t *testing.T) (*fixture.AccountRepositoryMocks, *AccountService) {
	t.Helper()

	mocks := fixture.NewAccountRepositoryMocks(t)

	return mocks, NewAccountService(mocks.Account, mocks.User, mocks.Currency)
}

func assertAccountServiceResult(t *testing.T, got, want any, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("result: want %#v, got %#v", want, got)
	}
}
