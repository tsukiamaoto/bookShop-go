package implement

import (
	"shopCart/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) GetOrderByUserId(userId uint) (*model.Order, error) {
	var (
		err   error
		order *model.Order
	)

	err = o.db.Model(&model.Order{UserID: userId}).Find(&order).Error

	return order, err
}

func (o *OrderRepository) CreateOrderWithUserId(userId uint) error {
	err := o.db.Create(&model.Order{UserID: userId}).Error

	return err
}

func (o *OrderRepository) AddOrderItemByUserId(orderItem *model.OrderItem, userId uint) error {
	// create orderItem
	if err := o.db.Create(&orderItem).Error; err != nil {
		return err
	}

	// append the orderItem to order
	err := o.db.Model(&model.Order{UserID: userId}).Association("OrderItems").Append(orderItem)

	return err
}

func (o *OrderRepository) UpdateTotalByOrderItemAndUserId(orderItem *model.OrderItem, userId uint) error {
	// get product price
	var (
		categoriesOfProduct []model.Category
		priceOfProduct      int
	)
	categoriesOfProduct = append(categoriesOfProduct, orderItem.Product.Categories...)
	priceOfProduct = categoriesOfProduct[0].Price

	// find the old total
	var total int
	if err := o.db.Model(&model.Order{UserID: userId}).Select("Total").First(&total).Error; err != nil {
		return err
	}
	// update the price of the order
	newTotal := total + priceOfProduct
	err := o.db.Model(&model.Order{UserID: userId}).UpdateColumn("Total", &newTotal).Error

	return err
}
