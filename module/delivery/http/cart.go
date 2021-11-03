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

func (handler *Handler) initCartRoutes(api *gin.RouterGroup, conf *config.Config) {
	carts := api.Group("/cart")
	carts.Use(
		cors.New(middleware.CorsConfig(conf)),
		middleware.AuthRequired,
	)
	{
		carts.GET("", handler.GetCartItemListByUserId)
		carts.POST("", handler.AddCartItemByUserId)
		carts.PUT("/:cartItemId", handler.UpdateCartItemById)
		carts.DELETE("/:cartItemId", handler.DeleteCartItem)
	}
}

func (handler *Handler) GetCartItemListByUserId(c *gin.Context) {
	userId, _ := middleware.GetUserId(c)
	if cartItemList, err := handler.services.Carts.GetCartItemListByUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, cartItemList)
	}
}

func (handler *Handler) AddCartItemByUserId(c *gin.Context) {
	var (
		cartItem = new(model.CartItem)
	)

	userId, _ := middleware.GetUserId(c)

	if err := c.ShouldBind(&cartItem); err != nil || cartItem == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if err := handler.services.Carts.AddCartItemByUserId(cartItem, userId); err != nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	c.JSON(200, "Added cartItem successfully!")
}

func (handler *Handler) UpdateCartItemById(c *gin.Context) {
	var cartItem = new(model.CartItem)

	uid64, _ := strconv.ParseUint(c.Params.ByName("cartItemId"), 10, 64)
	cartItemId := uint(uid64)

	if err := c.ShouldBind(&cartItem); err != nil || cartItem == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	if err := handler.services.Carts.UpdateCartItemById(cartItem, cartItemId); err != nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	c.JSON(200, "Updated cartItem successfully!")
}

func (handler *Handler) DeleteCartItem(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("cartItemId"), 10, 64)
	cartItemId := uint(uid64)

	if err := handler.services.Carts.DeleteCartItem(cartItemId); err != nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	c.JSON(200, "Deleted cartItem successfully!")
}
