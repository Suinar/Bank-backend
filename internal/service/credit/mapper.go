package credit

import (
	log "github.com/kVinsom/Bank-backend/internal/logging/service/credit"
	creditRepository "github.com/kVinsom/Bank-proto/repository/credit"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CreditToCore maps a repository credit into the domain model.
func CreditToCore(credit *creditRepository.Credit) *core.Credit {
	const operation = "credit_to_core"
	defer log.MappingStarted(operation)()
	if credit == nil {
		log.NilInput(operation)
		return nil
	}
	return &core.Credit{
		Id:             credit.Id,
		UserId:         credit.UserId,
		CurrencyId:     credit.CurrencyId,
		Amount:         credit.Amount,
		InterestRate:   credit.InterestRate,
		TermMonths:     int8(credit.TermMonths),
		MonthlyPayment: credit.MonthlyPayment,
		Status:         core.CreditStatus(credit.Status),
	}
}

// CreditToProto maps a domain credit into the repository contract.
func CreditToProto(credit *core.Credit) *creditRepository.Credit {
	const operation = "credit_to_proto"
	defer log.MappingStarted(operation)()
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

// CreditsToCore maps repository credits while preserving their order.
func CreditsToCore(credits []*creditRepository.Credit) []core.Credit {
	const operation = "credits_to_core"
	defer log.MappingStarted(operation)()
	result := make([]core.Credit, 0, len(credits))
	for _, credit := range credits {
		if converted := CreditToCore(credit); converted != nil {
			result = append(result, *converted)
		}
	}
	return result
}
