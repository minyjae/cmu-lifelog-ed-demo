package entities

import "time"

type Order struct {
	ID            uint           `json:"id"`
	Title         string         `json:"title"`
	IsActive      bool           `json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	OrderMappings []OrderMapping `json:"order_mappings"`
}

// OrderMapping table (many-to-many between ListQueue and Order)
type OrderMapping struct {
	ID          uint      `json:"id"`
	ListQueueID uint      `json:"list_queue_id"`
	OrderID     uint      `json:"order_id"`
	Checked     bool      `json:"checked"`
	ListQueue   ListQueue `json:"-"`
	Order       Order     `json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
