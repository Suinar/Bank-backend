package fixture

import (
	"testing"

	"github.com/golang/mock/gomock"
	repositoryMocks "github.com/kVinsom/Bank-backend/internal/mocks/repositories"
	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

type AccountRepositoryMocks struct {
	Account  *repositoryMocks.MockAccountRepositoryClient
	User     *repositoryMocks.MockUserRepositoryClient
	Currency *repositoryMocks.MockCurrencyRepositoryClient
}

func NewAccountRepositoryMocks(t *testing.T) *AccountRepositoryMocks {
	t.Helper()

	ctrl := gomock.NewController(t)

	return &AccountRepositoryMocks{
		Account:  repositoryMocks.NewMockAccountRepositoryClient(ctrl),
		User:     repositoryMocks.NewMockUserRepositoryClient(ctrl),
		Currency: repositoryMocks.NewMockCurrencyRepositoryClient(ctrl),
	}
}

// AccountCore returns a valid account fixture.
func AccountCore() core.Account {
	return core.Account{
		Id:         AccountId,
		UserId:     UserId,
		CurrencyId: CurrencyId,
		Name:       AccountName,
		Balance:    AccountBalance,
		Status:     AccountStatus,
	}
}

// AccountListCore returns a deterministic account collection for handler tests.
func AccountListCore() []core.Account {
	return []core.Account{AccountCore()}
}

func AccountProto() *accountRepository.Account {
	account := AccountCore()

	return &accountRepository.Account{
		Id:         account.Id,
		UserId:     account.UserId,
		CurrencyId: account.CurrencyId,
		Name:       account.Name,
		Balance:    account.Balance,
		Status:     accountRepository.AccountStatus(account.Status),
	}
}

func AccountListProto(accounts ...*accountRepository.Account) *accountRepository.AccountList {
	return &accountRepository.AccountList{
		Accounts: accounts,
	}
}

// AccountCreateInputCore returns a valid account creation fixture.
func AccountCreateInputCore() core.AccountCreateInput {
	return core.AccountCreateInput{
		UserId:     UserId,
		CurrencyId: CurrencyId,
		Name:       AccountName,
	}
}

// AccountUpdateInputCore returns a valid account update fixture.
func AccountUpdateInputCore() core.AccountUpdateInput {
	return core.AccountUpdateInput{
		Name: StringPointer(AccountName),
	}
}
