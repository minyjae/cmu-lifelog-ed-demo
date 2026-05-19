package repositories

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type UsersRepository interface {
	Create(string, string) (*entities.Users, error)
	Register(user *entities.Users) (*entities.Users, error)

	FindUserIsStaff() (*[]entities.Users, error)
	FindEmail(email string) (*entities.Users, error)
	FindAllUsers() (*[]entities.Users, error)

	FindUserByID(userID uint) (*entities.Users, error)
	Save(u *entities.Users) error

	UpdateInfo(email string, user *entities.Users) error

	FindUserByIDAndDelete(userID uint) error
}
