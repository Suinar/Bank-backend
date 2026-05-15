package core

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"os"
	"strconv"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	repository "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"
	notificationService "github.com/Suinar/Bank-backend/bankBackend/proto/notification"
	"github.com/joho/godotenv"
)

type CardService struct {
	cardRepository      repository.ICardRepository
	userRepository      repository.IUserRepository
	accountRepository   repository.IAccountRepository
	notificationService notificationService.NotificationServiceClient
}

func NewCardService(
	cardRepository repository.ICardRepository,
	userRepository repository.IUserRepository,
	accountRepository repository.IAccountRepository,
	notificationService notificationService.NotificationServiceClient) *CardService {
	return &CardService{
		cardRepository:      cardRepository,
		userRepository:      userRepository,
		accountRepository:   accountRepository,
		notificationService: notificationService,
	}
}

func (s *CardService) GetAll(ctx context.Context) ([]core.Card, error) {
	cards, err := s.cardRepository.GetAll(ctx)
	if err != nil {
		return nil, core.InternalServerError
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
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
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
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return card, nil
}

func (s *CardService) GetByNumber(ctx context.Context, number string) (*core.Card, error) {
	card, err := s.cardRepository.GetByNumber(ctx, number)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return card, nil
}

func (s *CardService) Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error) {
	if input == nil || input.UserId == "" || input.AccountId == "" {
		return nil, core.BadRequest
	}

	parsedUserId, err := strconv.ParseInt(input.UserId, 10, 64)
	if err != nil {
		return nil, core.BadRequest
	}

	parsedAccountId, err := strconv.ParseInt(input.AccountId, 10, 64)
	if err != nil {
		return nil, core.BadRequest
	}

	_, err = s.userRepository.GetById(ctx, parsedUserId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	number, err := s.GenerateCardNumber(ctx)
	if err != nil {
		return nil, core.InternalServerError
	}

	_, err = s.accountRepository.GetById(ctx, parsedAccountId)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return nil, core.BadRequest
		}

		return nil, core.InternalServerError
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
		return nil, core.InternalServerError
	}

	s.notificationService.SendEvent(ctx, &notificationService.NotificationEventRequest{
		Entity:   notificationService.EntityType_CARD,
		Action:   notificationService.ActionType_CREATE,
		EntityId: card.Id,
		UserId:   card.UserId,
	})

	return created, nil
}

func (s *CardService) Blocking(ctx context.Context, id int64) error {
	card, err := s.cardRepository.Blocking(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	s.notificationService.SendEvent(ctx, &notificationService.NotificationEventRequest{
		Entity:   notificationService.EntityType_CARD,
		Action:   notificationService.ActionType_BLOCK,
		EntityId: card.Id,
		UserId:   card.UserId,
	})

	return nil
}

func (s *CardService) Delete(ctx context.Context, id int64) error {
	userId, err := s.cardRepository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, core.NotFound) {
			return core.NotFound
		}

		return core.InternalServerError
	}

	s.notificationService.SendEvent(ctx, &notificationService.NotificationEventRequest{
		Entity:   notificationService.EntityType_CARD,
		Action:   notificationService.ActionType_DELETE,
		EntityId: id,
		UserId:   userId,
	})

	return nil
}

func (s *CardService) GenerateCardNumber(ctx context.Context) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", core.InternalServerError
	}

	bin := os.Getenv("CARD_BIN")
	if bin == "" {
		return "", core.InternalServerError
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

		// 3. перевірка унікальності
		_, err := s.cardRepository.GetByNumber(ctx, cardNumber)
		if err != nil {
			if errors.Is(err, core.NotFound) {
				return cardNumber, nil
			}

			return "", core.InternalServerError
		}
	}

	return "", core.InternalServerError
}
