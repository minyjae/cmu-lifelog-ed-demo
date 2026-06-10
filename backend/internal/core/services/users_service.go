package services

import (
	"context"
	"errors"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

const (
	cacheKeyUsersPrefix = "users:"
	cacheKeyUsersStaff  = "users:staff"
	cacheKeyUsersAll    = "users:all"
)

type usersService struct {
	repoUS repoPort.UsersRepository
	repoF  repoPort.FacultyRepository
	cache  *utils.Cache
}

func NewUsersServiceImpl(r repoPort.UsersRepository, f repoPort.FacultyRepository, redis rueidis.Client) *usersService {
	return &usersService{repoUS: r, repoF: f, cache: utils.NewCache(redis)}
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
	created, err := s.repoUS.Register(user)
	if err != nil {
		return nil, err
	}

	s.cache.InvalidatePrefix(context.Background(), cacheKeyUsersPrefix)
	return created, nil
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

	s.cache.InvalidatePrefix(context.Background(), cacheKeyUsersPrefix)
	return user, nil
}

func (s *usersService) GetStaff() (*[]entities.Users, error) {
	return utils.GetOrLoad(context.Background(), s.cache, cacheKeyUsersStaff, s.repoUS.FindUserIsStaff)
}

func (s *usersService) GetAllUsers() (*[]entities.Users, error) {
	return utils.GetOrLoad(context.Background(), s.cache, cacheKeyUsersAll, s.repoUS.FindAllUsers)
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

	ctx := context.Background()
	s.cache.InvalidatePrefix(ctx, cacheKeyUsersPrefix) // cache ของตัวเอง
	s.cache.InvalidatePrefix(ctx, cacheKeyListPrefix)  // list_queue preload Staff (Users) ไว้
	return nil
}

func (s *usersService) RemoveUser(userID uint) error {
	err := s.repoUS.FindUserByIDAndDelete(userID)
	if err != nil {
		return err
	}

	ctx := context.Background()
	s.cache.InvalidatePrefix(ctx, cacheKeyUsersPrefix) // cache ของตัวเอง
	s.cache.InvalidatePrefix(ctx, cacheKeyListPrefix)  // list_queue preload Staff (Users) ไว้
	return nil
}
