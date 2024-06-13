package authservice

import (
	"context"
	userproto "shopito/pkg/protobuf/users"
	jwtutil "shopito/pkg/util/jwt"
	"shopito/services/api-gw/config"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Service interface {
	RegisterUserService(user *userproto.CreateUserRequest) (int64, error)
	LoginUserService(user *userproto.CreateUserRequest) (*string, error)
	Close()
}

type AuthService struct {
	clientGRPC userproto.UserServiceClient
	conn       *grpc.ClientConn
}

func New() *AuthService {
	conn, err := grpc.NewClient(config.USERS_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	client := userproto.NewUserServiceClient(conn)
	return &AuthService{
		clientGRPC: client,
		conn:       conn,
	}
}

func (s *AuthService) RegisterUserService(user *userproto.CreateUserRequest) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := s.clientGRPC.CreateUser(ctx, user)
	if err != nil {
		return -1, err
	}
	return response.GetId(), nil
}

func (s *AuthService) LoginUserService(user *userproto.CreateUserRequest) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser, err := s.clientGRPC.GetUserByEmail(ctx, &userproto.GetUserByEmailRequest{Email: user.GetEmail()})
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.GetUser().GetPassword()), []byte(user.GetPassword()))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Wrong password")
	}
	token, err := jwtutil.GenerateToken(int(existingUser.GetUser().GetId()), existingUser.GetUser().GetIsAdmin(), existingUser.GetUser().GetIsVerified())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Incorrect Credentials")
	}
	return token, nil
}

func (s *AuthService) Close() {
	err := s.conn.Close()
	if err != nil {
		logrus.WithError(err).Error("Couldn't close users service grpc connection")
	} else {
		logrus.Info("users service grpc conn closed")
	}
}
