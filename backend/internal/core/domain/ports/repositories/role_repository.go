package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type RoleRepository interface {
	FindAll() (*[]entities.Role, error)
}
