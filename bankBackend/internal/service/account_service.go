package core

import (
	"context"
	"errors"
	"strconv"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type AccountService struct {
	accountRepository  repository.IAccountRepository
	userRepository     repository.IUserRepository
	currencyRepository repository.ICurrencyRepository
}

func NewAccountService(accountRepository repository.IAccountRepository,
	userRepository repository.IUserRepository,
	currencyRepository repository.ICurrencyRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository,
		userRepository:     userRepository,
		currencyRepository: currencyRepository}
}

func (s *AccountService) GetAll(ctx context.Context) ([]core.Account, error) {
	accounts, err := s.accountRepository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
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
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
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
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return account, nil
}

func (s *AccountService) Create(ctx context.Context, input *core.AccountCreateInput) (*core.Account, error) {
	if input == nil || input.UserId == "" || input.CurrencyId == "" ||
		len(input.Name) <= 0 || len(input.Name) > 15 {
		return nil, core.BadRequest
	}

	parsedUserId, err := strconv.ParseInt(input.UserId, 10, 64)
	if err != nil {
		return nil, core.BadRequest
	}

	parsedCurrencyId, err := strconv.ParseInt(input.CurrencyId, 10, 64)
	if err != nil {
		return nil, core.BadRequest
	}

	_, err = s.userRepository.GetById(ctx, parsedUserId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.BadRequest
		}

		return nil, core.InternalServerError
	}

	_, err = s.currencyRepository.GetById(ctx, parsedCurrencyId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.BadRequest
		}

		return nil, core.InternalServerError
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
		return nil, core.InternalServerError
	}

	// todo: notification

	return created, nil
}

func (s *AccountService) Blocking(ctx context.Context, id int64) error {
	err := s.accountRepository.Blocking(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// todo: notification

	return nil
}

func (s *AccountService) Close(ctx context.Context, id int64) error {
	err := s.accountRepository.Close(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// todo:notification

	return nil
}

func (s *AccountService) Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error) {
	if input == nil || input.Name == nil {
		return nil, core.BadRequest
	}

	if input.Name != nil {
		if len(*input.Name) == 0 || len(*input.Name) > 15 {
			return nil, core.BadRequest
		}
	}

	updated, err := s.accountRepository.Update(ctx, id, input)
	if err != nil {
		if errors.Is(err, core.BadRequest) {
			return nil, core.BadRequest
		}

		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	// todo: notification

	return updated, nil
}

func (s *AccountService) Delete(ctx context.Context, id int64) error {
	err := s.accountRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// todo: notification

	return nil
}
