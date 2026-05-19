package repositories

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type facultyRepositoryImpl struct {
	db *gorm.DB
}

func NewFacultyRepository(db *gorm.DB) *facultyRepositoryImpl {
	return &facultyRepositoryImpl{db: db}
}

func (r *facultyRepositoryImpl) FindAll() (*[]entities.Faculty, error) {
	var faculties []entities.Faculty
	if err := r.db.Find(&faculties).Error; err != nil {
		return nil, err
	}
	return &faculties, nil
}

func (r *facultyRepositoryImpl) FindByNameTH(name string) (*entities.Faculty, error) {
	var faculty entities.Faculty
	if err := r.db.Where("name_th = ?", name).First(&faculty).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Faculty not found
		}
		return nil, err // Other error
	}
	return &faculty, nil
}

func (r *facultyRepositoryImpl) FindByArg(arg string) (*entities.Faculty, error) {
	var faculty entities.Faculty
	if err := r.db.Where("code = ? OR name_th = ? OR name_en = ?", arg, arg, arg).First(&faculty).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Faculty not found
		}
		return nil, err // Other error
	}
	return &faculty, nil
}
