package repositories

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type orderMappingRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderMappingRepository(db *gorm.DB) *orderMappingRepositoryImpl {
	return &orderMappingRepositoryImpl{db: db}
}

func (s *orderMappingRepositoryImpl) Create(m *[]entities.OrderMapping) error {
	orderMapping := []models.OrderMapping{}

	for _, e := range *m {
		om := &models.OrderMapping{}
		om.FromEntity(&e)
		orderMapping = append(orderMapping, *om)
	}

	if err := s.db.Create(&orderMapping).Error; err != nil {
		return err
	}

	return nil
}

func (s *orderMappingRepositoryImpl) Update(m *entities.OrderMapping) error {
	orderMapping := models.OrderMapping{}
	orderMapping.FromEntity(m)

	return s.db.Save(orderMapping).Error
}

func (s *orderMappingRepositoryImpl) FindByListQueueID(listQueueID uint) (*[]entities.OrderMapping, error) {
	mapping := []models.OrderMapping{}

	if err := s.db.Where("list_queue_id = ?", listQueueID).
		Find(&mapping).Error; err != nil {
		return nil, err
	}

	result := make([]entities.OrderMapping, len(mapping))
	for i, m := range mapping {
		result[i] = *m.ToEntity()
	}

	return &result, nil

}

func (s *orderMappingRepositoryImpl) FindByOrderIDAndListID(orderID, listQueueID uint) (*entities.OrderMapping, error) {
	mapping := models.OrderMapping{}

	if err := s.db.Where("order_id = ? AND list_queue_id = ?", orderID, listQueueID).
		First(&mapping).Error; err != nil {
		return nil, err
	}

	result := mapping.ToEntity()
	return result, nil
}

func (s *orderMappingRepositoryImpl) FindByOrderIDAndDelete(orderID uint) error {
	// blocked := []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// // เช็คว่า orderID นี้อยู่ใน blocked list ไหม
	// var order models.Order
	// err := s.db.
	// 	Where("id IN ?", blocked).
	// 	Where("id = ?", orderID).
	// 	First(&order).Error

	// if err == nil {
	// 	// พบ orderID ใน list ที่ block → หยุดทันที
	// 	return fmt.Errorf("cannot delete protected order ID: %d", orderID)
	// }
	// if !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	// ถ้าไม่ใช่ ErrRecordNotFound → คือ error จริง → หยุดทันที
	// 	return err
	// }

	// ปลอดภัยแล้ว → ลบ OrderMapping ได้
	var ids []uint
	if err := s.db.
		Table("order_mappings AS om").
		Joins("JOIN list_queues AS lq ON lq.id = om.list_queue_id").
		Joins("JOIN staff_statuses AS ss ON ss.id = lq.staff_status_id").
		Where("ss.type NOT IN (?)", []string{"Done", "Cancel"}).
		Where("om.order_id = ?", orderID).
		Pluck("om.id", &ids).Error; err != nil {
		return err
	}

	if len(ids) == 0 {
		return nil
	}

	if err := s.db.
		Where("id IN ?", ids).
		Delete(&models.OrderMapping{}).Error; err != nil {
		return nil
	}
	return nil
}

func (s *orderMappingRepositoryImpl) DeleteOrder(o *[]entities.OrderMapping) error {
	if len(*o) == 0 {
		return nil
	}
	mapping := make([]models.OrderMapping, len(*o))
	for i, e := range *o {
		om := &models.OrderMapping{}
		om.FromEntity(&e)
		mapping[i] = *om
	}

	if err := s.db.Delete(&mapping).Error; err != nil {
		return err
	}

	return nil
}
