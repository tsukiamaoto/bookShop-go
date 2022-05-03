package http

import (
	"strconv"
	"tsukiamaoto/bookShop-go/config"
	"tsukiamaoto/bookShop-go/middleware"
	"tsukiamaoto/bookShop-go/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) initSellerRoutes(api *gin.RouterGroup, conf *config.Config) {
	sellers := api.Group("/seller")
	sellers.Use(
		middleware.AuthRequired,
	)
	{
		sellers.GET("", handler.GetProductListByUserId)
		sellers.POST("", handler.AddProductByUserId)
		sellers.PUT("/:productId", handler.UpdateProduct)
		sellers.DELETE("/:productId", handler.DeleteProductByUserId)
	}
}

// @Summary Get Product List
// @Tags Seller
// @Description get product list by user id
// @ModuleID getProductListByUserId
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]model.Product} "get products"
// @Failure 500 string Internal error
// @Router /seller [get]
func (handler *Handler) GetProductListByUserId(c *gin.Context) {
	userId, _ := middleware.GetUserId(c)
	if productList, err := handler.services.Sellers.GetProductListByUserId(userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
	} else {
		c.JSON(200, dataResponse{Data: productList})
	}
}

// @Summary Add Product
// @Tags Seller
// @Description add product by user id
// @ModuleID addProductByUserId
// @Accept json
// @Produce json
// @Param name formData string true "name"
// @Param description formData string true "description"
// @Param categories formData model.Category true "categories"
// @Success 200 {object} dataResponse{data=string} "Created new product successfully!"
// @Failure 500 string parameter error!
// @Failure 500 string Internal error!
// @Router /seller [post]
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

	c.JSON(200, dataResponse{Data: "Created a new product successfully!"})
}

// @Summary Update Product
// @Tags Seller
// @Description update product
// @ModuleID updateProduct
// @Accept json
// @Produce json
// @Param name formData string true "name"
// @Param description formData string true "description"
// @Param categories formData model.Category true "category"
// @Success 200 {object} dataResponse{data=string} "Updated the product successfully!"
// @Failure 500 string parameter error!
// @Failure 500 string Internal error
// @Router /seller/:productId [put]
func (handler *Handler) UpdateProduct(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("productId"), 10, 64)
	productId := uint(uid64)

	var product = new(model.Product)

	if err := c.ShouldBind(&product); err != nil || product == nil {
		log.Error(err)
		c.JSON(500, "parameter error!")
		return
	}
	product.ID = productId

	if err := handler.services.Sellers.UpdateProduct(product); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	c.JSON(200, dataResponse{Data: "Updated the product successfully!"})
}

// @Summary Delete Product
// @Tags Seller
// @Description delete product by user id
// @ModuleID deleteproductByUserId
// @Accept json
// @Produce json
// @Param productId path string true "product id"
// @Success 200 {object} dataResponse{data=string} "Deleted new product successfully!"
// @Failure 500 string Internal error!
// @Router /seller/:productId [delete]
func (handler *Handler) DeleteProductByUserId(c *gin.Context) {
	uid64, _ := strconv.ParseUint(c.Params.ByName("productId"), 10, 64)
	productId := uint(uid64)

	userId, _ := middleware.GetUserId(c)

	if err := handler.services.Sellers.DeleteProductByUserId(productId, userId); err != nil {
		log.Error(err)
		c.JSON(500, "Internal error!")
		return
	}

	c.JSON(200, dataResponse{Data: "Deleted the product successfully!"})
}
