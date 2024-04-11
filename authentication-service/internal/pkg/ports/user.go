package ports

import "github.com/deBeloper-code/authentication/internal/pkg/entity"

type UserRepository interface {
	GetAll(dest interface{}, conds ...interface{}) error
	GetByEmail(dest interface{}, email string) error
	GetById(dest interface{}, id int) error
	UpdateInfo(dest interface{}, newUpdate interface{}, id int) error
	Create(dest interface{}) error
	Delete(dest interface{}, id int) error
	ResetPassword(dest interface{}, id int, newPassword string) error
	Login(dest interface{}, email string) error
}
type UserService interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserById(id int) (entity.User, error)
	UpdateUserInfo(id int, updates interface{}) (entity.User, error)
	CreateUser(user entity.User) error
	DeleteUserById(id int) error
	ResetPasswordUser(id int, newPassword string, currentPassword string) error
	LoginUser(email, password string) (entity.User, error)
}
