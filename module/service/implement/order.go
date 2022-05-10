package implement

import (
	"github.com/tsukiamaoto/bookShop-go/model"
	repo "github.com/tsukiamaoto/bookShop-go/module/repository"
)

type OrdersService struct {
	repo repo.Orders
}

func NewOrdersService(repo repo.Orders) *OrdersService {
	return &OrdersService{
		repo: repo,
	}
}

func (o *OrdersService) GetOrderByUserId(userId uint) (*model.Order, error) {
	return o.repo.GetOrderByUserId(userId)
}

func (o *OrdersService) CreateOrderWithUserId(userId uint) error {
	return o.repo.CreateOrderWithUserId(userId)
}

func (o *OrdersService) UpdateOrderByUserId(order *model.Order, userId uint) error {
	return o.repo.UpdateOrderByUserId(order, userId)
}

func (o *OrdersService) UpdateTotalByOrderItemAndUserId(orderItem *model.OrderItem, userId uint) error {
	return o.repo.UpdateTotalByOrderItemAndUserId(orderItem, userId)
}
