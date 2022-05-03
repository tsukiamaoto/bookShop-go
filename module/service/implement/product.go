package implement

import (
	"tsukiamaoto/bookShop-go/model"
	repo "tsukiamaoto/bookShop-go/module/repository"
)

type ProductsService struct {
	repo repo.Products
}

func NewProductsService(repo repo.Products) *ProductsService {
	return &ProductsService{
		repo: repo,
	}
}

func (p *ProductsService) GetProductList() ([]*model.Product, error) {
	return p.repo.GetProductList()
}

func (p *ProductsService) GetProductById(productId uint) (*model.Product, error) {
	return p.repo.GetProductById(productId)
}
