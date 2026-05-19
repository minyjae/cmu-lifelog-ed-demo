package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type CourseStatusRepository interface {
	FindAll() (*[]entities.CourseStatus, error)
}
