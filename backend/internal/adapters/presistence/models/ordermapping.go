package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

// OrderMapping table (many-to-many between ListQueue and Order)
type OrderMapping struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ListQueueID uint      `gorm:"not null;index" json:"list_queue_id"`
	OrderID     uint      `gorm:"not null;index" json:"order_id"`
	Checked     bool      `gorm:"not null;default:false" json:"checked"`
	ListQueue   ListQueue `gorm:"foreignKey:ListQueueID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Order       Order     `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (o *OrderMapping) ToEntity() *entities.OrderMapping {
	return &entities.OrderMapping{
		ID:          o.ID,
		ListQueueID: o.ListQueueID,
		OrderID:     o.OrderID,
		Order:       *o.Order.ToEntity(),
		Checked:     o.Checked,
	}
}

func (o *OrderMapping) FromEntity(entity *entities.OrderMapping) {
	o.ID = entity.ID
	o.ListQueueID = entity.ListQueueID
	o.OrderID = entity.OrderID
	o.Order = Order{}
	o.Checked = entity.Checked
}
