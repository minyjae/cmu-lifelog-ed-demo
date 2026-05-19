package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type CourseStatus struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Status    string      `gorm:"not null;unique" json:"status"`
	Type      string      `gorm:"not null" json:"type"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Lists     []ListQueue `gorm:"foreignKey:CourseStatusID"`
}

func (s *CourseStatus) ToEntity() *entities.CourseStatus {
	return &entities.CourseStatus{
		ID:     s.ID,
		Status: s.Status,
		Type:   s.Type,
	}
}

func (s *CourseStatus) FromEntity(entity *entities.CourseStatus) {
	s.ID = entity.ID
	s.Status = entity.Status
	s.Type = entity.Type
}
