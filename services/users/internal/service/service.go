package service

import (
	"context"
	"database/sql"
	userproto "shopito/pkg/protobuf/users"
	"shopito/services/users/internal/repository"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	GetUserService(id int64) (*userproto.User, error)
	GetUserByEmailService(email string) (*userproto.User, error)
	InsertUserService(user *userproto.User) (int64, error)
	GetUsersService() ([]*userproto.User, error)
	DeleteUserService(id int64) error
	UpdateUserService(id int64, user *userproto.User) error
}

type UserService struct {
	repo *repository.Queries
}

func New(repo *repository.Queries) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) DeleteUserService(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	exists, err := s.repo.IsExistByID(ctx, id)
	if err != nil || !exists {
		return status.Errorf(codes.Internal, "Not found")
	}

	if err := s.repo.DeleteUser(ctx, id); err != nil {
		return status.Errorf(codes.Internal, "Internal Error")
	}
	return nil
}

func (s *UserService) GetUsersService() ([]*userproto.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	usersProto := make([]*userproto.User, len(users))
	for i, u := range users {
		usersProto[i] = &userproto.User{
			Id: u.ID,
		}
	}
	return usersProto, nil
}

func (s *UserService) GetUserByEmailService(email string) (*userproto.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	exist, err := s.repo.IsExistByEmail(ctx, email)
	if !exist {
		return nil, status.Errorf(codes.NotFound, "user not found")
	} else if err != nil {
		logrus.WithError(err).Error("Internal Error")
		return nil, status.Errorf(codes.Internal, "Something went wrong")
	}
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	userProto := userproto.User{
		Id:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		IsAdmin:    user.IsAdmin.Bool,
		IsVerified: user.IsVerified.Bool,
	}
	return &userProto, nil
}

func (s *UserService) GetUserService(id int64) (*userproto.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	exist, err := s.repo.IsExistByID(ctx, id)
	if !exist {
		return nil, status.Errorf(codes.NotFound, "User with such id not found")
	} else if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		logrus.WithError(err).Error("Internal Error")
		return nil, status.Errorf(codes.Internal, "Something went wrong")
	}
	protoUser := userproto.User{
		Id:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		IsAdmin:    user.IsAdmin.Bool,
		IsVerified: user.IsVerified.Bool,
	}
	return &protoUser, nil
}

func (s *UserService) InsertUserService(user *userproto.User) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	newUser := repository.CreateUserParams{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
	// if err := user.IsValid(); err != nil {
	// 	return -1, status.Errorf(codes.InvalidArgument, err.Error())
	// }
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)
	if err != nil {
		return -1, status.Errorf(codes.InvalidArgument, "Password is too long")
	}
	newUser.Password = string(hashedPassword)
	exist, err := s.repo.IsExistByEmail(ctx, newUser.Email)
	if !exist {
		return -1, status.Errorf(codes.AlreadyExists, "User with such email already exists")
	} else if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Server Error")
	}
	id, err := s.repo.CreateUser(ctx, newUser)
	if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Error")
	}
	return id, nil
}

func (s *UserService) UpdateUserService(id int64, user *userproto.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	updatedUser := repository.UpdateUserParams{
		ID:         id,
		Name:       user.GetName(),
		Email:      user.GetEmail(),
		Password:   user.GetPassword(),
		IsAdmin:    sql.NullBool{Bool: user.GetIsAdmin()},
		IsVerified: sql.NullBool{Bool: user.GetIsVerified()},
	}

	exist, err := s.repo.IsExistByID(ctx, id)
	if !exist {
		return status.Errorf(codes.AlreadyExists, "User with such email already exists")
	} else if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), 12)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Password is too long")
	}
	updatedUser.Password = string(newPassword)

	if err := s.repo.UpdateUser(ctx, updatedUser); err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return status.Errorf(codes.Internal, "Internal Error")
	}
	return nil
}
