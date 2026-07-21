package deposit

import (
	log "github.com/kVinsom/Bank-backend/internal/logging/service/deposit"
	depositRepository "github.com/kVinsom/Bank-proto/repository/deposit"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// DepositToCore maps a repository deposit into the domain model.
func DepositToCore(deposit *depositRepository.Deposit) *core.Deposit {
	const operation = "deposit_to_core"
	defer log.MappingStarted(operation)()
	if deposit == nil {
		log.NilInput(operation)
		return nil
	}
	return &core.Deposit{
		Id:           deposit.Id,
		UserId:       deposit.UserId,
		CurrencyId:   deposit.CurrencyId,
		Amount:       deposit.Amount,
		InterestRate: deposit.InterestRate,
		TermMonths:   int8(deposit.TermMonths),
		Status:       core.DepositStatus(deposit.Status),
	}
}

// DepositToProto maps a domain deposit into the repository contract.
func DepositToProto(deposit *core.Deposit) *depositRepository.Deposit {
	const operation = "deposit_to_proto"
	defer log.MappingStarted(operation)()
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

// DepositsToCore maps repository deposits while preserving their order.
func DepositsToCore(deposits []*depositRepository.Deposit) []core.Deposit {
	const operation = "deposits_to_core"
	defer log.MappingStarted(operation)()
	result := make([]core.Deposit, 0, len(deposits))
	for _, deposit := range deposits {
		if converted := DepositToCore(deposit); converted != nil {
			result = append(result, *converted)
		}
	}
	return result
}
