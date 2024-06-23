package userservice

import (
	"context"
	protouser "shopito/pkg/protobuf/user"
	userclient "shopito/services/api-gw/internal/client/users"
	"time"
)

type Service interface {
	GetUserService(id int64) (*protouser.User, error)
	ListUsersService() (*protouser.ListUsersResponse, error)
	CreateUserService(user *protouser.CreateUserRequest) (int64, error)
	UpdateUserService(id int64, user *protouser.User) error
	DeleteUserService(id int64) error
}

type UserService struct {
	client *userclient.UserClient
}

func New(client *userclient.UserClient) *UserService {
	return &UserService{
		client: client,
	}
}

func (s *UserService) DeleteUserService(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.GRPC.DeleteUser(ctx, &protouser.DeleteUserRequest{Id: id})
	return err
}

func (s *UserService) CreateUserService(user *protouser.CreateUserRequest) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.client.GRPC.CreateUser(ctx, user)
	return r.GetId(), err
}

func (s *UserService) GetUserService(id int64) (*protouser.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.client.GRPC.GetUserByID(ctx, &protouser.GetUserByIDRequest{Id: id})
	return r.GetUser(), err
}

func (s *UserService) ListUsersService() (*protouser.ListUsersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.client.GRPC.ListUsers(ctx, &protouser.ListUsersRequest{})
	return r, err
}

func (s *UserService) UpdateUserService(id int64, user *protouser.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.GRPC.UpdateUser(ctx, &protouser.UpdateUserRequest{
		Id:   id,
		User: user,
	})
	return err
}
