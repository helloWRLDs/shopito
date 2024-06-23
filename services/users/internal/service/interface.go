package service

import (
	"context"
	protouser "shopito/pkg/protobuf/user"
)

type Service interface {
	DeleteUserService(ctx context.Context, id int64) error
	GetUserByIDService(ctx context.Context, id int64) (*protouser.User, error)
	GetUserByEmailService(ctx context.Context, email string) (*protouser.User, error)
	CreateUserService(ctx context.Context, name, email, password string) (int64, error)
	UpdateUserService(ctx context.Context, id int64, user *protouser.User) error
	ListUserService(ctx context.Context) ([]*protouser.User, error)
}
