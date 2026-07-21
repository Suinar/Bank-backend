package credit

import (
	"context"
	"errors"
	log "github.com/kVinsom/Bank-backend/internal/logging/service/credit"

	"github.com/kVinsom/Bank-proto/repository/common"
	creditRepository "github.com/kVinsom/Bank-proto/repository/credit"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CreditService implements credit lifecycle use cases.
type CreditService struct {
	creditRepository   creditRepository.CreditRepositoryClient
	userRepository     userRepository.UserRepositoryClient
	currencyRepository currencyRepository.CurrencyRepositoryClient
}

// NewCreditService creates a credit service with its required repositories.
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
	const operation = "get_all"
	defer log.OperationStarted(operation)()
	response, err := s.creditRepository.GetAll(ctx, &common.Empty{})
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	credits := CreditsToCore(response.Credits)
	filtered := make([]core.Credit, 0, len(credits))

	for _, credit := range credits {
		if credit.Status != core.CreditStatusClosed {
			filtered = append(filtered, credit)
		}
	}

	return filtered, nil
}

func (s *CreditService) GetByUser(ctx context.Context, userId int64) ([]core.Credit, error) {
	const operation = "get_by_user"
	defer log.OperationStarted(operation)()
	response, err := s.creditRepository.GetByUser(ctx, &common.UserIdRequest{UserId: userId})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	credits := CreditsToCore(response.Credits)
	filtered := make([]core.Credit, 0, len(credits))

	for _, credit := range credits {
		if credit.Status != core.CreditStatusClosed {
			filtered = append(filtered, credit)
		}
	}

	return filtered, nil
}

func (s *CreditService) GetById(ctx context.Context, id int64) (*core.Credit, error) {
	const operation = "get_by_id"
	defer log.OperationStarted(operation)()
	credit, err := s.creditRepository.GetById(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return CreditToCore(credit), nil
}

func (s *CreditService) Create(ctx context.Context, input *core.CreditCreateInput) (*core.Credit, error) {
	const operation = "create"
	defer log.OperationStarted(operation)()
	if input == nil || input.UserId <= 0 || input.CurrencyId <= 0 || input.Amount < 1000 ||
		input.Amount > 100000 || input.TermMonths < 1 || input.TermMonths > 24 {
		return nil, coreErrors.BadRequest
	}

	_, err := s.userRepository.GetById(ctx, &common.IdRequest{Id: input.UserId})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}

	_, err = s.currencyRepository.GetById(ctx, &common.IdRequest{Id: input.CurrencyId})
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
		UserId:         input.UserId,
		CurrencyId:     input.CurrencyId,
		Amount:         input.Amount,
		InterestRate:   interestRate,
		TermMonths:     input.TermMonths,
		MonthlyPayment: monthlyPayment,
		Status:         core.CreditStatusActive,
	}

	created, err := s.creditRepository.Create(ctx, CreditToProto(credit))
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return CreditToCore(created), nil
}

func (s *CreditService) Repay(ctx context.Context, id int64, amount int) error {
	const operation = "repay"
	defer log.OperationStarted(operation)()
	_, err := s.creditRepository.Repay(ctx, &common.AmountRequest{Id: id, Amount: int64(amount)})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *CreditService) Delete(ctx context.Context, id int64) error {
	const operation = "delete"
	defer log.OperationStarted(operation)()
	_, err := s.creditRepository.Delete(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}
