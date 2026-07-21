package card

import (
	"context"
	"crypto/rand"
	"errors"
	log "github.com/kVinsom/Bank-backend/internal/logging/service/card"
	"math/big"
	"strconv"

	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
	cardRepository "github.com/kVinsom/Bank-proto/repository/card"
	"github.com/kVinsom/Bank-proto/repository/common"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// CardService implements card use cases and card-number generation rules.
type CardService struct {
	cardRepository    cardRepository.CardRepositoryClient
	userRepository    userRepository.UserRepositoryClient
	accountRepository accountRepository.AccountRepositoryClient
	cardBIN           string
}

// NewCardService creates a card service with its required repositories and issuer BIN.
func NewCardService(
	cardRepository cardRepository.CardRepositoryClient,
	userRepository userRepository.UserRepositoryClient,
	accountRepository accountRepository.AccountRepositoryClient,
	cardBIN string) *CardService {
	return &CardService{
		cardRepository:    cardRepository,
		userRepository:    userRepository,
		accountRepository: accountRepository,
		cardBIN:           cardBIN,
	}
}

func (s *CardService) GetAll(ctx context.Context) ([]core.Card, error) {
	const operation = "get_all"
	defer log.OperationStarted(operation)()
	response, err := s.cardRepository.GetAll(ctx, &common.Empty{})
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	cards := CardsToCore(response.Cards)
	filtered := make([]core.Card, 0, len(cards))

	for _, card := range cards {
		if card.Status != core.CardStatusClosed {
			filtered = append(filtered, card)
		}
	}

	return filtered, nil
}

func (s *CardService) GetByUser(ctx context.Context, userId int64) ([]core.Card, error) {
	const operation = "get_by_user"
	defer log.OperationStarted(operation)()
	response, err := s.cardRepository.GetByUser(ctx, &common.UserIdRequest{UserId: userId})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	cards := CardsToCore(response.Cards)
	filtered := make([]core.Card, 0, len(cards))

	for _, card := range cards {
		if card.Status != core.CardStatusClosed {
			filtered = append(filtered, card)
		}
	}

	return filtered, nil
}

func (s *CardService) GetById(ctx context.Context, id int64) (*core.Card, error) {
	const operation = "get_by_id"
	defer log.OperationStarted(operation)()
	card, err := s.cardRepository.GetById(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return CardToCore(card), nil
}

func (s *CardService) GetByNumber(ctx context.Context, number string) (*core.Card, error) {
	const operation = "get_by_number"
	defer log.OperationStarted(operation)()
	card, err := s.cardRepository.GetByNumber(ctx, &cardRepository.CardNumberRequest{Number: number})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return CardToCore(card), nil
}

func (s *CardService) Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error) {
	const operation = "create"
	defer log.OperationStarted(operation)()
	if input == nil || input.UserId <= 0 || input.AccountId <= 0 {
		return nil, coreErrors.BadRequest
	}

	_, err := s.userRepository.GetById(ctx, &common.IdRequest{Id: input.UserId})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	number, err := s.GenerateCardNumber(ctx)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	_, err = s.accountRepository.GetById(ctx, &common.IdRequest{Id: input.AccountId})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}

	card := &core.Card{
		UserId:      input.UserId,
		AccountId:   input.AccountId,
		Number:      number,
		ExpiryMonth: 6,
		ExpiryYear:  3,
		Status:      core.CardStatusActive,
	}

	created, err := s.cardRepository.Create(ctx, CardToProto(card))
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return CardToCore(created), nil
}

func (s *CardService) Blocking(ctx context.Context, id int64) error {
	const operation = "block"
	defer log.OperationStarted(operation)()
	_, err := s.cardRepository.Blocking(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *CardService) Delete(ctx context.Context, id int64) error {
	const operation = "delete"
	defer log.OperationStarted(operation)()
	_, err := s.cardRepository.Delete(ctx, &common.IdRequest{Id: id})
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *CardService) GenerateCardNumber(ctx context.Context) (string, error) {
	const operation = "generate_card_number"
	defer log.OperationStarted(operation)()
	bin := s.cardBIN
	if bin == "" {
		return "", coreErrors.InternalServerError
	}

	for i := 0; i < 5; i++ {
		randomPart := make([]byte, 9)

		for i := 0; i < 9; i++ {
			n, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				return "", err
			}
			randomPart[i] = byte('0' + n.Int64())
		}

		partial := bin + string(randomPart)

		checkDigit := func(number string) int {
			sum := 0
			double := true

			for i := len(number) - 1; i >= 0; i-- {
				digit := int(number[i] - '0')

				if double {
					digit *= 2
					if digit > 9 {
						digit -= 9
					}
				}

				sum += digit
				double = !double
			}

			return (10 - (sum % 10)) % 10
		}(partial)

		cardNumber := partial + strconv.Itoa(checkDigit)

		// 3. РїРµСЂРµРІС–СЂРєР° СѓРЅС–РєР°Р»СЊРЅРѕСЃС‚С–
		_, err := s.cardRepository.GetByNumber(ctx, &cardRepository.CardNumberRequest{Number: cardNumber})
		if err != nil {
			if errors.Is(err, coreErrors.NotFound) {
				return cardNumber, nil
			}

			return "", coreErrors.InternalServerError
		}
	}

	return "", coreErrors.InternalServerError
}
