package implement

import (
	"reflect"

	"github.com/tsukiamaoto/bookShop-go/model"

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
	var seller model.Seller
	if err := s.db.Model(&model.Seller{UserID: userId}).Find(&seller).Error; err != nil {
		return nil, err
	}

	if err := s.db.Preload("Products.Categories").Find(&seller).Error; err != nil {
		return nil, err
	}

	products := seller.Products

	return products, nil
}

func (s *SellerRepository) CreateSellerWithUserId(userId uint) error {
	err := s.db.Create(&model.Seller{UserID: userId}).Error

	return err
}

func (s *SellerRepository) AddProductByUserId(product *model.Product, userId uint) error {
	// found seller by userId
	var seller *model.Seller
	if err := s.db.Model(&model.Seller{UserID: userId}).First(&seller).Error; err != nil {
		return err
	}

	// added a new product assocation with seller
	if err := s.db.Model(&seller).Association("Products").Append(product); err != nil {
		return err
	}

	// added categories assocation with product
	if err := s.db.Model(&product).Association("Categories").Replace(&product.Categories); err != nil {
		return err
	}

	return nil
}

func (s *SellerRepository) UpdateProduct(product *model.Product) error {
	var categories = make([]model.Category, 0)
	var updatedCategories = make([]model.Category, 0)

	// find all origin categories of the product
	if err := s.db.Model(&model.Product{ID: product.ID}).Association("Categories").Find(&categories); err != nil {
		return err
	}

	// find updated categories
	for _, pCategory := range product.Categories {
		for _, category := range categories {
			if !reflect.DeepEqual(category, pCategory) ||
				!(category.Price == pCategory.Price) ||
				!(category.Inventory == pCategory.Inventory) {
				updatedCategories = append(updatedCategories, pCategory)
				break
			}
			if !reflect.DeepEqual(category.Images, pCategory.Images) {
				updatedCategories = append(updatedCategories, pCategory)
				break
			}
		}
	}

	// ignored categories to update product
	if err := s.db.Model(&model.Product{ID: product.ID}).Omit("Categories").Updates(&product).Error; err != nil {
		return err
	}

	// replaced origin categories of product to upated categories
	if err := s.db.Model(&model.Product{ID: product.ID}).Association("Categories").Replace(&updatedCategories); err != nil {
		return err
	}

	return nil
}

func (s *SellerRepository) DeleteProductByUserId(productId, userId uint) error {
	// delete product
	if err := s.db.Delete(&model.Product{ID: productId}).Error; err != nil {
		return err
	}

	// delete product association with seller
	err := s.db.Model(&model.Seller{UserID: userId}).Association("Products").Delete(&model.Product{ID: productId})

	return err
}
