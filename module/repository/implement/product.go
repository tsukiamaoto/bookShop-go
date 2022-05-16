package implement

import (
	"fmt"

	"github.com/tsukiamaoto/bookShop-go/model"
	"github.com/tsukiamaoto/bookShop-go/utils"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
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

func (p *ProductRepository) GetProductList(query model.Query) ([]*model.Product, paginator.Cursor, error) {
	var (
		err      error
		products = make([]*model.Product, 0)
	)

	keys, associationKeys := utils.ConvertProductModelColumn2MapString()

	var assocationOrder string
	if key, ok := associationKeys[query.SortType]; ok {
		assocationOrder = fmt.Sprintf("categories.%s %s", key, query.Order)
	}

	var whereStmt string
	if key, ok := keys[query.Column]; ok {
		whereStmt = fmt.Sprintf("%s = ?", key)
	}
	if key, ok := associationKeys[query.Column]; ok {
		if key == "types" {
			whereStmt = fmt.Sprintf("? = any(%s)", key)
		} else {
			whereStmt = fmt.Sprintf("categories.%s = ?", key)
		}
	}

	var stmt *gorm.DB
	if whereStmt != "" {
		stmt = p.db.Table("products").Select("products.*, categories.*").
			Joins("inner join product_categories on products.id = product_categories.product_id").
			Joins("inner join categories on categories.id = product_categories.category_id").
			Order(assocationOrder).
			Preload("Categories").
			Where(whereStmt, query.Content).
			Find(&products)
	} else {
		stmt = p.db.Table("products").Select("products.*, categories.*").
			Joins("inner join product_categories on products.id = product_categories.product_id").
			Joins("inner join categories on categories.id = product_categories.category_id").
			Order(assocationOrder).
			Preload("Categories").
			Find(&products)
	}
	
	pagination := utils.CreatePaginator(query, keys)

	_, cursor, err := pagination.Paginate(stmt, &products)
	if err != nil {
		return nil, paginator.Cursor{}, err
	}

	return products, cursor, nil
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
