package fixture

import (
	cardRepository "github.com/kVinsom/Bank-proto/repository/card"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CardCore returns a valid card fixture.
func CardCore() core.Card {
	return core.Card{
		Id:          CardId,
		UserId:      UserId,
		AccountId:   AccountId,
		Number:      CardNumber,
		ExpiryMonth: CardExpiryMonth,
		ExpiryYear:  CardExpiryYear,
		Status:      CardStatus,
	}
}

func CardListCore() []core.Card { return []core.Card{CardCore()} }

func CardProto() *cardRepository.Card {
	card := CardCore()
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

func CardListProto(cards ...*cardRepository.Card) *cardRepository.CardList {
	return &cardRepository.CardList{
		Cards: cards,
	}
}

// CardCreateInputCore returns a valid card creation fixture.
func CardCreateInputCore() core.CardCreateInput {
	return core.CardCreateInput{
		UserId:    UserId,
		AccountId: AccountId,
	}
}
