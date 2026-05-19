package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type StaffStatus struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Status    string      `gorm:"not null;unique" json:"status"`
	Type      string      `gorm:"not null" json:"type"`
	IsActive  bool        `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Lists     []ListQueue `gorm:"foreignKey:StaffStatusID"`
}

func (s *StaffStatus) ToEntity() *entities.StaffStatus {
	return &entities.StaffStatus{
		ID:       s.ID,
		Status:   s.Status,
		Type:     s.Type,
		IsActive: s.IsActive,
	}
}

func (s *StaffStatus) FromEntity(entity *entities.StaffStatus) {
	s.ID = entity.ID
	s.Status = entity.Status
	s.Type = entity.Type
	s.IsActive = entity.IsActive
}
