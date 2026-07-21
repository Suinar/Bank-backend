package deposit

import (
	"context"
	"errors"
	log "github.com/kVinsom/Bank-backend/internal/logging/service/deposit"

	"github.com/kVinsom/Bank-proto/repository/common"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	depositRepository "github.com/kVinsom/Bank-proto/repository/deposit"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// DepositService implements deposit lifecycle use cases.
type DepositService struct {
	depositRepository  depositRepository.DepositRepositoryClient
	userRepository     userRepository.UserRepositoryClient
	currencyRepository currencyRepository.CurrencyRepositoryClient
}

// NewDepositService creates a deposit service with its required repositories.
func NewDepositService(
	depositRepository depositRepository.DepositRepositoryClient,
	userRepository userRepository.UserRepositoryClient,
	currencyRepository currencyRepository.CurrencyRepositoryClient) *DepositService {
	return &DepositService{
		depositRepository:  depositRepository,
		userRepository:     userRepository,
		currencyRepository: currencyRepository,
	}
}

func (s *DepositService) GetAll(ctx context.Context) ([]core.Deposit, error) {
	const operation = "get_all"
	defer log.OperationStarted(operation)()
	response, err := s.depositRepository.GetAll(ctx, &common.Empty{})
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	deposits := DepositsToCore(response.Deposits)
	filtered := make([]core.Deposit, 0, len(deposits))

	for _, deposit := range deposits {
		if deposit.Status != core.DepositStatusClosed {
			filtered = append(filtered, deposit)
		}
	}

	return filtered, nil
}

func (s *DepositService) GetByUser(ctx context.Context, userId int64) ([]core.Deposit, error) {
	const operation = "get_by_user"
	defer log.OperationStarted(operation)()
	response, err := s.depositRepository.GetByUser(ctx, &common.UserIdRequest{UserId: userId})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	deposits := DepositsToCore(response.Deposits)
	filtered := make([]core.Deposit, 0, len(deposits))

	for _, deposit := range deposits {
		if deposit.Status != core.DepositStatusClosed {
			filtered = append(filtered, deposit)
		}
	}

	return filtered, nil
}

func (s *DepositService) GetById(ctx context.Context, id int64) (*core.Deposit, error) {
	const operation = "get_by_id"
	defer log.OperationStarted(operation)()
	deposit, err := s.depositRepository.GetById(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return DepositToCore(deposit), nil
}

func (s *DepositService) Create(ctx context.Context, input *core.DepositCreateInput) (*core.Deposit, error) {
	const operation = "create"
	defer log.OperationStarted(operation)()
	if input == nil || input.UserId <= 0 || input.CurrencyId <= 0 || input.Amount < 1000 ||
		input.Amount > 10000 || input.TermMonths < 1 || input.TermMonths > 24 {
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
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
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
		UserId:       input.UserId,
		CurrencyId:   input.CurrencyId,
		Amount:       input.Amount,
		InterestRate: interestRate,
		TermMonths:   input.TermMonths,
		Status:       core.DepositStatusActive,
	}

	created, err := s.depositRepository.Create(ctx, DepositToProto(deposit))
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return DepositToCore(created), nil
}

func (s *DepositService) Replenish(ctx context.Context, id int64, amount int) error {
	const operation = "replenish"
	defer log.OperationStarted(operation)()
	_, err := s.depositRepository.Replenish(ctx, &common.AmountRequest{Id: id, Amount: int64(amount)})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *DepositService) Delete(ctx context.Context, id int64) error {
	const operation = "delete"
	defer log.OperationStarted(operation)()
	_, err := s.depositRepository.Delete(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}
