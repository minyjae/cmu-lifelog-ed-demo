package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type StaffStatusRepository interface {
	Create(status *entities.StaffStatus) (*entities.StaffStatus, error)
	FindAll() (*[]entities.StaffStatus, error)
	FindByID(id uint) (*entities.StaffStatus, error)
	FindByIDToSoftDelete(id uint) (*entities.StaffStatus, error)
	Save(*entities.StaffStatus) (*entities.StaffStatus, error)
	// FindStaffStatusByID(id uint) (*entities.StaffStatus, error)
}
