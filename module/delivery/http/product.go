package http

import (
	"shopCart/config"
	"shopCart/middleware"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) initProductRoutes(api *gin.RouterGroup, conf *config.Config) {
	products := api.Group("/product")
	products.Use(
		cors.New(middleware.CorsConfig(conf)),
		middleware.AuthRequired,
	)
	{
		products.GET("", handler.GetProductList)
		products.GET("/:productId", handler.GetProductById)
	}
}

func (handler *Handler) GetProductList(c *gin.Context) {
	if productList, err := handler.services.Products.GetProductList(); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, productList)
	}
}

func (handler *Handler) GetProductById(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	productId := uint(uid64)

	if product, err := handler.services.Products.GetProductById(productId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, product)
	}
}
