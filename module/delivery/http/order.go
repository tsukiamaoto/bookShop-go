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

// @Summary Get Order
// @Tags Order
// @Description get order by user id
// @ModuleID getOrderByUserId
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=model.Order} "get the order"
// @Failure 500 string parameter error!
// @Router /order [get]
func (handler *Handler) GetOrderByUserId(c *gin.Context) {
	userId, _ := middleware.GetUserId(c)
	if order, err := handler.services.Orders.GetOrderByUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, dataResponse{Data: order})
	}
}

// @Summary Add OrderItem
// @Tags Order
// @Description add order item by user id
// @ModuleID addOrderItemByUserId
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=string} "Created orderItem successfully!"
// @Failure 500 string parameter error!
// @Router /order [post]
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

	c.JSON(200, dataResponse{Data: "Created orderItem successfully!"})
}
