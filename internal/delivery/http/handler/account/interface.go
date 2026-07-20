package account

import "github.com/gin-gonic/gin"

type IAccountHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	Create(*gin.Context)
	BlockingById(*gin.Context)
	CloseById(*gin.Context)
	UpdateById(*gin.Context)
	DeleteById(*gin.Context)
}
