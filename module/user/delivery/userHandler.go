package delivery

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetUserList(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	ModifyUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
