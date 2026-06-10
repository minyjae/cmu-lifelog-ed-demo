package services

import (
	"context"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
	"github.com/redis/rueidis"
)

type roleService struct {
	repoR repoPort.RoleRepository
	cache *utils.Cache
}

const (
	cacheKeyRoleAll = "role:all"
)

func NewRoleServiceImpl(r repoPort.RoleRepository, redis rueidis.Client) *roleService {
	return &roleService{repoR: r, cache: utils.NewCache(redis)}
}

func (s *roleService) GetRole() (*[]entities.Role, error) {
	// roles, err := s.repoR.FindAll()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to find course status: %w", err)
	// }

	// return roles, nil

	return utils.GetOrLoad(context.Background(), s.cache, cacheKeyRoleAll, s.repoR.FindAll)
}
