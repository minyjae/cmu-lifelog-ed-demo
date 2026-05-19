package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type OrderMappingRepository interface {
	Create(m *[]entities.OrderMapping) error
	Update(m *entities.OrderMapping) error

	FindByListQueueID(listQueueID uint) (*[]entities.OrderMapping, error)
	// FindByOrderID(orderID uint) (*[]entities.OrderMapping, error)
	FindByOrderIDAndListID(orderID, listQueueID uint) (*entities.OrderMapping, error)

	FindByOrderIDAndDelete(orderID uint) error

	DeleteOrder(o *[]entities.OrderMapping) error
}
