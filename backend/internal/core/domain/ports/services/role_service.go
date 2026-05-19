package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type RoleService interface {
	GetRole() (*[]entities.Role, error)
}
