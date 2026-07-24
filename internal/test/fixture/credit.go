package fixture

import (
	"github.com/kVinsom/Bank-backend/internal/test"
	creditRepository "github.com/kVinsom/Bank-proto/repository/credit"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CreditCore returns a valid credit fixture.
func CreditCore() core.Credit {
	return core.Credit{
		Id:             test.CreditId,
		UserId:         test.UserId,
		CurrencyId:     test.CurrencyId,
		Amount:         test.CreditAmount,
		InterestRate:   test.CreditInterestRate,
		TermMonths:     test.CreditTermMonths,
		MonthlyPayment: test.CreditMonthlyPayment,
		Status:         test.CreditStatus,
	}
}

func CreditListCore() []core.Credit { return []core.Credit{CreditCore()} }

func CreditProto() *creditRepository.Credit {
	credit := CreditCore()
	return &creditRepository.Credit{
		Id:             credit.Id,
		UserId:         credit.UserId,
		CurrencyId:     credit.CurrencyId,
		Amount:         credit.Amount,
		InterestRate:   credit.InterestRate,
		TermMonths:     int32(credit.TermMonths),
		MonthlyPayment: credit.MonthlyPayment,
		Status:         creditRepository.CreditStatus(credit.Status),
	}
}

func CreditListProto(credits ...*creditRepository.Credit) *creditRepository.CreditList {
	return &creditRepository.CreditList{
		Credits: credits,
	}
}

// CreditCreateInputCore returns a valid credit creation fixture.
func CreditCreateInputCore() core.CreditCreateInput {
	return core.CreditCreateInput{
		UserId:     test.UserId,
		CurrencyId: test.CurrencyId,
		Amount:     test.CreditAmount,
		TermMonths: test.CreditTermMonths,
	}
}
