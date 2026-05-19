package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type UsersService interface {
	// @Auth
	Register(user *entities.Users, password string) (*entities.Users, error)
	SignIn(email, password string) (*entities.Users, error)

	// @Create
	CreateUser(string, string) (*entities.Users, error)

	// @Read
	GetStaff() (*[]entities.Users, error)
	GetAllUsers() (*[]entities.Users, error)
	FindEmail(email string) (*entities.Users, error)

	// @Update
	UpdateInfo(string, *entities.Users) error
	// UpdateUserFaculty(string, uint) (*entities.Users, error)

	// @Delete
	RemoveUser(userID uint) error

	// @Extra functions
	IsUserExist(string) (bool, error)
}
