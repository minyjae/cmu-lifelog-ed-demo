package services

import (
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
)

type courseStatusService struct {
	repoCS repoPort.CourseStatusRepository
}

func NewCourseStatusServiceImpl(r repoPort.CourseStatusRepository) *courseStatusService {
	return &courseStatusService{repoCS: r}
}

func (s *courseStatusService) GetCourseStatus() (*[]entities.CourseStatus, error) {
	statuses, err := s.repoCS.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find course status: %w", err)
	}

	return statuses, nil
}
