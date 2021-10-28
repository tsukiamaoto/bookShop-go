package delivery

import (
	"shopCart/config"
	"shopCart/module/service"
	v1 "shopCart/module/delivery/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUserList(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	ModifyUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(conf *config.Config) *gin.Engine {
	// create server
	router := gin.Default()
	
	// init api
	h.initApi(router, conf)
	
	return router
}

func (h *Handler) initApi(router *gin.Engine, conf *config.Config){
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api, conf)
	}
}