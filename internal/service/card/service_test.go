package card

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	cardRepository "github.com/kVinsom/Bank-proto/repository/card"
	"github.com/kVinsom/Bank-proto/repository/common"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

const testCardBIN = "424242"

func TestCardService_Read_Success(t *testing.T) {
	t.Parallel()
	closed := fixture.CardProto()
	closed.Status = cardRepository.CardStatus(core.CardStatusClosed)
	tests := []struct {
		name   string
		expect func(*fixture.CardRepositoryMocks)
		invoke func(*CardService) (any, error)
		want   any
	}{
		{name: "get all", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(fixture.CardListProto(fixture.CardProto(), closed), nil)
		}, invoke: func(s *CardService) (any, error) { return s.GetAll(context.Background()) }, want: fixture.CardListCore()},
		{name: "get by user", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetByUser(gomock.Any(), &common.UserIdRequest{UserId: test.UserId}).Return(fixture.CardListProto(fixture.CardProto()), nil)
		}, invoke: func(s *CardService) (any, error) { return s.GetByUser(context.Background(), test.UserId) }, want: fixture.CardListCore()},
		{name: "get by id", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetById(gomock.Any(), &common.IdRequest{Id: test.CardId}).Return(fixture.CardProto(), nil)
		}, invoke: func(s *CardService) (any, error) { return s.GetById(context.Background(), test.CardId) }, want: cardCorePointer()},
		{name: "get by number", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(fixture.CardProto(), nil)
		}, invoke: func(s *CardService) (any, error) { return s.GetByNumber(context.Background(), test.CardNumber) }, want: cardCorePointer()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCardServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertCardResult(t, got, tt.want, err)
		})
	}
}

func TestCardService_Write_Success(t *testing.T) {
	t.Parallel()
	t.Run("create", func(t *testing.T) {
		mocks, sut := newCardServiceSUT(t)
		input := fixture.CardCreateInputCore()
		mocks.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
		mocks.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		mocks.Account.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.AccountProto(), nil)
		mocks.Card.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fixture.CardProto(), nil)
		got, err := sut.Create(context.Background(), &input)
		assertCardResult(t, got, cardCorePointer(), err)
	})
	for _, operation := range []struct {
		name   string
		expect func(*fixture.CardRepositoryMocks)
		invoke func(*CardService) error
	}{
		{name: "blocking", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().Blocking(gomock.Any(), gomock.Any()).Return(fixture.CardProto(), nil)
		}, invoke: func(s *CardService) error { return s.Blocking(context.Background(), test.CardId) }},
		{name: "delete", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&common.Empty{}, nil)
		}, invoke: func(s *CardService) error { return s.Delete(context.Background(), test.CardId) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			mocks, sut := newCardServiceSUT(t)
			operation.expect(mocks)
			if err := operation.invoke(sut); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestCardService_GenerateCardNumber(t *testing.T) {
	t.Parallel()
	mocks, sut := newCardServiceSUT(t)
	mocks.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
	number, err := sut.GenerateCardNumber(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(number) != 16 || !strings.HasPrefix(number, testCardBIN) {
		t.Fatalf("invalid generated number: %q", number)
	}
	if !validLuhn(number) {
		t.Fatalf("number does not satisfy Luhn: %q", number)
	}
}

func newCardServiceSUT(t *testing.T) (*fixture.CardRepositoryMocks, *CardService) {
	t.Helper()
	mocks := fixture.NewCardRepositoryMocks(t)
	return mocks, NewCardService(mocks.Card, mocks.User, mocks.Account, testCardBIN)
}
func cardCorePointer() any { value := fixture.CardCore(); return &value }
func assertCardResult(t *testing.T, got, want any, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("result: want %#v, got %#v", want, got)
	}
}
func validLuhn(number string) bool {
	sum := 0
	alternate := false
	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}
