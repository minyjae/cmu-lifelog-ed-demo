package services

import (
	"errors"

	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
)

type requireRoleService struct {
	usersRepo repoPort.UsersRepository
}

func NewRequireRoleService(usersRepo repoPort.UsersRepository) *requireRoleService {
	return &requireRoleService{
		usersRepo: usersRepo,
	}
}

func (s *requireRoleService) GetRoleByEmail(email string) (string, error) {
	user, err := s.usersRepo.FindEmail(email)
	if err != nil {
		return "", errors.New("failed to find user by email: " + email)
	}

	switch user.Role {
	case "admin":
		return "admin", nil
	case "staff":
		return "staff", nil
	case "LE":
		return "LE", nil
	case "officer":
		return "officer", nil
	default:
		return "user", nil
	}
}
