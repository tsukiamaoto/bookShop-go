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

func (c *CartRepository) GetCartByUserId(userId uint) (*model.Cart, error) {
	var cart *model.Cart

	if err := c.db.Model(&model.Cart{}).Where("user_id = ?", userId).Preload("CartItems.Product").Preload("CartItems.Category").Preload("CartItems.Product.Categories").First(&cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func (c *CartRepository) CreateCartWithUserId(userId uint) error {
	err := c.db.Create(&model.Cart{UserID: userId}).Error

	return err
}

func (c *CartRepository) AddCartItemByUserId(cartItem *model.CartItem, userId uint) error {
	// found cartId
	var cartId uint
	if err := c.db.Model(&model.Cart{}).Where("user_id = ?", userId).Select("id").First(&cartId).Error; err != nil {
		return err
	}

	// created a new cartItem
	cartItem.CartID = cartId
	if err := c.db.Model(&model.CartItem{}).Create(&cartItem).Error; err != nil {
		return err
	}

	// appended the cartItem to cart
	if err := c.db.Model(&model.Cart{ID: cartId, UserID: userId}).Omit("CartItems.*").Association("CartItems").Append(cartItem); err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) UpdateCartItemById(cartItem *model.CartItem, cartItemId uint) error {
	// found the cartItem and update
	if err := c.db.Model(&model.CartItem{ID: cartItemId}).Updates(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) DeleteCartItem(userId, cartItemId uint) error {
	// found cartId
	var cartId uint
	if err := c.db.Model(&model.Cart{}).Where("user_id = ?", userId).Select("id").First(&cartId).Error; err != nil {
		return err
	}

	// found cartItem
	var cartItem *model.CartItem
	if err := c.db.Model(&model.CartItem{}).Where("id = ?", cartItemId).First(&cartItem).Error; err != nil {
		return err
	}

	// removed the cartItem relationshop with the cart
	if err := c.db.Model(&model.Cart{ID: cartId, UserID: userId}).Association("CartItems").Delete(cartItem); err != nil {
		return err
	}

	// deleted cartItem object
	if err := c.db.Delete(cartItem).Error; err != nil {
		return err
	}

	return nil
}
