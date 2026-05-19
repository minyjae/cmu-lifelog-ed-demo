package entities

import "time"

type StaffStatus struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CourseStatus struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	Type      string    ` json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
