package services

import (
	"errors"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
)

type facultyService struct {
	repoFaculty repoPort.FacultyRepository
}

func NewFacultyServiceImpl(r repoPort.FacultyRepository) *facultyService {
	return &facultyService{repoFaculty: r}
}

func (s *facultyService) GetAllFaculty() (*[]entities.Faculty, error) {
	faculties, err := s.repoFaculty.FindAll()
	if err != nil {
		return nil, err
	}

	return faculties, nil
}

func (s *facultyService) CheckFacultyExist(facultyName string) (*entities.Faculty, error) {
	faculty, err := s.repoFaculty.FindByNameTH(facultyName)
	if err != nil {
		return nil, err
	}

	if faculty == nil {
		return nil, errors.New("faculty not found")
	}

	return faculty, nil // Faculty exists
}

// func (s *facultyService) CreateNewFaculty(faculty *entities.Faculty) (*entities.Faculty, error) {
// 	_, err := s.repoFaculty.FindByCode(faculty.Code)
// 	if err == nil {
// 		return nil, err // Faculty already exists
// 	}

// 	createdFaculty, err := s.repoFaculty.Create(faculty)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return createdFaculty, nil
// }

// func (s *facultyService) DeleteFaculty(facultyCode string) error {
// 	faculty, err := s.repoFaculty.FindByCode(facultyCode)
// 	if err != nil {
// 		return err
// 	}
// 	if faculty == nil {
// 		return nil // Faculty does not exist, nothing to delete
// 	}
// 	faculty.IsActive = false

// 	err = s.repoFaculty.Save(faculty)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
