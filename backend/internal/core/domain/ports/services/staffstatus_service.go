package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type StaffStatusService interface {
	// @Create
	CreateStaffStatus(status *entities.StaffStatus) (*entities.StaffStatus, error)

	// @Read
	GetStaffStatus() (*[]entities.StaffStatus, error)
	GetStaffStatusByID(id uint) (*entities.StaffStatus, error)

	// @Update
	UpdateStaffStatusName(id uint, name string) (*entities.StaffStatus, error)

	// @Delete
	RemoveStaffStatus(id uint) error
}
