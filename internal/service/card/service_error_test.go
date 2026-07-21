package card

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	coreErrors "github.com/kVinsom/Bank-repository-service/pkg"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

var cardRepositoryError = errors.New("repository failure")

func TestCardService_Create_InvalidInput(t *testing.T) {
	t.Parallel()
	for _, input := range []*core.CardCreateInput{nil, {UserId: 0, AccountId: 1}, {UserId: -1, AccountId: 1}, {UserId: 1, AccountId: 0}, {UserId: 1, AccountId: -1}} {
		got, err := (&CardService{}).Create(context.Background(), input)
		assertCardError(t, got, err, coreErrors.BadRequest)
	}
}

func TestCardService_RepositoryErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.CardRepositoryMocks)
		invoke func(*CardService) (any, error)
		want   error
	}{
		{name: "get all", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
		}, invoke: func(s *CardService) (any, error) { return s.GetAll(context.Background()) }, want: coreErrors.InternalServerError},
		{name: "get by user not found", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetByUser(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CardService) (any, error) { return s.GetByUser(context.Background(), fixture.UserId) }, want: coreErrors.NotFound},
		{name: "get by id", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
		}, invoke: func(s *CardService) (any, error) { return s.GetById(context.Background(), fixture.CardId) }, want: coreErrors.InternalServerError},
		{name: "get by number not found", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CardService) (any, error) { return s.GetByNumber(context.Background(), fixture.CardNumber) }, want: coreErrors.NotFound},
		{name: "blocking", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().Blocking(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
		}, invoke: func(s *CardService) (any, error) { return nil, s.Blocking(context.Background(), fixture.CardId) }, want: coreErrors.InternalServerError},
		{name: "delete not found", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CardService) (any, error) { return nil, s.Delete(context.Background(), fixture.CardId) }, want: coreErrors.NotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCardServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertCardError(t, got, err, tt.want)
		})
	}
}

func TestCardService_Create_DependencyErrors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		setup func(*fixture.CardRepositoryMocks)
		want  error
	}{
		{name: "user not found", setup: func(m *fixture.CardRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, want: coreErrors.NotFound},
		{name: "user failure", setup: func(m *fixture.CardRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
		}, want: coreErrors.InternalServerError},
		{name: "number lookup failure", setup: func(m *fixture.CardRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
			m.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
		}, want: coreErrors.InternalServerError},
		{name: "account not found", setup: func(m *fixture.CardRepositoryMocks) {
			m.User.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(fixture.UserProto(), nil)
			m.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
			m.Account.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, want: coreErrors.BadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCardServiceSUT(t)
			tt.setup(mocks)
			input := fixture.CardCreateInputCore()
			got, err := sut.Create(context.Background(), &input)
			assertCardError(t, got, err, tt.want)
		})
	}
}

func TestCardService_GenerateCardNumber_Errors(t *testing.T) {
	if number, err := NewCardService(nil, nil, nil, "").GenerateCardNumber(context.Background()); number != "" || !errors.Is(err, coreErrors.InternalServerError) {
		t.Fatalf("missing BIN: got %q, %v", number, err)
	}
	mocks, sut := newCardServiceSUT(t)
	mocks.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
	number, err := sut.GenerateCardNumber(context.Background())
	if number != "" || !errors.Is(err, coreErrors.InternalServerError) {
		t.Fatalf("repository failure: got %q, %v", number, err)
	}
}

func assertCardError(t *testing.T, got any, err, want error) {
	t.Helper()
	if got != nil && !reflect.ValueOf(got).IsNil() {
		t.Fatalf("result: want nil, got %#v", got)
	}
	if !errors.Is(err, want) {
		t.Fatalf("error: want %v, got %v", want, err)
	}
}

func TestCardService_RepositoryErrorBranches(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect func(*fixture.CardRepositoryMocks)
		invoke func(*CardService) (any, error)
		want   error
	}{
		{name: "get by user failure", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetByUser(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
		}, invoke: func(s *CardService) (any, error) { return s.GetByUser(context.Background(), fixture.UserId) }, want: coreErrors.InternalServerError},
		{name: "get by id not found", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CardService) (any, error) { return s.GetById(context.Background(), fixture.CardId) }, want: coreErrors.NotFound},
		{name: "get by number failure", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
		}, invoke: func(s *CardService) (any, error) { return s.GetByNumber(context.Background(), fixture.CardNumber) }, want: coreErrors.InternalServerError},
		{name: "blocking not found", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().Blocking(gomock.Any(), gomock.Any()).Return(nil, coreErrors.NotFound)
		}, invoke: func(s *CardService) (any, error) { return nil, s.Blocking(context.Background(), fixture.CardId) }, want: coreErrors.NotFound},
		{name: "delete failure", expect: func(m *fixture.CardRepositoryMocks) {
			m.Card.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, cardRepositoryError)
		}, invoke: func(s *CardService) (any, error) { return nil, s.Delete(context.Background(), fixture.CardId) }, want: coreErrors.InternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks, sut := newCardServiceSUT(t)
			tt.expect(mocks)
			got, err := tt.invoke(sut)
			assertCardError(t, got, err, tt.want)
		})
	}
}

func TestCardService_GenerateCardNumber_Collisions(t *testing.T) {
	mocks, sut := newCardServiceSUT(t)
	mocks.Card.EXPECT().GetByNumber(gomock.Any(), gomock.Any()).Return(fixture.CardProto(), nil).Times(5)
	number, err := sut.GenerateCardNumber(context.Background())
	if number != "" || err != coreErrors.InternalServerError {
		t.Fatalf("want empty number and InternalServerError, got %q and %v", number, err)
	}
}
