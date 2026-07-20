package credit

import (
	"context"
	"errors"
	"strconv"

	creditRepository "github.com/kVinsom/Bank-proto/repository/credit"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type CreditService struct {
	creditRepository   creditRepository.CreditRepositoryClient
	userRepository     userRepository.UserRepositoryClient
	currencyRepository currencyRepository.CurrencyRepositoryClient
}

func NewCreditService(
	creditRepository creditRepository.CreditRepositoryClient,
	userRepository userRepository.UserRepositoryClient,
	currencyRepository currencyRepository.CurrencyRepositoryClient) *CreditService {
	return &CreditService{
		creditRepository:   creditRepository,
		userRepository:     userRepository,
		currencyRepository: currencyRepository,
	}
}

func (s *CreditService) GetAll(ctx context.Context) ([]core.Credit, error) {
	credits, err := s.creditRepository.GetAll(ctx)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	filtered := make([]core.Credit, len(credits))

	for _, credit := range credits {
		if credit.Status != core.CreditStatusClosed {
			filtered = append(filtered, credit)
		}
	}

	return filtered, nil
}

func (s *CreditService) GetByUser(ctx context.Context, userId int64) ([]core.Credit, error) {
	credits, err := s.creditRepository.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	filtered := make([]core.Credit, len(credits))

	for _, credit := range credits {
		if credit.Status != core.CreditStatusClosed {
			filtered = append(filtered, credit)
		}
	}

	return filtered, nil
}

func (s *CreditService) GetById(ctx context.Context, id int64) (*core.Credit, error) {
	credit, err := s.creditRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return credit, nil
}

func (s *CreditService) Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error) {
	if input == nil || input.UserId == "" || input.CurrencyId == "" || input.Amount < 1000 ||
		input.Amount > 100000 || input.TermMonths < 1 || input.TermMonths > 24 {
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

	monthlyPayment := input.Amount/int64(input.TermMonths) + input.Amount/10

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
		return nil, coreErrors.InternalServerError
	}

	return created, nil
}

func (s *CreditService) Repay(ctx context.Context, id int64, amount int) error {
	credit, err := s.creditRepository.Repay(ctx, id, amount)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *CreditService) Delete(ctx context.Context, id int64) error {
	userId, err := s.creditRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}
