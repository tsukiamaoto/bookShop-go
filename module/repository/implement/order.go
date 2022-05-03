package implement

import (
	"tsukiamaoto/bookShop-go/model"

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
	var order *model.Order
	// get order details
	if err := o.db.Model(&model.Order{}).Where("user_id = ?", userId).Preload("OrderItems.Product").Preload("OrderItems.Category").Preload("OrderItems.Product.Categories").Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderRepository) CreateOrderWithUserId(userId uint) error {
	err := o.db.Create(&model.Order{UserID: userId}).Error

	return err
}

func (o *OrderRepository) UpdateOrderByUserId(order *model.Order, userId uint) error {
	// found orderId
	var orderId uint
	if err := o.db.Model(&model.Order{}).Where("user_id = ?", userId).Select("id").First(&orderId).Error; err != nil {
		return err
	}

	// created OrderItem instances
	var orderItems = make([]model.OrderItem, 0)
	for _, orderItem := range order.OrderItems {
		orderItem.OrderID = orderId
		orderItems = append(orderItems, orderItem)
	}
	if err := o.db.Model(&model.OrderItem{}).Create(&orderItems).Error; err != nil {
		return err
	}

	// appended orderItems to order
	if err := o.db.Model(&model.Order{ID: orderId, UserID: userId}).Omit("OrderItems.*").Association("OrderItems").Append(&orderItems); err != nil {
		return err
	}

	return nil
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
