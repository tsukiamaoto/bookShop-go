package implement

import (
	"shopCart/model"
	repo "shopCart/module/repository"
)

type CartsService struct {
	repo repo.Carts
}

func NewCartsService(repo repo.Carts) *CartsService {
	return &CartsService{
		repo: repo,
	}
}

func (c *CartsService) GetCartItemListByUserId(userId uint) ([]*model.CartItem, error) {
	return c.repo.GetCartItemListByUserId(userId)
}

func (c *CartsService) CreateCartWithUserId(userId uint) error {
	return c.repo.CreateCartWithUserId(userId)
}

func (c *CartsService) AddCartItemByUserId(cartItem *model.CartItem, userId uint) error {
	return c.repo.AddCartItemByUserId(cartItem, userId)
}

func (c *CartsService) UpdateCartItemById(cartItem *model.CartItem, cartItemId uint) error {
	return c.repo.UpdateCartItemById(cartItem, cartItemId)
}

func (c *CartsService) DeleteCartItem(cartItemId uint) error {
	return c.repo.DeleteCartItem(cartItemId)
}
