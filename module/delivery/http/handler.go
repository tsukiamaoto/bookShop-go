package http

import (
	"tsukiamaoto/bookShop-go/config"
	"tsukiamaoto/bookShop-go/module/service"

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
		h.initCartRoutes(v1, conf)
		h.initOrderRoutes(v1, conf)
		h.initSellerRoutes(v1, conf)
		h.initProductRoutes(v1, conf)
	}

}
