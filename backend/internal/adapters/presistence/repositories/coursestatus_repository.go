package repositories

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type courseStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewCourseStatusRepository(db *gorm.DB) *courseStatusRepositoryImpl {
	return &courseStatusRepositoryImpl{db: db}
}

func (s *courseStatusRepositoryImpl) FindAll() (*[]entities.CourseStatus, error) {
	statuses := []models.CourseStatus{}

	if err := s.db.Find(&statuses).Error; err != nil {
		return nil, err
	}

	result := make([]entities.CourseStatus, len(statuses))
	for i, l := range statuses {
		result[i] = *l.ToEntity()
	}

	return &result, nil
}
