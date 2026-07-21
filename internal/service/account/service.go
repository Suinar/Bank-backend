package account

import (
	"context"
	"errors"

	log "github.com/kVinsom/Bank-backend/internal/logging/service/account"
	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
	"github.com/kVinsom/Bank-proto/repository/common"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// AccountService implements account use cases over repository clients.
type AccountService struct {
	accountRepository  accountRepository.AccountRepositoryClient
	userRepository     userRepository.UserRepositoryClient
	currencyRepository currencyRepository.CurrencyRepositoryClient
}

// NewAccountService creates an account service with its required repositories.
func NewAccountService(
	accountRepository accountRepository.AccountRepositoryClient,
	userRepository userRepository.UserRepositoryClient,
	currencyRepository currencyRepository.CurrencyRepositoryClient) *AccountService {
	return &AccountService{
		accountRepository:  accountRepository,
		userRepository:     userRepository,
		currencyRepository: currencyRepository,
	}
}

func (s *AccountService) GetAll(ctx context.Context) ([]core.Account, error) {
	const operation = "get_all"
	defer log.OperationStarted(operation)()

	response, err := s.accountRepository.GetAll(ctx, &common.Empty{})
	if err != nil {
		log.DependencyFailed(operation, "account_repository", err)
		return nil, coreErrors.InternalServerError
	}

	accounts := AccountsToCore(response.Accounts)
	filtered := make([]core.Account, 0, len(accounts))

	for _, account := range accounts {
		if account.Status != core.AccountStatusClosed {
			filtered = append(filtered, account)
		}
	}

	return filtered, nil
}

func (s *AccountService) GetByUser(ctx context.Context, userId int64) ([]core.Account, error) {
	const operation = "get_by_user"
	defer log.OperationStarted(operation)()

	response, err := s.accountRepository.GetByUser(ctx, &common.UserIdRequest{UserId: userId})
	if err != nil {
		log.DependencyFailed(operation, "account_repository", err)
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	accounts := AccountsToCore(response.Accounts)
	filtered := make([]core.Account, 0, len(accounts))

	for _, account := range accounts {
		if account.Status != core.AccountStatusClosed {
			filtered = append(filtered, account)
		}
	}

	return filtered, nil
}

func (s *AccountService) GetById(ctx context.Context, id int64) (*core.Account, error) {
	const operation = "get_by_id"
	defer log.OperationStarted(operation)()

	account, err := s.accountRepository.GetById(ctx, &common.IdRequest{Id: id})
	if err != nil {
		log.DependencyFailed(operation, "account_repository", err)
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return AccountToCore(account), nil
}

func (s *AccountService) Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error) {
	const operation = "create"
	defer log.OperationStarted(operation)()

	if input == nil || input.UserId <= 0 || input.CurrencyId <= 0 ||
		len(input.Name) <= 0 || len(input.Name) > 15 {
		log.ValidationFailed(operation, "invalid account input")
		return nil, coreErrors.BadRequest
	}

	_, err := s.userRepository.GetById(ctx, &common.IdRequest{Id: input.UserId})
	if err != nil {
		log.DependencyFailed(operation, "user_repository", err)
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}

	_, err = s.currencyRepository.GetById(ctx, &common.IdRequest{Id: input.CurrencyId})
	if err != nil {
		log.DependencyFailed(operation, "currency_repository", err)
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}

	account := &core.Account{
		UserId:     input.UserId,
		CurrencyId: input.CurrencyId,
		Name:       input.Name,
		Balance:    0,
		Status:     core.AccountStatusActive,
	}

	created, err := s.accountRepository.Create(ctx, AccountToProto(account))
	if err != nil {
		log.DependencyFailed(operation, "account_repository", err)
		return nil, coreErrors.InternalServerError
	}

	return AccountToCore(created), nil
}

func (s *AccountService) Blocking(ctx context.Context, id int64) error {
	const operation = "block"
	defer log.OperationStarted(operation)()

	_, err := s.accountRepository.Blocking(ctx, &common.IdRequest{Id: id})
	if err != nil {
		log.DependencyFailed(operation, "account_repository", err)
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *AccountService) Close(ctx context.Context, id int64) error {
	const operation = "close"
	defer log.OperationStarted(operation)()

	_, err := s.accountRepository.Close(ctx, &common.IdRequest{Id: id})
	if err != nil {
		log.DependencyFailed(operation, "account_repository", err)
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *AccountService) Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error) {
	const operation = "update"
	defer log.OperationStarted(operation)()

	if input == nil || input.Name == nil {
		log.ValidationFailed(operation, "missing account name")
		return nil, coreErrors.BadRequest
	}

	if input.Name != nil {
		if len(*input.Name) == 0 || len(*input.Name) > 15 {
			log.ValidationFailed(operation, "invalid account name length")
			return nil, coreErrors.BadRequest
		}
	}

	updated, err := s.accountRepository.Update(ctx, &accountRepository.UpdateAccountRequest{
		Id: id,
		Input: &accountRepository.AccountUpdateInput{
			Name: input.Name,
		},
	})
	if err != nil {
		log.DependencyFailed(operation, "account_repository", err)
		if errors.Is(err, coreErrors.BadRequest) {
			return nil, coreErrors.BadRequest
		}

		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return AccountToCore(updated), nil
}

func (s *AccountService) Delete(ctx context.Context, id int64) error {
	const operation = "delete"
	defer log.OperationStarted(operation)()

	_, err := s.accountRepository.Delete(ctx, &common.IdRequest{Id: id})
	if err != nil {
		log.DependencyFailed(operation, "account_repository", err)
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}
