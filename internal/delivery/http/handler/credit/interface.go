package credit

import "github.com/gin-gonic/gin"

type ICreditHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	Create(*gin.Context)
	RepayById(*gin.Context)
	DeleteById(*gin.Context)
}
