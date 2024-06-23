package service

import (
	"context"
	protouser "shopito/pkg/protobuf/user"
	"shopito/services/users/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	repo *repository.Queries
}

func New(repo *repository.Queries) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) DeleteUserService(ctx context.Context, id int64) error {
	exist, err := s.repo.IsExistByID(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "internal server error")
	}
	if !exist {
		return status.Errorf(codes.NotFound, "user not found")
	}
	if err := s.repo.DeleteUser(ctx, id); err != nil {
		return status.Errorf(codes.Internal, "internal server error")
	}
	return nil
}

func (s *UserService) GetUserByIDService(ctx context.Context, id int64) (*protouser.User, error) {
	exist, err := s.repo.IsExistByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}
	if !exist {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}
	return user.MapToProto(), nil
}

func (s *UserService) GetUserByEmailService(ctx context.Context, email string) (*protouser.User, error) {
	exist, err := s.repo.IsExistByEmail(ctx, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}
	if !exist {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}
	return user.MapToProto(), nil
}

func (s *UserService) CreateUserService(ctx context.Context, name, email, password string) (int64, error) {
	exist, err := s.repo.IsExistByEmail(ctx, email)
	if err != nil {
		return -1, status.Errorf(codes.Internal, "internal server error")
	}
	if exist {
		return -1, status.Errorf(codes.AlreadyExists, "user already exist")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return -1, status.Errorf(codes.InvalidArgument, "password is too long")
	}
	id, err := s.repo.CreateUser(ctx, repository.CreateUserParams{
		Name: name, Email: email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return -1, status.Errorf(codes.Internal, "internal server error")
	}
	return id, nil
}

func (s *UserService) UpdateUserService(ctx context.Context, id int64, user *protouser.User) error {
	exist, err := s.repo.IsExistByID(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "internal server error")
	}
	if !exist {
		return status.Errorf(codes.NotFound, "user not found")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 12)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "password is too long")
	}
	params := repository.UpdateUserParams{
		ID:         id,
		Name:       user.GetName(),
		Email:      user.GetEmail(),
		Password:   string(hashedPassword),
		IsAdmin:    user.GetIsAdmin(),
		IsVerified: user.GetIsVerified(),
	}
	if err := s.repo.UpdateUser(ctx, params); err != nil {
		return err
	}
	return nil
}

func (s *UserService) ListUserService(ctx context.Context) ([]*protouser.User, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}
	protoUsers := make([]*protouser.User, len(users))
	for i, user := range users {
		protoUsers[i] = user.MapToProto()
	}
	return protoUsers, nil
}
