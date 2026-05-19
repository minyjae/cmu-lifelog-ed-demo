package models

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Role string `gorm:"not null;unique" json:"role"`
}

func (s *Role) ToEntity() *entities.Role {
	return &entities.Role{
		ID:   s.ID,
		Role: s.Role,
	}
}

func (s *Role) FromEntity(entity *entities.Role) {
	s.ID = entity.ID
	s.Role = entity.Role
}
