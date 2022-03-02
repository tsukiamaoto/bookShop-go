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
		carts.GET("", handler.GetCartByUserId)
		carts.POST("", handler.AddCartItemByUserId)
		carts.PUT("/:cartItemId", handler.UpdateCartItemById)
		carts.DELETE("/:cartItemId", handler.DeleteCartItem)
	}
}

// @Summary Get CartItem List
// @Tags Cart
// @Description get cart item list by user id
// @ModuleID GetCartByUserId
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=model.Cart} "get cart"
// @Failure 500 string Internal error
// @Router /cart [get]
func (handler *Handler) GetCartByUserId(c *gin.Context) {
	userId, _ := middleware.GetUserId(c)
	if cartItemList, err := handler.services.Carts.GetCartByUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, dataResponse{Data: cartItemList})
	}
}

// @Summary Add CartItem
// @Tags Cart
// @Description add cart item by user id
// @ModuleID addCartItemByUserId
// @Accept json
// @Produce json
// @Param product formData model.Product true "product"
// @Param quantity formData int true "quantity"
// @Success 200 {object} dataResponse{data=string} "Added cartItem successfully!"
// @Failure 500 string Internal error
// @Router /cart [post]
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

	c.JSON(200, dataResponse{Data: "Added cartItem successfully!"})
}

// @Summary Update CartItem
// @Tags Cart
// @Description updated cart item by cart item id
// @ModuleID updateCartItemById
// @Accept json
// @Produce json
// @Param cartItemId path string true "cartItem id"
// @Success 200 {object} dataResponse{data=string} "Updated cartItem successfully!"
// @Failure 500 string parameter error!
// @Router /cart/:cartItemId [put]
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

	c.JSON(200, dataResponse{Data: "Updated cartItem successfully!"})
}

// @Summary Delete CartItem
// @Tags Cart
// @Description delete cart item by cart item id
// @ModuleID deleteCartItem
// @Accept json
// @Produce json
// @Param cartItemId path string true "cartItem id"
// @Success 200 {object} dataResponse{data=string} "Deleted cartItem successfully!"
// @Failure 500 string parameter error!
// @Router /cart/:cartItemId [delete]
func (handler *Handler) DeleteCartItem(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("cartItemId"), 10, 64)
	cartItemId := uint(uid64)

	userId, _ := middleware.GetUserId(c)

	if err := handler.services.Carts.DeleteCartItem(userId, cartItemId); err != nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}

	c.JSON(200, dataResponse{Data: "Deleted cartItem successfully!"})
}
