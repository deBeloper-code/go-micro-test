package user

import (
	"github.com/deBeloper-code/authentication/internal/pkg/entity"
	"github.com/deBeloper-code/authentication/internal/pkg/ports"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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
func (s *service) LoginUser(email, password string) (entity.User, error) {
	var usersFormat entity.User
	err := s.repo.Login(&usersFormat, email)
	if err != nil {
		return usersFormat, err
	}

	if err := tryMatchPassword(usersFormat.Password, password); err != nil {
		log.New().Errorf(err.Error())
		return usersFormat, err
	}

	return usersFormat, err
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

func (s *service) ResetPasswordUser(id int, newPassword string, currentPassword string) error {
	var usersFormat entity.User
	// 1. Looking for User
	userFounded, err := s.GetUserById(id)
	if err != nil {
		log.New().Errorf(err.Error())
		return err
	}
	// 2. Trying match password
	if err := tryMatchPassword(userFounded.Password, currentPassword); err != nil {
		log.New().Errorf(err.Error())
		return err
	}
	// 3. Reset password
	errReset := s.repo.ResetPassword(&usersFormat, id, newPassword)
	if errReset != nil {
		log.New().Errorf(errReset.Error())
		return errReset
	}
	return nil
}

func (s *service) DeleteUserById(id int) error {
	var usersFormat entity.User
	err := s.repo.Delete(&usersFormat, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateUser(user entity.User) error {
	// Hash password
	user.Password = hashAndSalt(user.Password)

	err := s.repo.Create(&user)
	if err != nil {
		return err
	}
	return nil
}

func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}
	return string(hash)
}

func tryMatchPassword(userPassword, credentialsPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(credentialsPassword))
	if err != nil {
		return err
	}
	return nil
}
