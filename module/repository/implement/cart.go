package implement

import (
	"shopCart/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (c *CartRepository) GetCartItemListByUserId(userId uint) ([]*model.CartItem, error) {
	var (
		err       error
		cartItems = make([]*model.CartItem, 0)
	)

	err = c.db.Model(&model.Cart{UserID: userId}).Association("CartItems").Find(&cartItems)

	return cartItems, err
}

func (c *CartRepository) CreateCartWithUserId(userId uint) error {
	err := c.db.Create(&model.Cart{UserID: userId}).Error

	return err
}

func (c *CartRepository) AddCartItemByUserId(cartItem *model.CartItem, userId uint) error {
	// create a cartItem
	if err := c.db.Create(&cartItem).Error; err != nil {
		return err
	}

	// add the cartItem to cart
	if err := c.db.Model(&model.Cart{UserID: userId}).Association("CartItems").Append(&cartItem); err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) UpdateCartItemById(cartItem *model.CartItem, cartItemId uint) error {
	// find the cartItem and update
	err := c.db.Model(&model.CartItem{ID: cartItemId}).Updates(cartItem).Error

	return err
}

func (c *CartRepository) DeleteCartItem(cartItemId uint) error {
	err := c.db.Delete(&model.CartItem{ID: cartItemId}).Error

	return err
}
