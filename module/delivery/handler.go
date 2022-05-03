package delivery

import (
	"tsukiamaoto/bookShop-go/config"
	docs "tsukiamaoto/bookShop-go/docs"
	"tsukiamaoto/bookShop-go/middleware"
	v1 "tsukiamaoto/bookShop-go/module/delivery/http"
	"tsukiamaoto/bookShop-go/module/service"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type UserHandler interface {
	GetUserList(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	ModifyUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	Signup(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type CartHandler interface {
	GetCartByUserId(c *gin.Context)
	AddCartItemByUserId(c *gin.Context)
	UpdateCartItemById(c *gin.Context)
	DeleteCartItem(c *gin.Context)
}

type OrderHandler interface {
	GetOrderByUserId(c *gin.Context)
	UpdateOrderByUserId(c *gin.Context)
}

type SellerHandler interface {
	GetProductListBySellerId(c *gin.Context)
	GetProductById(c *gin.Context)
	AddProductByUserId(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProductByUserId(c *gin.Context)
}

type ProductHandler interface {
	GetProductList(c *gin.Context)
	GetProductById(c *gin.Context)
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
	router.Use(
		cors.New(middleware.CorsConfig(conf)),
	)

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", "localhost", 8080)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// init api
	h.initApi(router, conf)

	return router
}

func (h *Handler) initApi(router *gin.Engine, conf *config.Config) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api, conf)
	}
}
