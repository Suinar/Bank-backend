package fixture

import (
	"github.com/kVinsom/Bank-backend/internal/test"
	depositRepository "github.com/kVinsom/Bank-proto/repository/deposit"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// DepositCore returns a valid deposit fixture.
func DepositCore() core.Deposit {
	return core.Deposit{
		Id:           test.DepositId,
		UserId:       test.UserId,
		CurrencyId:   test.CurrencyId,
		Amount:       test.DepositAmount,
		InterestRate: test.DepositInterestRate,
		TermMonths:   test.DepositTermMonths,
		Status:       test.DepositStatus,
	}
}

func DepositListCore() []core.Deposit { return []core.Deposit{DepositCore()} }

func DepositProto() *depositRepository.Deposit {
	deposit := DepositCore()
	return &depositRepository.Deposit{
		Id:           deposit.Id,
		UserId:       deposit.UserId,
		CurrencyId:   deposit.CurrencyId,
		Amount:       deposit.Amount,
		InterestRate: deposit.InterestRate,
		TermMonths:   int32(deposit.TermMonths),
		Status:       depositRepository.DepositStatus(deposit.Status),
	}
}

func DepositListProto(deposits ...*depositRepository.Deposit) *depositRepository.DepositList {
	return &depositRepository.DepositList{
		Deposits: deposits,
	}
}

// DepositCreateInputCore returns a valid deposit creation fixture.
func DepositCreateInputCore() core.DepositCreateInput {
	return core.DepositCreateInput{
		UserId:     test.UserId,
		CurrencyId: test.CurrencyId,
		Amount:     10_000,
		TermMonths: test.DepositTermMonths,
	}
}
