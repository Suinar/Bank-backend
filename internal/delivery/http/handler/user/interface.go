package user

import "github.com/gin-gonic/gin"

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
