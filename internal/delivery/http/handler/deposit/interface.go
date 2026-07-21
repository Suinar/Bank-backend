package deposit

import "github.com/gin-gonic/gin"

// IDebitHandler defines the deposit endpoints registered by the router.
type IDebitHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	Create(*gin.Context)
	ReplenishById(*gin.Context)
	DeleteById(*gin.Context)
}
