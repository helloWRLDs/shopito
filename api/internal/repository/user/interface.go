package userrepository

import userdomain "shopito/api/internal/domain/user"

type UserRepository interface {
	GetById(id int) (*userdomain.User, error)
	Insert(user *userdomain.User) (int, error)
	ExistById(id int) bool
	IsVerified(id int) bool
	Update(id int, user *userdomain.User) error
	Delete(id int) error
	GetByEmail(email string) (*userdomain.User, error)
	GetAll() (*[]userdomain.User, error)
	ExistByEmail(email string) bool
}
