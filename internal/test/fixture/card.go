package fixture

import (
	"github.com/kVinsom/Bank-backend/internal/test"
	cardRepository "github.com/kVinsom/Bank-proto/repository/card"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CardCore returns a valid card fixture.
func CardCore() core.Card {
	return core.Card{
		Id:          test.CardId,
		UserId:      test.UserId,
		AccountId:   test.AccountId,
		Number:      test.CardNumber,
		ExpiryMonth: test.CardExpiryMonth,
		ExpiryYear:  test.CardExpiryYear,
		Status:      test.CardStatus,
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
		UserId:    test.UserId,
		AccountId: test.AccountId,
	}
}
