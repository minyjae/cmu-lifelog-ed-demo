package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type FacultyService interface {
	GetAllFaculty() (*[]entities.Faculty, error)
	CheckFacultyExist(string) (*entities.Faculty, error)
	// CreateNewFaculty(faculty *entities.Faculty) (*entities.Faculty, error)
	// DeleteFaculty(facultyCode string) error
}
