package services

import (
	"log"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
)

type orderService struct {
	orderRepo   repoPort.OrderRepository
	mappingRepo repoPort.OrderMappingRepository
	listRepo    repoPort.ListQueueRepository
}

func NewOrderServiceImpl(r repoPort.OrderRepository, m repoPort.OrderMappingRepository, l repoPort.ListQueueRepository) *orderService {
	return &orderService{orderRepo: r, mappingRepo: m, listRepo: l}
}

func (s *orderService) CreateOrder(order *entities.Order) (*entities.Order, error) {
	_, err := s.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	if err := s.AddNewOrderToListQueue(order); err != nil {
		log.Printf("[CreateOrder] warning: cannot add order to list queue: %v", err)
	}

	result, err := s.orderRepo.FindByIDWithMappings(order.ID)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *orderService) CreateOrderForListQueue(listID uint) error {
	o, err := s.orderRepo.FindAll()
	if err != nil {
		return err
	}

	for _, order := range *o {
		orderMapping := entities.OrderMapping{
			ListQueueID: listID,
			OrderID:     order.ID,
		}

		err := s.mappingRepo.Create(&[]entities.OrderMapping{orderMapping})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *orderService) GetOrderFromListQueueID(id uint) (*[]entities.Order, error) {
	var orders []entities.Order
	m, err := s.mappingRepo.FindByListQueueID(id)
	if err != nil {
		return nil, err
	}

	for _, mapping := range *m {
		o, err := s.orderRepo.FindByID(mapping.OrderID)
		if err != nil {
			return nil, err
		}

		orders = append(orders, *o)
	}

	return &orders, nil
}

func (s *orderService) UpdateOrder(orderMapping *entities.OrderMapping) (*entities.Order, error) {
	m, err := s.mappingRepo.FindByOrderIDAndListID(orderMapping.OrderID, orderMapping.ListQueueID)
	if err != nil {
		return nil, err
	}

	m.Checked = orderMapping.Checked

	if err := s.mappingRepo.Update(m); err != nil {
		return nil, err
	}

	p, err := s.orderRepo.FindByIDWithMappings(orderMapping.OrderID)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *orderService) UpdateOrderName(orderID uint, title string) (*entities.Order, error) {
	o, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}
	o.Title = title

	updatedOrder, err := s.orderRepo.Save(o)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}

func (s *orderService) RemoveOrder(orderID uint) error {
	err := s.mappingRepo.FindByOrderIDAndDelete(orderID)
	if err != nil {
		return err
	}

	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return err
	}

	order.IsActive = false

	_, err = s.orderRepo.Save(order)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) AddNewOrderToListQueue(order *entities.Order) error {
	l, err := s.listRepo.FindStaffStatusInListQueue()
	if err != nil {
		return err
	}

	var mappings []entities.OrderMapping
	for _, list := range *l {
		mappings = append(mappings, entities.OrderMapping{
			ListQueueID: list.ID,
			OrderID:     order.ID,
			Checked:     false,
		})
	}

	err = s.mappingRepo.Create(&mappings)
	if err != nil {
		return err
	}

	return nil
}
