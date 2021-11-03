package http

import (
	"shopCart/config"
	"shopCart/middleware"
	"shopCart/model"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) initSellerRoutes(api *gin.RouterGroup, conf *config.Config) {
	sellers := api.Group("/seller")
	sellers.Use(
		cors.New(middleware.CorsConfig(conf)),
		middleware.AuthRequired,
	)
	{
		sellers.GET("", handler.GetProductListByUserId)
		sellers.POST("", handler.AddOrderItemByUserId)
		sellers.PUT("/:productId", handler.UpdateProduct)
		sellers.DELETE("/:productId", handler.DeleteProductByUserId)
	}
}

func (handler *Handler) GetProductListByUserId(c *gin.Context) {
	userId, _ := middleware.GetUserId(c)
	if productList, err := handler.services.Sellers.GetProductListByUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, productList)
	}
}

func (handler *Handler) AddProductByUserId(c *gin.Context) {
	var product = new(model.Product)

	userId, _ := middleware.GetUserId(c)

	if err := c.ShouldBind(&product); err != nil || product == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if err := handler.services.Sellers.AddProductByUserId(product, userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	c.JSON(200, "Created new product successfully!")
}

func (handler *Handler) UpdateProduct(c *gin.Context) {
	var product = new(model.Product)

	if err := c.ShouldBind(&product); err != nil || product == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if err := handler.services.Sellers.UpdateProduct(product); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	c.JSON(200, "Updated the product successfully!")
}

func (handler *Handler) DeleteProductByUserId(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("productId"), 10, 64)
	productId := uint(uid64)

	userId, _ := middleware.GetUserId(c)

	if err := handler.services.Sellers.DeleteProductByUserId(productId, userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	c.JSON(200, "Deleted new product successfully!")
}
