package implement

import (
	"github.com/tsukiamaoto/bookShop-go/model"

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

func (p *ProductRepository) GetTypeList() ([][]*model.Type, error) {
	var types = make([]*model.Type, 0)
	if err := p.db.Preload("Parent").Find(&types).Error; err != nil {
		return nil, err
	}

	// calculated level size
	typeLevelSize := 0
	for _, typeValue := range types {
		if typeLevelSize <= typeValue.Level {
			typeLevelSize = typeValue.Level + 1
		}
	}

	var typeList = make([][]*model.Type, typeLevelSize)
	for _, typeValue := range types {
		typeList[typeValue.Level] = append(typeList[typeValue.Level], typeValue)
	}

	return typeList, nil
}
