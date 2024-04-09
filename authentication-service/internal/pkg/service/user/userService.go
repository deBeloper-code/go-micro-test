package user

import (
	"github.com/deBeloper-code/authentication/internal/pkg/entity"
	"github.com/deBeloper-code/authentication/internal/pkg/ports"
)

type service struct {
	repo ports.UserRepository
}

func NewService(repo ports.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllUsers() ([]entity.User, error) {
	var usersFormat []entity.User
	err := s.repo.GetAll(&usersFormat)
	if err != nil {
		return nil, err
	}
	return usersFormat, nil
}
func (s *service) GetUserByEmail(email string) (entity.User, error) {
	var usersFormat entity.User
	err := s.repo.GetByEmail(&usersFormat, email)
	if err != nil {
		return usersFormat, err
	}
	return usersFormat, nil
}

func (s *service) GetUserById(id int) (entity.User, error) {
	var usersFormat entity.User
	err := s.repo.GetById(&usersFormat, id)
	if err != nil {
		return usersFormat, err
	}
	return usersFormat, nil
}

func (s *service) UpdateUserInfo(id int, updates interface{}) (entity.User, error) {
	var usersFormat entity.User
	err := s.repo.UpdateInfo(&usersFormat, updates, id)
	if err != nil {
		return usersFormat, err
	}
	return usersFormat, nil
}
