package implement

import (
	"github.com/tsukiamaoto/bookShop-go/model"
	repo "github.com/tsukiamaoto/bookShop-go/module/repository"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

type ProductsService struct {
	repo repo.Products
}

func NewProductsService(repo repo.Products) *ProductsService {
	return &ProductsService{
		repo: repo,
	}
}

func (p *ProductsService) GetProductList(query model.Query) ([]*model.Product, paginator.Cursor, error) {
	return p.repo.GetProductList(query)
}

func (p *ProductsService) GetProductById(productId uint) (*model.Product, error) {
	return p.repo.GetProductById(productId)
}

func (p *ProductsService) GetTypeList() ([][]*model.Type, error) {
	return p.repo.GetTypeList()
}
