package services

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type ListQueueService interface {
	// @Create
	CreateListQueue(req *entities.ListQueue) (*entities.ListQueue, error)

	// @Read
	GetListQueue() (*[]entities.ListQueue, error)
	GetListQueueNotYet() (*[]entities.ListQueue, error)
	GetListQueueByID(id uint) (*entities.ListQueue, error)
	GetListQueueByStaffStatus([]uint) (*[]entities.ListQueue, error)
	GetListQueueByFaculty(f string, ids []uint) (*[]entities.ListQueue, error)
	GetListQueueByCourseStatus([]uint) (*[]entities.ListQueue, error)
	GetListQueueByOwner(string, []uint) (*[]entities.ListQueue, error)

	// @Update
	UpdateListQueue(req *entities.ListQueue) (*entities.ListQueue, error)
	UpdateStaffStatus(id, statusID uint) (*entities.ListQueue, error)
	UpdatePriority(uint, uint) (*entities.ListQueue, error)

	// @Delete
	RemoveListQueueForDev(id uint) error

	// @Custom
	DaysBetweenCeil(a, b time.Time, loc *time.Location) int
	DecorateDateLeft(lq *entities.ListQueue)
}
