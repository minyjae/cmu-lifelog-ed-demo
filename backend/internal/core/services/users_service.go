package services

import (
	"errors"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
	"gorm.io/gorm"
)

type usersService struct {
	repoUS repoPort.UsersRepository
	repoF  repoPort.FacultyRepository
}

func NewUsersServiceImpl(r repoPort.UsersRepository, f repoPort.FacultyRepository) *usersService {
	return &usersService{repoUS: r, repoF: f}
}

func (s *usersService) Register(user *entities.Users, password string) (*entities.Users, error) {
	exist, err := s.IsUserExist(user.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("email already exists")
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user.Password = hashed
	user.Role = "user"
	return s.repoUS.Register(user)
}

func (s *usersService) SignIn(email, password string) (*entities.Users, error) {
	user, err := s.repoUS.FindEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if !utils.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

// รับ email มาพร้อม role ในการเพิ่ม user แล้วให้ user นั้นๆเข้าระบบมาจะเก็บข้อมูลเต็มให้
func (s *usersService) CreateUser(role string, email string) (*entities.Users, error) {
	u, _ := s.repoUS.FindEmail(email)
	if u != nil {
		return nil, errors.New("this email already exist")
	}

	user, err := s.repoUS.Create(role, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) GetStaff() (*[]entities.Users, error) {
	u, err := s.repoUS.FindUserIsStaff()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *usersService) GetAllUsers() (*[]entities.Users, error) {
	u, err := s.repoUS.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *usersService) FindEmail(email string) (*entities.Users, error) {
	user, err := s.repoUS.FindEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found") // ไม่เจอ user
		}
		return nil, err // error อื่น
	}
	return user, nil // เจอ user แล้ว
}

func (s *usersService) IsUserExist(email string) (bool, error) {
	_, err := s.repoUS.FindEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // ไม่เจอ user
		}
		return false, err // error อื่น
	}
	return true, nil // เจอ user แล้ว
}

// func (s *usersService) UpdateUserFaculty(code string, id uint) (*entities.Users, error) {
// 	u, err := s.repoUS.FindUserByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	f, err := s.repoF.FindByCode(code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	u.OrganizationCode = f.Code
// 	u.OrganizationNameTH = f.NameTH
// 	u.OrganizationNameEN = f.NameTH

// 	err = s.repoUS.Save(u)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }

func (s *usersService) UpdateInfo(email string, user *entities.Users) error {
	err := s.repoUS.UpdateInfo(email, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *usersService) RemoveUser(userID uint) error {
	err := s.repoUS.FindUserByIDAndDelete(userID)
	if err != nil {
		return err
	}

	return nil
}
