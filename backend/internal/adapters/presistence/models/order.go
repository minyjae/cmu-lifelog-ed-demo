package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

// Order table
type Order struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Title         string         `gorm:"not null" json:"title"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	OrderMappings []OrderMapping `gorm:"foreignKey:OrderID" json:"order_mappings"`
}

func (o *Order) ToEntity() *entities.Order {
	return &entities.Order{
		ID:        o.ID,
		Title:     o.Title,
		IsActive:  o.IsActive,
		UpdatedAt: o.UpdatedAt,
		CreatedAt: o.CreatedAt,
	}
}

func (o *Order) FromEntity(entity *entities.Order) {
	o.ID = entity.ID
	o.Title = entity.Title
	o.IsActive = entity.IsActive
	o.UpdatedAt = entity.UpdatedAt
	o.CreatedAt = entity.CreatedAt
}
