package userservice

import (
	"context"
	"fmt"
	userproto "shopito/pkg/protobuf/users"
	"shopito/services/api-gw/config"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service interface {
	GetUserService(id int64) (*userproto.User, error)
	ListUsersService() (*userproto.GetUsersResponse, error)
	CreateUserService(user *userproto.CreateUserRequest) (int64, error)
	Close()
	UpdateUserService(id int64, user *userproto.User) error
	DeleteUserService(id int64) error
}

type UserService struct {
	clientGRPC userproto.UserServiceClient
	conn       *grpc.ClientConn
}

func New() *UserService {
	conn, err := grpc.NewClient(config.USERS_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	logrus.Info("user service conn established")
	client := userproto.NewUserServiceClient(conn)
	return &UserService{
		clientGRPC: client,
		conn:       conn,
	}
}

func (s *UserService) DeleteUserService(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.clientGRPC.DeleteUser(ctx, &userproto.DeleteUserRequest{Id: id})
	return err
}

func (s *UserService) CreateUserService(user *userproto.CreateUserRequest) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.clientGRPC.CreateUser(ctx, user)
	return r.GetId(), err
}

func (s *UserService) GetUserService(id int64) (*userproto.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.clientGRPC.GetUser(ctx, &userproto.GetUserRequest{Id: id})
	fmt.Println(r.GetUser().GetEmail())
	return r.GetUser(), err
}

func (s *UserService) ListUsersService() (*userproto.GetUsersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.clientGRPC.GetUsers(ctx, &userproto.GetUsersRequest{})
	return r, err
}

func (s *UserService) UpdateUserService(id int64, user *userproto.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.clientGRPC.UpdateUser(ctx, &userproto.UpdateUserRequest{
		Id:   id,
		User: user,
	})
	return err
}

func (s *UserService) Close() {
	err := s.conn.Close()
	if err != nil {
		logrus.WithError(err).Error("Couldn't close users service grpc connection")
	} else {
		logrus.Info("users service grpc conn closed")
	}
}
