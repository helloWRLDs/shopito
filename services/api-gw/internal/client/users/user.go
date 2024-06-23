package userclient

import (
	protouser "shopito/pkg/protobuf/user"
	"shopito/services/api-gw/config"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	Close()
}

type UserClient struct {
	GRPC protouser.UserServiceClient
	conn *grpc.ClientConn
}

func New() *UserClient {
	conn, err := grpc.NewClient(config.USERS_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	logrus.Info("user service conn established")
	client := protouser.NewUserServiceClient(conn)
	return &UserClient{
		GRPC: client,
		conn: conn,
	}
}

func (s *UserClient) Close() {
	err := s.conn.Close()
	if err != nil {
		logrus.WithError(err).Error("Couldn't close users service grpc connection")
	} else {
		logrus.Info("users service grpc conn closed")
	}
}
