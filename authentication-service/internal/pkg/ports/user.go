package ports

import "github.com/deBeloper-code/authentication/internal/pkg/entity"

type UserRepository interface {
	GetAll(dest interface{}, conds ...interface{}) error
	GetByEmail(dest interface{}, email string) error
	GetById(dest interface{}, id int) error
	UpdateInfo(dest interface{}, newUpdate interface{}, id int) error
}
type UserService interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserById(id int) (entity.User, error)
	UpdateUserInfo(id int, updates interface{}) (entity.User, error)
}
