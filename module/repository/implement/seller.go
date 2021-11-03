package implement

import (
	"shopCart/model"

	"gorm.io/gorm"
)

type SellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{
		db: db,
	}
}

func (s *SellerRepository) GetProductListByUserId(userId uint) ([]*model.Product, error) {
	var (
		err      error
		products = make([]*model.Product, 0)
	)

	err = s.db.Model(&model.Seller{UserId: userId}).Association("Products").Find(&products)

	return products, err
}

func (s *SellerRepository) AddProductByUserId(product *model.Product, userId uint) error {
	// create categories and update categories id
	// categories := product.Categories
	// if err := s.db.Create(&categories).Error; err != nil{
	// 	return err
	// }
	// product.Categories = categories

	// create prodcut
	if err := s.db.Create(&product).Error; err != nil {
		return err
	}

	// add a product association with seller products
	err := s.db.Model(&model.Seller{UserId: userId}).Association("Products").Append(product)

	return err
}

func (s *SellerRepository) UpdateProduct(product *model.Product) error {
	// update product
	err := s.db.Model(&model.Product{ID: product.ID}).Updates(&product).Error

	return err
}

func (s *SellerRepository) DeleteProductByUserId(productId, userId uint) error {
	// delete product
	if err := s.db.Delete(&model.Product{ID: productId}).Error; err != nil {
		return err
	}

	// delete product association with seller
	err := s.db.Model(&model.Seller{UserId: userId}).Association("Products").Delete(&model.Product{ID: productId})

	return err
}
