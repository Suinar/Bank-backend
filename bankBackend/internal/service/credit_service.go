package core

import (
	"context"
	"errors"
	"strconv"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type CreditService struct {
	creditRepository   repository.ICreditRepository
	userRepository     repository.ICreditRepository
	currencyRepository repository.ICreditRepository
}

func NewCreditService(creditRepository repository.ICreditRepository,
	userRepository repository.ICreditRepository,
	currencyRepository repository.ICreditRepository) *CreditService {
	return &CreditService{
		creditRepository:   creditRepository,
		userRepository:     userRepository,
		currencyRepository: currencyRepository,
	}
}

func (s *CreditService) GetAll(ctx context.Context) ([]core.Credit, error) {
	credits, err := s.creditRepository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	filtered := make([]core.Credit, len(credits))

	for _, credit := range credits {
		if credit.Status != core.CreditStatusClosed {
			filtered = append(filtered, credit)
		}
	}

	return filtered, nil
}

func (s *CreditService) GetByUser(ctx context.Context, userId uint64) ([]core.Credit, error) {
	credits, err := s.creditRepository.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	filtered := make([]core.Credit, len(credits))

	for _, credit := range credits {
		if credit.Status != core.CreditStatusClosed {
			filtered = append(filtered, credit)
		}
	}

	return filtered, nil
}

func (s *CreditService) GetById(ctx context.Context, id uint64) (*core.Credit, error) {
	credit, err := s.creditRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return credit, nil
}

func (s *CreditService) Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error) {
	if input == nil || input.UserId == "" || input.CurrencyId == "" || input.Amount < 1000 ||
		input.Amount > 100000 || input.TermMonths < 1 || input.TermMonths > 24 {
		return nil, core.BadRequest
	}

	parsedUserId, err := strconv.ParseUint(input.UserId, 10, 64)
	if err != nil {
		return nil, core.BadRequest
	}

	parsedCurrencyId, err := strconv.ParseUint(input.CurrencyId, 10, 64)
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

	monthlyPayment := input.Amount/uint64(input.TermMonths) + input.Amount/10

	var interestRate float32
	if input.Amount < 100000 {
		interestRate = 5.0
	} else if input.Amount < 50000 {
		interestRate = 7.5
	} else if input.Amount < 10000 {
		interestRate = 10.0
	}

	credit := &core.Credit{
		UserId:         parsedUserId,
		CurrencyId:     parsedCurrencyId,
		Amount:         input.Amount,
		InterestRate:   interestRate,
		TermMonths:     input.TermMonths,
		MonthlyPayment: monthlyPayment,
		Status:         core.CreditStatusActive,
	}

	created, err := s.creditRepository.Create(ctx, credit)
	if err != nil {
		return nil, core.InternalServerError
	}

	// todo: notification

	return created, nil
}

func (s *CreditService) Repay(ctx context.Context, id uint64, amount int) error {
	// todo: repay service

	return nil
}

func (s *CreditService) Delete(ctx context.Context, id uint64) error {
	err := s.creditRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	// todo: notification

	return nil
}
