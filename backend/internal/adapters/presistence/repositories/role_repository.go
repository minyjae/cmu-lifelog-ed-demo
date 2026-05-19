package repositories

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepositoryImpl {
	return &roleRepositoryImpl{db: db}
}

func (s *roleRepositoryImpl) FindAll() (*[]entities.Role, error) {
	roles := []models.Role{}

	if err := s.db.Find(&roles).Error; err != nil {
		return nil, err
	}

	result := make([]entities.Role, len(roles))
	for i, l := range roles {
		result[i] = *l.ToEntity()
	}

	return &result, nil
}
