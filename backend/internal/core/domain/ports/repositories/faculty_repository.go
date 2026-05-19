package repositories

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type FacultyRepository interface {
	// Create(faculty *entities.Faculty) (*entities.Faculty, error)
	FindAll() (*[]entities.Faculty, error)
	FindByNameTH(string) (*entities.Faculty, error)
	FindByArg(string) (*entities.Faculty, error)
	// Save(*entities.Faculty) error
}
