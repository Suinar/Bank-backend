package card

import "github.com/gin-gonic/gin"

type ICardHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	GetByNumber(*gin.Context)
	Create(*gin.Context)
	BlockingById(*gin.Context)
	DeleteById(*gin.Context)
}
