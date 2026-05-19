package services

import (
	"fmt"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
)

type roleService struct {
	repoR repoPort.RoleRepository
}

func NewRoleServiceImpl(r repoPort.RoleRepository) *roleService {
	return &roleService{repoR: r}
}

func (s *roleService) GetRole() (*[]entities.Role, error) {
	roles, err := s.repoR.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find course status: %w", err)
	}

	return roles, nil
}
