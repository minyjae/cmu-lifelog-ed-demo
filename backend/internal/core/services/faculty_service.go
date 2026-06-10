package services

import (
	"context"
	"errors"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
	"github.com/redis/rueidis"
)

type facultyService struct {
	repoFaculty repoPort.FacultyRepository
	cache       *utils.Cache
}

const (
	cacheKeyFacultyAll    = "faculty:all"
	cacheKeyFacultyNameTH = "faculty:nameTH"
)

func NewFacultyServiceImpl(r repoPort.FacultyRepository, redis rueidis.Client) *facultyService {
	return &facultyService{repoFaculty: r, cache: utils.NewCache(redis)}
}

func (s *facultyService) GetAllFaculty() (*[]entities.Faculty, error) {
	return utils.GetOrLoad(context.Background(), s.cache, cacheKeyFacultyAll, s.repoFaculty.FindAll)
}

func (s *facultyService) CheckFacultyExist(facultyName string) (*entities.Faculty, error) {
	key := cacheKeyFacultyNameTH + ":" + facultyName
	faculty, err := utils.GetOrLoad(context.Background(), s.cache, key, func() (*entities.Faculty, error) {
		return s.repoFaculty.FindByNameTH(facultyName)
	})
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
