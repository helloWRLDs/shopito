package userservice

import (
	"context"
	"shopito/services/api-gw/config"
	"shopito/services/api-gw/protobuf"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service interface {
	GetUserService(id int64) (*protobuf.GetUserResponse, error)
	ListUsersService() (*protobuf.GetUsersResponse, error)
	CreateUserService(user *protobuf.CreateUserRequest) (int64, error)
	Close()
	// UpdateUserService()
	DeleteUserService(id int64) error
}

type UserService struct {
	clientGRPC protobuf.UserServiceClient
	conn       *grpc.ClientConn
}

func New() *UserService {
	conn, err := grpc.NewClient(config.USERS_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	client := protobuf.NewUserServiceClient(conn)
	return &UserService{
		clientGRPC: client,
		conn:       conn,
	}
}

func (s *UserService) DeleteUserService(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.clientGRPC.DeleteUser(ctx, &protobuf.DeleteUserRequest{Id: id})
	return err
}

func (s *UserService) CreateUserService(user *protobuf.CreateUserRequest) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.clientGRPC.CreateUser(ctx, user)
	return r.GetId(), err
}

func (s *UserService) GetUserService(id int64) (*protobuf.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.clientGRPC.GetUser(ctx, &protobuf.GetUserRequest{Id: id})
	return r, err
}

func (s *UserService) ListUsersService() (*protobuf.GetUsersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.clientGRPC.GetUsers(ctx, &protobuf.GetUsersRequest{})
	return r, err
}

func (s *UserService) Close() {
	err := s.conn.Close()
	if err != nil {
		logrus.WithError(err).Error("Couldn't close users service grpc connection")
	} else {
		logrus.Info("users service grpc conn closed")
	}
}
