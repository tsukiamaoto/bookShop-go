package implement

import (
	"tsukiamaoto/bookShop-go/model"
	repo "tsukiamaoto/bookShop-go/module/repository"
)

type SellersService struct {
	repo repo.Sellers
}

func NewSellersService(repo repo.Sellers) *SellersService {
	return &SellersService{
		repo: repo,
	}
}

func (s *SellersService) GetProductListByUserId(userId uint) ([]*model.Product, error) {
	return s.repo.GetProductListByUserId(userId)
}

func (s *SellersService) CreateSellerWithUserId(userId uint) error {
	return s.repo.CreateSellerWithUserId(userId)
}

func (s *SellersService) AddProductByUserId(product *model.Product, userId uint) error {
	return s.repo.AddProductByUserId(product, userId)
}

func (s *SellersService) UpdateProduct(product *model.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *SellersService) DeleteProductByUserId(productId, userId uint) error {
	return s.repo.DeleteProductByUserId(productId, userId)
}
