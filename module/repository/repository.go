package repository

import (
	"github.com/tsukiamaoto/bookShop-go/model"
	repo "github.com/tsukiamaoto/bookShop-go/module/repository/implement"

	"gorm.io/gorm"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
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
	GetCartByUserId(userId uint) (*model.Cart, error)
	CreateCartWithUserId(userId uint) error
	AddCartItemByUserId(cartItem *model.CartItem, userId uint) error
	UpdateCartItemById(cartItem *model.CartItem, cartItemId uint) error
	DeleteCartItem(userId uint, cartItemId uint) error
}

type Orders interface {
	GetOrderByUserId(userId uint) (*model.Order, error)
	CreateOrderWithUserId(userId uint) error
	UpdateOrderByUserId(order *model.Order, userId uint) error
	UpdateTotalByOrderItemAndUserId(orderItem *model.OrderItem, userId uint) error
}

type Sellers interface {
	GetProductListByUserId(userId uint) ([]*model.Product, error)
	CreateSellerWithUserId(userId uint) error
	AddProductByUserId(product *model.Product, userId uint) error
	UpdateProduct(product *model.Product) error
	DeleteProductByUserId(productId, userId uint) error
}

type Products interface {
	GetProductList(query model.Query) ([]*model.Product, paginator.Cursor, error)
	GetProductById(productId uint) (*model.Product, error)
	GetTypeList() ([][]*model.Type, error)
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
