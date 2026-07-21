package card

import (
	log "github.com/kVinsom/Bank-backend/internal/logging/service/card"
	cardRepository "github.com/kVinsom/Bank-proto/repository/card"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CardToCore maps a repository card into the domain model.
func CardToCore(card *cardRepository.Card) *core.Card {
	const operation = "card_to_core"
	defer log.MappingStarted(operation)()
	if card == nil {
		log.NilInput(operation)
		return nil
	}
	return &core.Card{
		Id:          card.Id,
		UserId:      card.UserId,
		AccountId:   card.AccountId,
		Number:      card.Number,
		ExpiryMonth: int8(card.ExpiryMonth),
		ExpiryYear:  int8(card.ExpiryYear),
		Status:      core.CardStatus(card.Status),
	}
}

// CardToProto maps a domain card into the repository contract.
func CardToProto(card *core.Card) *cardRepository.Card {
	const operation = "card_to_proto"
	defer log.MappingStarted(operation)()
	return &cardRepository.Card{
		Id:          card.Id,
		UserId:      card.UserId,
		AccountId:   card.AccountId,
		Number:      card.Number,
		ExpiryMonth: int32(card.ExpiryMonth),
		ExpiryYear:  int32(card.ExpiryYear),
		Status:      cardRepository.CardStatus(card.Status),
	}
}

// CardsToCore maps repository cards while preserving their order.
func CardsToCore(cards []*cardRepository.Card) []core.Card {
	const operation = "cards_to_core"
	defer log.MappingStarted(operation)()
	result := make([]core.Card, 0, len(cards))
	for _, card := range cards {
		if converted := CardToCore(card); converted != nil {
			result = append(result, *converted)
		}
	}
	return result
}
