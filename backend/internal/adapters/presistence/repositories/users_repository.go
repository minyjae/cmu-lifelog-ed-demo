package repositories

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type usersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *usersRepositoryImpl {
	return &usersRepositoryImpl{db: db}
}

func (s *usersRepositoryImpl) Create(role string, email string) (*entities.Users, error) {
	user := models.Users{
		Role:  role,
		Email: email,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	result := user.ToEntity()

	return result, nil
}

func (s *usersRepositoryImpl) Register(user *entities.Users) (*entities.Users, error) {
	m := models.Users{}
	m.FromEntity(user)

	if err := s.db.Create(&m).Error; err != nil {
		return nil, err
	}

	return m.ToEntity(), nil
}

func (s *usersRepositoryImpl) FindUserIsStaff() (*[]entities.Users, error) {
	user := []models.Users{}
	roles := []string{"staff", "admin"}
	if err := s.db.Where("role IN ?", roles).Find(&user).Error; err != nil {
		return nil, err
	}

	result := make([]entities.Users, len(user))
	for i, m := range user {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *usersRepositoryImpl) FindAllUsers() (*[]entities.Users, error) {
	user := []models.Users{}
	if err := s.db.Find(&user).Error; err != nil {
		return nil, err
	}

	result := make([]entities.Users, len(user))
	for i, m := range user {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *usersRepositoryImpl) FindEmail(email string) (*entities.Users, error) {
	user := models.Users{}

	if err := s.db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}

	result := user.ToEntity()
	return result, nil
}
func (s *usersRepositoryImpl) FindUserByID(userID uint) (*entities.Users, error) {
	user := models.Users{}

	if err := s.db.Where("id = ?", userID).
		First(&user).Error; err != nil {
		return nil, err
	}

	result := user.ToEntity()
	return result, nil
}

func (s *usersRepositoryImpl) FindUserByIDAndDelete(userID uint) error {
	blocked := []uint{1}
	user := models.Users{}

	if err := s.db.Where("id NOT IN ?", blocked).Where("id =?", userID).
		Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *usersRepositoryImpl) Save(u *entities.Users) error {
	user := models.Users{}
	user.FromEntity(u)

	if err := s.db.Save(&user).Error; err != nil {
		return nil
	}

	return nil
}

func (s *usersRepositoryImpl) UpdateInfo(email string, user *entities.Users) error {
	u := models.Users{}
	u.FromEntity(user)

	if err := s.db.Model(&u).Where("email = ?", email).Updates(u).Error; err != nil {
		return err
	}

	return nil
}
