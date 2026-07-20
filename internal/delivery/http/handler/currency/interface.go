package currency

import "github.com/gin-gonic/gin"

type ICurrencyHandler interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	GetByIso(*gin.Context)
	GetBySymbol(*gin.Context)
	Convert(*gin.Context)
	Create(*gin.Context)
	UpdateById(*gin.Context)
	DeleteById(*gin.Context)
	GetAllRanking(*gin.Context)
	GetRelativeRanking(*gin.Context)
}
