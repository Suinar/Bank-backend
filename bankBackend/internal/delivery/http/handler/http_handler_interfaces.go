package core

import (
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	GetByEmail(*gin.Context)
	GetByPhoneNumber(*gin.Context)
	GetMe(*gin.Context)
	Create(*gin.Context)
	UpdateById(*gin.Context)
	ChangePasswordById(*gin.Context)
	DeleteById(*gin.Context)
}

type IAccountHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	GetByCreateTime(*gin.Context)
	Create(*gin.Context)
	UpdateById(*gin.Context)
	DeleteById(*gin.Context)
}

type ICardHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	GetByNumber(*gin.Context)
	GetByCreateTime(*gin.Context)
	BlockingById(*gin.Context)
	Create(*gin.Context)
	DeleteById(*gin.Context)
}

type ICreditHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	GetByCreateTime(*gin.Context)
	GetByRepayTime(*gin.Context)
	Create(*gin.Context)
	RepayById(*gin.Context)
	DeleteById(*gin.Context)
}

type IDebitHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	GetByCreateTime(*gin.Context)
	GetByCompletionTime(*gin.Context)
	Create(*gin.Context)
	RepayById(*gin.Context)
	DeleteById(*gin.Context)
}

type ICurrencyHandler interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	GetByIsoCod(*gin.Context)
	GetByNumberCod(*gin.Context)
	GetBySymbol(*gin.Context)
	Convert(*gin.Context)
	Create(*gin.Context)
	UpdateById(*gin.Context)
	DeleteById(*gin.Context)
}

type IExchangeRateHandler interface {
	GetAll(*gin.Context)
	GetForCurrencies(*gin.Context)
}
