package account

import (
	"context"
	"errors"
	"strconv"

	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type AccountService struct {
	accountRepository  accountRepository.AccountRepositoryClient
	userRepository     userRepository.UserRepositoryClient
	currencyRepository currencyRepository.CurrencyRepositoryClient
}

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
	accounts, err := s.accountRepository.GetAll(ctx)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	filtered := make([]core.Account, 0, len(accounts))

	for _, account := range accounts {
		if account.Status != core.AccountStatusClosed {
			filtered = append(filtered, account)
		}
	}

	return filtered, nil
}

func (s *AccountService) GetByUser(ctx context.Context, userId int64) ([]core.Account, error) {
	accounts, err := s.accountRepository.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	filtered := make([]core.Account, 0, len(accounts))

	for _, account := range accounts {
		if account.Status != core.AccountStatusClosed {
			filtered = append(filtered, account)
		}
	}

	return filtered, nil
}

func (s *AccountService) GetById(ctx context.Context, id int64) (*core.Account, error) {
	account, err := s.accountRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return account, nil
}

func (s *AccountService) Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error) {
	if input == nil || input.UserId == "" || input.CurrencyId == "" ||
		len(input.Name) <= 0 || len(input.Name) > 15 {
		return nil, coreErrors.BadRequest
	}

	parsedUserId, err := strconv.ParseInt(input.UserId, 10, 64)
	if err != nil {
		return nil, coreErrors.BadRequest
	}

	parsedCurrencyId, err := strconv.ParseInt(input.CurrencyId, 10, 64)
	if err != nil {
		return nil, coreErrors.BadRequest
	}

	_, err = s.userRepository.GetById(ctx, parsedUserId)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}

	_, err = s.currencyRepository.GetById(ctx, parsedCurrencyId)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}

	account := &core.Account{
		UserId:     parsedUserId,
		CurrencyId: parsedCurrencyId,
		Name:       input.Name,
		Balance:    0,
		Status:     core.AccountStatusActive,
	}

	created, err := s.accountRepository.Create(ctx, account)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return created, nil
}

func (s *AccountService) Blocking(ctx context.Context, id int64) error {
	account, err := s.accountRepository.Blocking(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *AccountService) Close(ctx context.Context, id int64) error {
	account, err := s.accountRepository.Close(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *AccountService) Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error) {
	if input == nil || input.Name == nil {
		return nil, coreErrors.BadRequest
	}

	if input.Name != nil {
		if len(*input.Name) == 0 || len(*input.Name) > 15 {
			return nil, coreErrors.BadRequest
		}
	}

	updated, err := s.accountRepository.Update(ctx, id, input)
	if err != nil {
		if errors.Is(err, coreErrors.BadRequest) {
			return nil, coreErrors.BadRequest
		}

		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return updated, nil
}

func (s *AccountService) Delete(ctx context.Context, id int64) error {
	userId, err := s.accountRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}
