package repository

import (
	"shopito/services/users/protobuf"
)

type Repository interface {
	GetById(id int64) (*protobuf.User, error)
	Insert(user *protobuf.User) (int64, error)
	ExistById(id int64) bool
	IsVerified(id int64) bool
	Update(id int64, user *protobuf.User) error
	Delete(id int64) error
	GetByEmail(email string) (*protobuf.User, error)
	GetAll() ([]*protobuf.User, error)
	ExistByEmail(email string) bool
}
