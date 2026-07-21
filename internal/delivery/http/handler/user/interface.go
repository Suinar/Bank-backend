package user

import "github.com/gin-gonic/gin"

// IUserHandler defines the user endpoints registered by the router.
type IUserHandler interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	GetByEmail(*gin.Context)
	GetByPhoneNumber(*gin.Context)
	Create(*gin.Context)
	UpdateById(*gin.Context)
	ChangePasswordById(*gin.Context)
	DeleteById(*gin.Context)
}
