package services

import (
	"context"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
	"github.com/redis/rueidis"
)

type courseStatusService struct {
	repoCS repoPort.CourseStatusRepository
	cache  *utils.Cache
}

const (
	cacheKeyCourseStatusAll = "course_status:all"
)

func NewCourseStatusServiceImpl(r repoPort.CourseStatusRepository, redis rueidis.Client) *courseStatusService {
	return &courseStatusService{repoCS: r, cache: utils.NewCache(redis)}
}

func (s *courseStatusService) GetCourseStatus() (*[]entities.CourseStatus, error) {
	return utils.GetOrLoad(context.Background(), s.cache, cacheKeyCourseStatusAll, s.repoCS.FindAll)
}
