package repositories

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type ListQueueRepository interface {
	Create(q *entities.ListQueue) (*entities.ListQueue, error)
	Save(q *entities.ListQueue) error
	Update(q *entities.ListQueue) (*entities.ListQueue, error)

	FindByID(id uint) (*entities.ListQueue, error)
	FindByIDWithRelations(id uint) (*entities.ListQueue, error)
	FindAllWithFacultyAndRelations() (*[]entities.ListQueue, error)
	FindAllWithRelations() (*[]entities.ListQueue, error)
	FindNotYetWithRelation() (*[]entities.ListQueue, error)
	FindByFacultyWithRelation(uint, []uint) (*[]entities.ListQueue, error)
	FindByStaffStatusWithRelation([]uint) (*[]entities.ListQueue, error)
	FindStaffStatusInListQueue() (*[]entities.ListQueue, error)
	FindByCourseStatusWithRelation([]uint) (*[]entities.ListQueue, error)
	FindByOwnerEmailWithRelation(string, []uint) (*[]entities.ListQueue, error)

	DeleteByID(id uint) error

	ChangeStaffStatusToNone(statusID uint) error

	HasSignificantChanges(*entities.ListQueue, uint, *time.Location) (bool, error)
	UpdateFields(uint, map[string]interface{}) error
	ShiftDownFrom(from uint) error
	ShiftPriorityAndMove(uint, uint, uint) error
}
