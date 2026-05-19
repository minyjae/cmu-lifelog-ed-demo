package repositories

import (
	"errors"
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type staffStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewStaffStatusRepository(db *gorm.DB) *staffStatusRepositoryImpl {
	return &staffStatusRepositoryImpl{db: db}
}

func (s *staffStatusRepositoryImpl) Create(status *entities.StaffStatus) (*entities.StaffStatus, error) {
	cs := models.StaffStatus{}
	cs.FromEntity(status)
	if err := s.db.Create(&cs).Error; err != nil {
		return nil, err
	}

	result := cs.ToEntity()

	return result, nil
}

func (s *staffStatusRepositoryImpl) FindAll() (*[]entities.StaffStatus, error) {
	status := []models.StaffStatus{}

	if err := s.db.Where("is_active = ?", true).Find(&status).Error; err != nil {
		return nil, err
	}

	result := make([]entities.StaffStatus, len(status))
	for i, m := range status {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *staffStatusRepositoryImpl) FindByID(id uint) (*entities.StaffStatus, error) {
	status := models.StaffStatus{}

	if err := s.db.Where("id = ?", id).First(&status).Error; err != nil {
		return nil, err
	}
	result := status.ToEntity()

	return result, nil
}

func (s *staffStatusRepositoryImpl) FindByIDToSoftDelete(id uint) (*entities.StaffStatus, error) {
	// 1) guard: กันลบ id ที่ห้ามลบ
	blocked := map[uint]struct{}{1: {}}
	if _, ok := blocked[id]; ok {
		return nil, fmt.Errorf("cannot delete protected status ID: %d", id)
	}

	// 2) ดึงเรคอร์ดจริง ๆ จาก DB
	var status models.StaffStatus
	if err := s.db.First(&status, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("staff status id %d not found", id)
		}
		return nil, err
	}

	// 3) คืน entity ให้ layer ถัดไปไปทำ soft delete ต่อ
	return status.ToEntity(), nil
}

func (s *staffStatusRepositoryImpl) Save(status *entities.StaffStatus) (*entities.StaffStatus, error) {
	cs := models.StaffStatus{}
	cs.FromEntity(status)
	if err := s.db.Save(&cs).Error; err != nil {
		return nil, err
	}
	result := cs.ToEntity()

	return result, nil
}
