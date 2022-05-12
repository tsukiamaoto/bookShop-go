package http

import (
	"strconv"

	"github.com/tsukiamaoto/bookShop-go/config"
	"github.com/tsukiamaoto/bookShop-go/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) initProductRoutes(api *gin.RouterGroup, conf *config.Config) {
	products := api.Group("/product")
	{
		products.GET("", handler.GetProductList)
		products.GET("/:productId", handler.GetProductById)
		products.GET("/types", handler.GetTypeList)
	}
}

// @Summary Get ProductList
// @Tags Product
// @Description get product list
// @ModuleID getProductList
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]model.Product} "get products"
// @Failure 500 string parameter error!
// @Router /product [get]
func (handler *Handler) GetProductList(c *gin.Context) {
	var query model.Query
	if err := c.BindQuery(&query); err != nil {
		log.Error("Failed to bind json with query, the reason is ", err)
	}

	if products, cursor, err := handler.services.Products.GetProductList(query); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		var nextPage, prevPage = "", ""
		if cursor.After != nil {
			nextPage = *cursor.After
		}
		if cursor.Before != nil {
			prevPage = *cursor.Before
		}

		c.JSON(200, dataResponse{
			Data:     products,
			NextPage: nextPage,
			PrevPage: prevPage,
		})
	}
}

// @Summary Get Product
// @Tags Product
// @Description get product by product id
// @ModuleID getProductById
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]model.Product} "get the product"
// @Failure 500 string parameter error!
// @Router /product/:productId [get]
func (handler *Handler) GetProductById(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	productId := uint(uid64)

	if product, err := handler.services.Products.GetProductById(productId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, dataResponse{Data: product})
	}
}

func (handler *Handler) GetTypeList(c *gin.Context) {
	if typeList, err := handler.services.Products.GetTypeList(); err != nil {
		log.Error("Failed to get typeList, the reason is", err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, dataResponse{Data: typeList})
	}
}
