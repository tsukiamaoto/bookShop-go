package service

import (
	"github.com/tsukiamaoto/bookShop-go/model"
	repo "github.com/tsukiamaoto/bookShop-go/module/repository"
	service "github.com/tsukiamaoto/bookShop-go/module/service/implement"
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
	GetProductList() ([]*model.Product, error)
	GetProductById(productId uint) (*model.Product, error)
}

type Services struct {
	Users    Users
	Carts    Carts
	Orders   Orders
	Sellers  Sellers
	Products Products
}

type Repos struct {
	Repos *repo.Repositories
}

func NewServices(repos Repos) *Services {
	usersService := service.NewUsersService(repos.Repos.Users)
	cartsService := service.NewCartsService(repos.Repos.Carts)
	ordersService := service.NewOrdersService(repos.Repos.Orders)
	sellersService := service.NewSellersService(repos.Repos.Sellers)
	productsService := service.NewProductsService(repos.Repos.Products)

	return &Services{
		Users:    usersService,
		Carts:    cartsService,
		Orders:   ordersService,
		Sellers:  sellersService,
		Products: productsService,
	}
}
