package card

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kVinsom/Bank-backend/internal/mocks/services"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestCardHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	expected := fixture.CardListCore()

	service.EXPECT().GetAll(gomock.Any()).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/cards", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCardHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	expected := fixture.CardListCore()

	service.EXPECT().GetByUser(gomock.Any(), fixture.UserId).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1/cards", nil)
	ctx.AddParam("user_id", "1")

	sut.GetByUser(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCardHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	expected := fixture.CardCore()

	service.EXPECT().GetById(gomock.Any(), fixture.CardId).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/cards/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCardHandler_GetByNumber_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	expected := fixture.CardCore()

	service.EXPECT().GetByNumber(gomock.Any(), fixture.CardNumber).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/cards/"+fixture.CardNumber, nil)
	ctx.AddParam("number", fixture.CardNumber)

	sut.GetByNumber(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCardHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	input := fixture.CardCreateInputCore()
	expected := fixture.CardCore()

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/cards", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestCardHandler_BlockingById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	service.EXPECT().Blocking(gomock.Any(), fixture.CardId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/cards/1/block", nil)
	ctx.AddParam("id", "1")

	sut.BlockingById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func TestCardHandler_DeleteById_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	service.EXPECT().Delete(gomock.Any(), fixture.CardId).Return(nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/cards/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, nil)
}

func NewCardSUT(t *testing.T) (*mocks.MockICardService, *CardHandler) {
	t.Helper()
	return fixture.NewMockSUT(t, mocks.NewMockICardService, func(service *mocks.MockICardService) *CardHandler {
		return NewCardHandler(service)
	})
}
