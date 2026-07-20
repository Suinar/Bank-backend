package card

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	accountRepository "github.com/kVinsom/Bank-proto/repository/account"
	cardRepository "github.com/kVinsom/Bank-proto/repository/card"
	userRepository "github.com/kVinsom/Bank-proto/repository/user"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

type CardService struct {
	cardRepository    cardRepository.CardRepositoryClient
	userRepository    userRepository.UserRepositoryClient
	accountRepository accountRepository.AccountRepositoryClient
}

func NewCardService(
	cardRepository cardRepository.CardRepositoryClient,
	userRepository userRepository.UserRepositoryClient,
	accountRepository accountRepository.AccountRepositoryClient) *CardService {
	return &CardService{
		cardRepository:    cardRepository,
		userRepository:    userRepository,
		accountRepository: accountRepository,
	}
}

func (s *CardService) GetAll(ctx context.Context) ([]core.Card, error) {
	cards, err := s.cardRepository.GetAll(ctx)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	filtered := make([]core.Card, len(cards))

	for _, card := range cards {
		if card.Status != core.CardStatusClosed {
			filtered = append(filtered, card)
		}
	}

	return filtered, nil
}

func (s *CardService) GetByUser(ctx context.Context, userId int64) ([]core.Card, error) {
	cards, err := s.cardRepository.GetByUser(ctx, userId)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	filtered := make([]core.Card, len(cards))

	for _, card := range cards {
		if card.Status != core.CardStatusClosed {
			filtered = append(filtered, card)
		}
	}

	return filtered, nil
}

func (s *CardService) GetById(ctx context.Context, id int64) (*core.Card, error) {
	card, err := s.cardRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return card, nil
}

func (s *CardService) GetByNumber(ctx context.Context, number string) (*core.Card, error) {
	card, err := s.cardRepository.GetByNumber(ctx, number)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.NotFound
		}

		return nil, coreErrors.InternalServerError
	}

	return card, nil
}

func (s *CardService) Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error) {
	if input == nil || input.UserId == "" || input.AccountId == "" {
		return nil, coreErrors.BadRequest
	}

	parsedUserId, err := strconv.ParseInt(input.UserId, 10, 64)
	if err != nil {
		return nil, coreErrors.BadRequest
	}

	parsedAccountId, err := strconv.ParseInt(input.AccountId, 10, 64)
	if err != nil {
		return nil, coreErrors.BadRequest
	}

	_, err = s.userRepository.GetById(ctx, parsedUserId)
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

	_, err = s.accountRepository.GetById(ctx, parsedAccountId)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return nil, coreErrors.BadRequest
		}

		return nil, coreErrors.InternalServerError
	}

	card := &core.Card{
		UserId:      parsedUserId,
		AccountId:   parsedAccountId,
		Number:      number,
		ExpiryMonth: 6,
		ExpiryYear:  3,
		Status:      core.CardStatusActive,
	}

	created, err := s.cardRepository.Create(ctx, card)
	if err != nil {
		return nil, coreErrors.InternalServerError
	}

	return created, nil
}

func (s *CardService) Blocking(ctx context.Context, id int64) error {
	card, err := s.cardRepository.Blocking(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *CardService) Delete(ctx context.Context, id int64) error {
	userId, err := s.cardRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, coreErrors.NotFound) {
			return coreErrors.NotFound
		}

		return coreErrors.InternalServerError
	}

	return nil
}

func (s *CardService) GenerateCardNumber(ctx context.Context) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", coreErrors.InternalServerError
	}

	bin := os.Getenv("CARD_BIN")
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
		_, err := s.cardRepository.GetByNumber(ctx, cardNumber)
		if err != nil {
			if errors.Is(err, coreErrors.NotFound) {
				return cardNumber, nil
			}

			return "", coreErrors.InternalServerError
		}
	}

	return "", coreErrors.InternalServerError
}
