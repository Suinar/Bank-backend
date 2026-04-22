package core

import (
	"context"
	"errors"
	"strconv"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type DepositService struct {
	depositRepository  repository.IDepositRepository
	userRepository     repository.IUserRepository
	currencyRepository repository.ICurrencyRepository
}

func NewDepositService(depositRepository repository.IDepositRepository,
	userRepository repository.IUserRepository,
	currencyRepository repository.ICurrencyRepository) *DepositService {
	return &DepositService{depositRepository: depositRepository,
		userRepository:     userRepository,
		currencyRepository: currencyRepository}
}

func (s *DepositService) GetAll(ctx context.Context) ([]core.Deposit, error) {
	deposits, err := s.depositRepository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	filtered := make([]core.Deposit, 0, len(deposits))

	for _, deposit := range deposits {
		if deposit.Status != core.DepositStatusClosed {
			deposits = append(deposits, deposit)
		}
	}

	return filtered, nil
}

func (s *DepositService) GetByUser(ctx context.Context, userId int64) ([]core.Deposit, error) {
	deposits, err := s.depositRepository.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	filtered := make([]core.Deposit, len(deposits))

	for _, deposit := range deposits {
		if deposit.Status != core.DepositStatusClosed {
			filtered = append(filtered, deposit)
		}
	}

	return filtered, nil
}

func (s *DepositService) GetById(ctx context.Context, id int64) (*core.Deposit, error) {
	deposit, err := s.depositRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return deposit, nil
}

func (s *DepositService) Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error) {
	if input == nil || input.UserId == "" || input.CurrencyId == "" || input.Amount < 1000 ||
		input.Amount > 10000 || input.TermMonths < 1 || input.TermMonths > 24 {
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

	var interestRate float32
	if input.Amount < 100000 {
		interestRate = 6.0
	} else if input.Amount < 50000 {
		interestRate = 4.5
	} else if input.Amount < 10000 {
		interestRate = 2.0
	}

	deposit := &core.Deposit{
		UserId:       parsedUserId,
		CurrencyId:   parsedCurrencyId,
		Amount:       input.Amount,
		InterestRate: interestRate,
		TermMonths:   input.TermMonths,
		Status:       core.DepositStatusActive,
	}

	created, err := s.depositRepository.Create(ctx, deposit)
	if err != nil {
		return nil, core.InternalServerError
	}

	// todo: notification

	return created, nil
}

func (s *DepositService) Replenish(ctx context.Context, id int64, amount int) error {
	// todo: replenish service

	return nil
}

func (s *DepositService) Delete(ctx context.Context, id int64) error {
	err := s.depositRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// todo: notification

	return nil
}
