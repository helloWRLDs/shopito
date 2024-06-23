package authservice

import (
	"context"
	protouser "shopito/pkg/protobuf/user"
	jwtutil "shopito/pkg/util/jwt"
	userclient "shopito/services/api-gw/internal/client/users"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	RegisterUserService(user *protouser.CreateUserRequest) (int64, error)
	LoginUserService(user *protouser.CreateUserRequest) (*string, error)
}

type AuthService struct {
	client *userclient.UserClient
}

func New(client *userclient.UserClient) *AuthService {
	return &AuthService{
		client: client,
	}
}

func (s *AuthService) RegisterUserService(user *protouser.CreateUserRequest) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := s.client.GRPC.CreateUser(ctx, user)
	if err != nil {
		return -1, err
	}
	return response.GetId(), nil
}

func (s *AuthService) LoginUserService(user *protouser.CreateUserRequest) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser, err := s.client.GRPC.GetUserByEmail(ctx, &protouser.GetUserByEmailRequest{Email: user.GetEmail()})
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
