package currency

import (
	"context"
	"errors"
	log "github.com/kVinsom/Bank-backend/internal/logging/service/currency"

	"github.com/kVinsom/Bank-proto/repository/common"
	currencyRepository "github.com/kVinsom/Bank-proto/repository/currency"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CurrencyService implements currency catalog use cases.
type CurrencyService struct {
	currencyRepository currencyRepository.CurrencyRepositoryClient
}

// NewCurrencyService creates a currency service backed by the currency repository.
func NewCurrencyService(
	repository currencyRepository.CurrencyRepositoryClient) *CurrencyService {
	return &CurrencyService{
		currencyRepository: repository,
	}
}

func (s *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {
	const operation = "get_all"
	defer log.OperationStarted(operation)()
	response, err := s.currencyRepository.GetAll(ctx, &common.Empty{})
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return CurrenciesToCore(response.Currencies), nil
}

func (s *CurrencyService) GetById(ctx context.Context, id int64) (*core.Currency, error) {
	const operation = "get_by_id"
	defer log.OperationStarted(operation)()
	currency, err := s.currencyRepository.GetById(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}
		return nil, coreErrors.InternalServerError
	}

	return CurrencyToCore(currency), nil
}

func (s *CurrencyService) GetByIso(ctx context.Context, isoCode string) (*core.Currency, error) {
	const operation = "get_by_iso"
	defer log.OperationStarted(operation)()
	currency, err := s.currencyRepository.GetByIso(ctx, &currencyRepository.IsoCodeRequest{IsoCode: isoCode})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return CurrencyToCore(currency), nil
}

func (s *CurrencyService) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	const operation = "get_by_symbol"
	defer log.OperationStarted(operation)()
	currency, err := s.currencyRepository.GetBySymbol(ctx, &currencyRepository.SymbolRequest{Symbol: string(symbol)})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return CurrencyToCore(currency), nil
}

func (s *CurrencyService) Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error) {
	const operation = "create"
	defer log.OperationStarted(operation)()
	if input == nil || input.Name == "" || input.Symbol == '0' ||
		input.MinorUnits <= 0 || input.MinorUnits > 50 ||
		input.IsoCode == "" {
		return nil, coreErrors.BadRequest
	}

	currency := &core.Currency{
		Name:       input.Name,
		Symbol:     input.Symbol,
		IsoCode:    input.IsoCode,
		MinorUnits: input.MinorUnits,
	}

	created, err := s.currencyRepository.Create(ctx, CurrencyToProto(currency))
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return CurrencyToCore(created), nil
}

func (s *CurrencyService) Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
	const operation = "update"
	defer log.OperationStarted(operation)()
	if input == nil {
		return nil, coreErrors.BadRequest
	}

	if input.Name != nil {
		if len(*input.Name) == 0 {
			return nil, coreErrors.BadRequest
		}
	}

	if input.MinorUnits != nil {
		if *input.MinorUnits <= 0 {
			return nil, coreErrors.BadRequest
		}
	}

	if input.IsoCode != nil {
		if len(*input.IsoCode) <= 0 || len(*input.IsoCode) > 3 {
			return nil, coreErrors.BadRequest
		}
	}

	var symbol *string
	if input.Symbol != nil {
		value := string(*input.Symbol)
		symbol = &value
	}

	var minorUnits *int32
	if input.MinorUnits != nil {
		value := int32(*input.MinorUnits)
		minorUnits = &value
	}

	currency, err := s.currencyRepository.Update(ctx, &currencyRepository.UpdateCurrencyRequest{
		Id: id,
		Input: &currencyRepository.CurrencyUpdateInput{
			Name:       input.Name,
			Symbol:     symbol,
			IsoCode:    input.IsoCode,
			MinorUnits: minorUnits,
		},
	})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return CurrencyToCore(currency), nil
}

func (s *CurrencyService) Delete(ctx context.Context, id int64) error {
	const operation = "delete"
	defer log.OperationStarted(operation)()
	_, err := s.currencyRepository.Delete(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}
