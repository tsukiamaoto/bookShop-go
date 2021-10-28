package http

import (
	"shopCart/config"
	"shopCart/module/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup, conf *config.Config) {
	v1 := api.Group("/v1")
	{
		h.initUserRoutes(v1, conf)
	}
	
}