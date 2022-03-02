package implement

import (
	"shopCart/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) GetProductList() ([]*model.Product, error) {
	var (
		err      error
		products = make([]*model.Product, 0)
	)

	if err = p.db.Preload("Categories").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) GetProductById(productId uint) (*model.Product, error) {
	var product *model.Product
	err := p.db.Preload("Categories").Model(&model.Product{ID: productId}).First(&product).Error

	return product, err
}
