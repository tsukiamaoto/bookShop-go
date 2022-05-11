package implement

import (
	"github.com/tsukiamaoto/bookShop-go/model"
	repo "github.com/tsukiamaoto/bookShop-go/module/repository"
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

func (p *ProductsService) GetTypeList() ([][]*model.Type, error) {
	return p.repo.GetTypeList()
}
