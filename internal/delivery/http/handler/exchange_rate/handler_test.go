package exchange_rate

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kVinsom/Bank-backend/internal/mocks/services"
	"github.com/kVinsom/Bank-backend/internal/test"
	"github.com/kVinsom/Bank-backend/internal/test/fixture"
)

func TestExchangeRateHandler_GetAllRanking_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewExchangeRateSUT(t)

	expected := fixture.RankingListCore()

	service.EXPECT().GetAllRanking(gomock.Any(), test.ExchangeISOFrom).Return(expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/exchange-rates/840", nil)
	ctx.AddParam("currency_iso_from", strconv.Itoa(test.ExchangeISOFrom))

	sut.GetAllRanking(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func TestExchangeRateHandler_GetRelativeRanking_Success(t *testing.T) {
	t.Parallel()

	service, sut := NewExchangeRateSUT(t)

	expected := fixture.RankingCore()

	service.EXPECT().GetRelativeRanking(gomock.Any(), test.ExchangeISOFrom, test.ExchangeISOTo).Return(&expected, nil).Times(1)

	ctx, recorder := fixture.NewHTTPContext(t, http.MethodGet, "/exchange-rates/840/978", nil)
	ctx.AddParam("currency_iso_from", strconv.Itoa(test.ExchangeISOFrom))
	ctx.AddParam("currency_iso_to", strconv.Itoa(test.ExchangeISOTo))

	sut.GetRelativeRanking(ctx)

	fixture.AssertHTTPResponse(t, recorder, http.StatusOK, expected)
}

func NewExchangeRateSUT(t *testing.T) (*mocks.MockIExchangeRateService, *ExchangeRateHandler) {
	t.Helper()
	return fixture.NewMockSUT(t, mocks.NewMockIExchangeRateService, func(service *mocks.MockIExchangeRateService) *ExchangeRateHandler {
		return NewExchangeRateHandler(service)
	})
}
