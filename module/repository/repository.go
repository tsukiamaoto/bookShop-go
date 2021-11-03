package repository

import (
	"shopCart/model"
	repo "shopCart/module/repository/implement"

	"gorm.io/gorm"
)

type Users interface {
	GetUserList() ([]*model.User, error)
	GetUser(user *model.User) (*model.User, error)
	GetUserById(userId uint) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(userId uint) error
}

type Carts interface {
	GetCartItemListByUserId(userId uint) ([]*model.CartItem, error)
	CreateCartWithUserId(userId uint) error
	AddCartItemByUserId(cartItem *model.CartItem, userId uint) error
	UpdateCartItemById(cartItem *model.CartItem, cartItemId uint) error
	DeleteCartItem(cartItemId uint) error
}

type Orders interface {
	GetOrderByUserId(userId uint) (*model.Order, error)
	CreateOrderWithUserId(userId uint) error
	AddOrderItemByUserId(orderItem *model.OrderItem, userId uint) error
	UpdateTotalByOrderItemAndUserId(orderItem *model.OrderItem, userId uint) error
}

type Sellers interface {
	GetProductListByUserId(userId uint) ([]*model.Product, error)
	AddProductByUserId(product *model.Product, userId uint) error
	UpdateProduct(product *model.Product) error
	DeleteProductByUserId(productId, userId uint) error
}

type Products interface {
	GetProductList() ([]*model.Product, error)
	GetProductById(productId uint) (*model.Product, error)
}

type Repositories struct {
	Users    Users
	Carts    Carts
	Orders   Orders
	Sellers  Sellers
	Products Products
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:    repo.NewUserRepository(db),
		Carts:    repo.NewCartRepository(db),
		Orders:   repo.NewOrderRepository(db),
		Products: repo.NewProductRepository(db),
		Sellers:  repo.NewSellerRepository(db),
	}
}
