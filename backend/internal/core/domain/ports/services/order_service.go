package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type OrderService interface {
	// @Create
	CreateOrder(order *entities.Order) (*entities.Order, error)
	CreateOrderForListQueue(listQueueID uint) error

	// @Read
	GetOrderFromListQueueID(id uint) (*[]entities.Order, error)

	// @Update
	UpdateOrder(orderMapping *entities.OrderMapping) (*entities.Order, error)
	UpdateOrderName(orderID uint, title string) (*entities.Order, error)

	// @Delete
	RemoveOrder(orderID uint) error

	// @Extra functions
	AddNewOrderToListQueue(order *entities.Order) error
}
