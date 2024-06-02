package repository

import (
	userproto "shopito/pkg/protobuf/users"
)

type Repository interface {
	GetById(id int64) (*userproto.User, error)
	Insert(user *userproto.User) (int64, error)
	ExistById(id int64) bool
	IsVerified(id int64) bool
	Update(id int64, user *userproto.User) error
	Delete(id int64) error
	GetByEmail(email string) (*userproto.User, error)
	GetAll() ([]*userproto.User, error)
	ExistByEmail(email string) bool
}
