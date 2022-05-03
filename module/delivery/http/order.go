package http

import (
	"tsukiamaoto/bookShop-go/config"
	"tsukiamaoto/bookShop-go/middleware"
	"tsukiamaoto/bookShop-go/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) initOrderRoutes(api *gin.RouterGroup, conf *config.Config) {
	orders := api.Group("/order")
	orders.Use(
		middleware.AuthRequired,
	)
	{
		orders.GET("", handler.GetOrderByUserId)
		orders.POST("", handler.UpdateOrderByUserId)
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
// @ModuleID UpdateOrderByUserId
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=string} "Created orderItem successfully!"
// @Failure 500 string parameter error!
// @Router /order [post]
func (handler *Handler) UpdateOrderByUserId(c *gin.Context) {
	var (
		order = new(model.Order)
	)

	userId, _ := middleware.GetUserId(c)

	if err := c.ShouldBind(&order); err != nil || order == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if err := handler.services.Orders.UpdateOrderByUserId(order, userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	// if err := handler.services.Orders.UpdateTotalByOrderItemAndUserId(order.Total, userId); err != nil {
	// 	log.Error(err)
	// 	c.JSON(500, "Internal error!")
	// 	return
	// }

	c.JSON(200, dataResponse{Data: "Created orderItem successfully!"})
}
