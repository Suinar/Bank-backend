package exchange_rate

import "github.com/gin-gonic/gin"

// IExchangeRateHandler defines the exchange-rate endpoints registered by the router.
type IExchangeRateHandler interface {
	GetAllRanking(*gin.Context)
	GetRelativeRanking(*gin.Context)
}
