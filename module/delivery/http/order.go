package http

import (
	"shopCart/config"
	"shopCart/middleware"
	"shopCart/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) initOrderRoutes(api *gin.RouterGroup, conf *config.Config) {
	orders := api.Group("/order")
	orders.Use(
		cors.New(middleware.CorsConfig(conf)),
		middleware.AuthRequired,
	)
	{
		orders.GET("", handler.GetOrderByUserId)
		orders.POST("", handler.AddOrderItemByUserId)
	}
}

func (handler *Handler) GetOrderByUserId(c *gin.Context) {
	userId, _ := middleware.GetUserId(c)
	if order, err := handler.services.Orders.GetOrderByUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, order)
	}
}

func (handler *Handler) AddOrderItemByUserId(c *gin.Context) {
	var (
		orderItem = new(model.OrderItem)
	)

	userId, _ := middleware.GetUserId(c)

	if err := c.ShouldBind(&orderItem); err != nil || orderItem == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if err := handler.services.Orders.AddOrderItemByUserId(orderItem, userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	if err := handler.services.Orders.UpdateTotalByOrderItemAndUserId(orderItem, userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	c.JSON(200, "Created orderItem successfully!")
}
