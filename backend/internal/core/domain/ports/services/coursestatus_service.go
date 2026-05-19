package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type CourseStatusService interface {
	GetCourseStatus() (*[]entities.CourseStatus, error)
}
