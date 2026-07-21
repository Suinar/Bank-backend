package fixture

import (
	creditRepository "github.com/kVinsom/Bank-proto/repository/credit"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CreditCore returns a valid credit fixture.
func CreditCore() core.Credit {
	return core.Credit{
		Id:             CreditId,
		UserId:         UserId,
		CurrencyId:     CurrencyId,
		Amount:         CreditAmount,
		InterestRate:   CreditInterestRate,
		TermMonths:     CreditTermMonths,
		MonthlyPayment: CreditMonthlyPayment,
		Status:         CreditStatus,
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
		UserId:     UserId,
		CurrencyId: CurrencyId,
		Amount:     CreditAmount,
		TermMonths: CreditTermMonths,
	}
}
