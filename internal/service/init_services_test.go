package service

import (
	"testing"

	"github.com/kVinsom/Bank-backend/internal/configs"
	grpcClient "github.com/kVinsom/Bank-backend/internal/delivery/grps/client"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestInitServices(t *testing.T) {
	mocks := fixture.NewAccountRepositoryMocks(t)
	cardMocks := fixture.NewCardRepositoryMocks(t)
	creditMocks := fixture.NewCreditRepositoryMocks(t)
	depositMocks := fixture.NewDepositRepositoryMocks(t)
	exchangeMocks := fixture.NewExchangeRateRepositoryMocks(t)
	cfg := &configs.Config{}
	cfg.CardConfig.BIN = "424242"
	repositories := &grpcClient.RepositoryClients{
		Account:  mocks.Account,
		Card:     cardMocks.Card,
		Credit:   creditMocks.Credit,
		Currency: mocks.Currency,
		Deposit:  depositMocks.Deposit,
		User:     mocks.User,
	}

	services := InitServices(cfg, repositories, exchangeMocks.Ranking)

	if services.Account == nil || services.Card == nil || services.Credit == nil ||
		services.Currency == nil || services.Deposit == nil || services.ExchangeRate == nil || services.User == nil {
		t.Fatalf("InitServices returned incomplete container: %#v", services)
	}
}
