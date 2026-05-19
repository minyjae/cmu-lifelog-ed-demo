package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type OrderRepository interface {
	Create(o *entities.Order) (*entities.Order, error)
	FindByID(id uint) (*entities.Order, error)
	FindAll() (*[]entities.Order, error)
	Save(o *entities.Order) (*entities.Order, error)

	// preload/relations
	FindByIDWithMappings(id uint) (*entities.Order, error)
}
