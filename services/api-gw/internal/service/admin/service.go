package adminservice

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
	PromoteUserService(id string) error
	DemoteUserService(id string) error
	Close() error
}

type AdminService struct {
	clientGRPC protobuf.UserServiceClient
	conn       *grpc.ClientConn
}

func New() *AdminService {
	conn, err := grpc.NewClient(config.USERS_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	client := protobuf.NewUserServiceClient(conn)
	return &AdminService{
		clientGRPC: client,
		conn:       conn,
	}
}

func (s *AdminService) PromoteUserService(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rGet, err := s.clientGRPC.GetUser(ctx, &protobuf.GetUserRequest{Id: id})
	if err != nil {
		return err
	}
	_, err = s.clientGRPC.UpdateUser(ctx, &protobuf.UpdateUserRequest{
		Id: id,
		User: &protobuf.User{
			Name:       rGet.GetUser().GetEmail(),
			Email:      rGet.GetUser().GetEmail(),
			Password:   rGet.GetUser().GetPassword(),
			IsAdmin:    true,
			IsVerified: rGet.GetUser().GetIsVerified(),
		},
	})
	return err
}

func (s *AdminService) DemoteUserService(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rGet, err := s.clientGRPC.GetUser(ctx, &protobuf.GetUserRequest{Id: id})
	if err != nil {
		return err
	}
	_, err = s.clientGRPC.UpdateUser(ctx, &protobuf.UpdateUserRequest{
		Id: id,
		User: &protobuf.User{
			Name:       rGet.GetUser().GetName(),
			Email:      rGet.GetUser().GetEmail(),
			Password:   rGet.GetUser().GetPassword(),
			IsAdmin:    false,
			IsVerified: rGet.GetUser().GetIsVerified(),
		},
	})
	return err
}

func (s *AdminService) Close() error {
	return s.conn.Close()
}
