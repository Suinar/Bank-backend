package account

import "github.com/gin-gonic/gin"

// IAccountHandler defines the account endpoints registered by the router.
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
