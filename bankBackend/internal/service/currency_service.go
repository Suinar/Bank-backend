package core

import (
	"context"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
)

type CurrencyService struct {
	repository repository.ICurrencyRepository
}

func NewCurrencyService(repository repository.ICurrencyRepository) *CurrencyService {
	return &CurrencyService{repository: repository}
}

func (service *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {}

func (service *CurrencyService) GetById(ctx context.Context, id string) (*core.Currency, error) {}

func (service *CurrencyService) GetByIsoCod(ctx context.Context, isoCode string) (*core.Currency, error) {
}

func (service *CurrencyService) GetByNumberCod(ctx context.Context, numberCode string) (*core.Currency, error) {
}

func (service *CurrencyService) GetBySymbol(ctx context.Context, symbol string) (*core.Currency, error) {
}

func (service *CurrencyService) Convert(ctx context.Context, currencyIdFrom string, amount int, currencyIdTo string) (float64, error) {
}

func (service *CurrencyService) Create(ctx context.Context, input *core.CurrencyCreateInput) (*core.Currency, error) {
}

func (service *CurrencyService) Update(ctx context.Context, id string, input *core.CurrencyUpdateInput) (*core.Currency, error) {
}

func (service *CurrencyService) Delete(ctx context.Context, id string) error {}

func (service *CurrencyService) GetAllRanking(ctx context.Context, currencyIdFrom string) ([]core.ExchangeRate, error) {
}

func (service *CurrencyService) GetRelativeRanking(ctx context.Context, currencyIdFrom string, currencyIdTo string) (*core.ExchangeRate, error) {
}
