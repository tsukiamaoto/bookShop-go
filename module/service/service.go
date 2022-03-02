package service

import (
	"shopCart/model"
	repo "shopCart/module/repository"
	service "shopCart/module/service/implement"
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

type Deps struct {
	Repos *repo.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := service.NewUsersService(deps.Repos.Users)
	cartsService := service.NewCartsService(deps.Repos.Carts)
	ordersService := service.NewOrdersService(deps.Repos.Orders)
	sellersService := service.NewSellersService(deps.Repos.Sellers)
	productsService := service.NewProductsService(deps.Repos.Products)

	return &Services{
		Users:    usersService,
		Carts:    cartsService,
		Orders:   ordersService,
		Sellers:  sellersService,
		Products: productsService,
	}
}
