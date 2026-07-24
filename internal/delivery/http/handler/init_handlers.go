package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	accountHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/account"
	cardHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/card"
	creditHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/credit"
	currencyHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/currency"
	depositHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/deposit"
	exchangeRateHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/exchange_rate"
	userHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler/user"
	serviceContainer "github.com/kVinsom/Bank-backend/internal/service"
)

// Handlers groups HTTP handlers by application resource.
type Handlers struct {
	Account      *accountHandler.AccountHandler
	Card         *cardHandler.CardHandler
	Credit       *creditHandler.CreditHandler
	Currency     *currencyHandler.CurrencyHandler
	Deposit      *depositHandler.DepositHandler
	ExchangeRate *exchangeRateHandler.ExchangeRateHandler
	User         *userHandler.UserHandler
}

// InitHandlers builds the HTTP delivery layer and registers all routes.
func InitHandlers(services *serviceContainer.Services, middleware ...gin.HandlerFunc) *gin.Engine {
	handlers := &Handlers{
		Account:      accountHandler.NewAccountHandler(services.Account),
		Card:         cardHandler.NewCardHandler(services.Card),
		Credit:       creditHandler.NewCreditHandler(services.Credit),
		Currency:     currencyHandler.NewCurrencyHandler(services.Currency),
		Deposit:      depositHandler.NewDepositHandler(services.Deposit),
		ExchangeRate: exchangeRateHandler.NewExchangeRateHandler(services.ExchangeRate),
		User:         userHandler.NewUserHandler(services.User),
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(middleware...)
	_ = router.SetTrustedProxies(nil)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	registerUserRoutes(api.Group("/users"), handlers)
	registerAccountRoutes(api.Group("/accounts"), handlers)
	registerCardRoutes(api.Group("/cards"), handlers)
	registerCreditRoutes(api.Group("/credits"), handlers)
	registerDepositRoutes(api.Group("/deposits"), handlers)
	registerCurrencyRoutes(api.Group("/currencies"), handlers)
	registerExchangeRateRoutes(api.Group("/exchange-rates"), handlers)

	return router
}

func registerUserRoutes(group *gin.RouterGroup, handlers *Handlers) {
	group.GET("", handlers.User.GetAll)
	group.GET("/:id", handlers.User.GetById)
	group.GET("/email/:email", handlers.User.GetByEmail)
	group.GET("/phone/:phone_number", handlers.User.GetByPhoneNumber)
	group.POST("", handlers.User.Create)
	group.PUT("/:id", handlers.User.UpdateById)
	group.PATCH("/:id/password", handlers.User.ChangePasswordById)
	group.DELETE("/:id", handlers.User.DeleteById)
}

func registerAccountRoutes(group *gin.RouterGroup, handlers *Handlers) {
	group.GET("", handlers.Account.GetAll)
	group.GET("/:id", handlers.Account.GetById)
	group.GET("/user/:user_id", handlers.Account.GetByUser)
	group.POST("", handlers.Account.Create)
	group.PUT("/:id", handlers.Account.UpdateById)
	group.PATCH("/:id/block", handlers.Account.BlockingById)
	group.PATCH("/:id/close", handlers.Account.CloseById)
	group.DELETE("/:id", handlers.Account.DeleteById)
}

func registerCardRoutes(group *gin.RouterGroup, handlers *Handlers) {
	group.GET("", handlers.Card.GetAll)
	group.GET("/:id", handlers.Card.GetById)
	group.GET("/user/:user_id", handlers.Card.GetByUser)
	group.GET("/number/:number", handlers.Card.GetByNumber)
	group.POST("", handlers.Card.Create)
	group.PATCH("/:id/block", handlers.Card.BlockingById)
	group.DELETE("/:id", handlers.Card.DeleteById)
}

func registerCreditRoutes(group *gin.RouterGroup, handlers *Handlers) {
	group.GET("", handlers.Credit.GetAll)
	group.GET("/:id", handlers.Credit.GetById)
	group.GET("/user/:user_id", handlers.Credit.GetByUser)
	group.POST("", handlers.Credit.Create)
	group.POST("/:id/repay/:amount", handlers.Credit.RepayById)
	group.DELETE("/:id", handlers.Credit.DeleteById)
}

func registerDepositRoutes(group *gin.RouterGroup, handlers *Handlers) {
	group.GET("", handlers.Deposit.GetAll)
	group.GET("/:id", handlers.Deposit.GetById)
	group.GET("/user/:user_id", handlers.Deposit.GetByUser)
	group.POST("", handlers.Deposit.Create)
	group.POST("/:id/replenish/:amount", handlers.Deposit.ReplenishById)
	group.DELETE("/:id", handlers.Deposit.DeleteById)
}

func registerCurrencyRoutes(group *gin.RouterGroup, handlers *Handlers) {
	group.GET("", handlers.Currency.GetAll)
	group.GET("/:id", handlers.Currency.GetById)
	group.GET("/iso/:iso_code", handlers.Currency.GetByIso)
	group.GET("/symbol/:symbol", handlers.Currency.GetBySymbol)
	group.POST("", handlers.Currency.Create)
	group.PUT("/:id", handlers.Currency.UpdateById)
	group.DELETE("/:id", handlers.Currency.DeleteById)
}

func registerExchangeRateRoutes(group *gin.RouterGroup, handlers *Handlers) {
	group.GET("/:currency_iso_from", handlers.ExchangeRate.GetAllRanking)
	group.GET("/:currency_iso_from/:currency_iso_to", handlers.ExchangeRate.GetRelativeRanking)
}
