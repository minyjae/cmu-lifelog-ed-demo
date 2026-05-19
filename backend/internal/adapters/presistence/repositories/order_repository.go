package repositories

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepositoryImpl {
	return &orderRepositoryImpl{db: db}
}

func (s *orderRepositoryImpl) Create(o *entities.Order) (*entities.Order, error) {
	order := models.Order{}
	order.FromEntity(o)

	if err := s.db.Create(&order).Error; err != nil {
		return nil, err
	}

	o.ID = order.ID
	result := order.ToEntity()

	return result, nil
}

func (s *orderRepositoryImpl) FindByID(id uint) (*entities.Order, error) {
	order := models.Order{}

	if err := s.db.First(&order, id).Error; err != nil {
		return nil, err
	}

	result := order.ToEntity()
	return result, nil
}

func (s *orderRepositoryImpl) FindAll() (*[]entities.Order, error) {
	order := []models.Order{}
	if err := s.db.Where("is_active = ?", true).Find(&order).Error; err != nil {
		return nil, err
	}

	result := make([]entities.Order, len(order))
	for i, m := range order {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *orderRepositoryImpl) Save(o *entities.Order) (*entities.Order, error) {
	order := models.Order{}
	order.FromEntity(o)
	if err := s.db.Save(&order).Error; err != nil {
		return nil, err
	}

	result := order.ToEntity()
	return result, nil
}

func (s *orderRepositoryImpl) FindByIDWithMappings(id uint) (*entities.Order, error) {
	order := models.Order{}

	if err := s.db.Preload("OrderMappings").
		Find(&order, id).Error; err != nil {
		return nil, err
	}

	result := order.ToEntity()

	return result, nil
}
