package credit

import "github.com/gin-gonic/gin"

// ICreditHandler defines the credit endpoints registered by the router.
type ICreditHandler interface {
	GetAll(*gin.Context)
	GetByUser(*gin.Context)
	GetById(*gin.Context)
	Create(*gin.Context)
	RepayById(*gin.Context)
	DeleteById(*gin.Context)
}
