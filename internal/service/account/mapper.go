package account

import (
	log "github.com/kVinsom/Bank-backend/internal/logging/service/account"
	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// AccountToCore maps a repository account into the domain model.
func AccountToCore(account *accountRepository.Account) *core.Account {
	const operation = "account_to_core"
	defer log.MappingStarted(operation)()

	if account == nil {
		log.NilInput(operation)
		return nil
	}

	return &core.Account{
		Id:         account.Id,
		UserId:     account.UserId,
		CurrencyId: account.CurrencyId,
		Name:       account.Name,
		Balance:    account.Balance,
		Status:     core.AccountStatus(account.Status),
	}
}

// AccountToProto maps a domain account into the repository contract.
func AccountToProto(account *core.Account) *accountRepository.Account {
	const operation = "account_to_proto"
	defer log.MappingStarted(operation)()

	return &accountRepository.Account{
		Id:         account.Id,
		UserId:     account.UserId,
		CurrencyId: account.CurrencyId,
		Name:       account.Name,
		Balance:    account.Balance,
		Status:     accountRepository.AccountStatus(account.Status),
	}
}

// AccountsToCore maps repository accounts while preserving their order.
func AccountsToCore(accounts []*accountRepository.Account) []core.Account {
	const operation = "accounts_to_core"
	defer log.MappingStarted(operation)()

	result := make([]core.Account, 0, len(accounts))
	for _, account := range accounts {
		if converted := AccountToCore(account); converted != nil {
			result = append(result, *converted)
		}
	}
	return result
}
