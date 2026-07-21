package card

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

var cardHandlerTestError = errors.New("service failure")

func TestCardHandler_GetAll_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	service.EXPECT().GetAll(gomock.Any()).Return(nil, cardHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/cards", nil)

	sut.GetAll(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, cardHandlerTestError.Error())
}

func TestCardHandler_GetByUser_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	service.EXPECT().GetByUser(gomock.Any(), fixture.UserId).Return(nil, cardHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/users/1/cards", nil)
	ctx.AddParam("user_id", "1")

	sut.GetByUser(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, cardHandlerTestError.Error())
}

func TestCardHandler_GetById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	service.EXPECT().GetById(gomock.Any(), fixture.CardId).Return(nil, cardHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/cards/1", nil)
	ctx.AddParam("id", "1")

	sut.GetById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, cardHandlerTestError.Error())
}

func TestCardHandler_GetByNumber_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	service.EXPECT().GetByNumber(gomock.Any(), fixture.CardNumber).Return(nil, cardHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/cards/"+fixture.CardNumber, nil)
	ctx.AddParam("number", fixture.CardNumber)

	sut.GetByNumber(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, cardHandlerTestError.Error())
}

func TestCardHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	input := core.CardCreateInput{}

	service.EXPECT().Create(gomock.Any(), gomock.Eq(&input)).Return(nil, cardHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPost, "/cards", input)

	sut.Create(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, cardHandlerTestError.Error())
}

func TestCardHandler_BlockingById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	service.EXPECT().Blocking(gomock.Any(), fixture.CardId).Return(cardHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodPatch, "/cards/1/block", nil)
	ctx.AddParam("id", "1")

	sut.BlockingById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, cardHandlerTestError.Error())
}

func TestCardHandler_DeleteById_Error(t *testing.T) {
	t.Parallel()

	service, sut := NewCardSUT(t)

	service.EXPECT().Delete(gomock.Any(), fixture.CardId).Return(cardHandlerTestError).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodDelete, "/cards/1", nil)
	ctx.AddParam("id", "1")

	sut.DeleteById(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusInternalServerError, cardHandlerTestError.Error())
}

func TestCardHandler_GetById_InvalidId(t *testing.T) {
	t.Parallel()

	_, sut := NewCardSUT(t)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/cards/invalid", nil)
	ctx.AddParam("id", "invalid")

	sut.GetById(ctx)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("want %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
}
