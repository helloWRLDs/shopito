package service

import (
	"shopito/services/users/internal/repository"
	"shopito/services/users/protobuf"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	GetUserService(id int64) (*protobuf.User, error)
	GetUserByEmailService(email string) (*protobuf.User, error)
	InsertUserService(user *protobuf.User) (int64, error)
	GetUsersService() ([]*protobuf.User, error)
	DeleteUserService(id int64) error
}

type UserService struct {
	repo repository.Repository
}

func New(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) DeleteUserService(id int64) error {
	if !s.repo.ExistById(id) {
		return status.Errorf(codes.Internal, "Not found")
	}
	if err := s.repo.Delete(id); err != nil {
		return status.Errorf(codes.Internal, "Internal Error")
	}
	return nil
}

func (s *UserService) GetUsersService() ([]*protobuf.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}
	return users, nil
}

func (s *UserService) GetUserByEmailService(email string) (*protobuf.User, error) {
	if !s.repo.ExistByEmail(email) {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	return user, nil
}

func (s *UserService) GetUserService(id int64) (*protobuf.User, error) {
	if !s.repo.ExistById(id) {
		return nil, status.Errorf(codes.NotFound, "User with such id not found")
	}
	user, err := s.repo.GetById(id)
	if err != nil {
		logrus.WithError(err).Error("Internal Error")
		return nil, status.Errorf(codes.Internal, "Something went wrong")
	}
	return user, nil
}

func (s *UserService) InsertUserService(user *protobuf.User) (int64, error) {
	// if err := user.IsValid(); err != nil {
	// 	return -1, status.Errorf(codes.InvalidArgument, err.Error())
	// }
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return -1, status.Errorf(codes.InvalidArgument, "Password is too long")
	}
	user.Password = string(hashedPassword)
	if s.repo.ExistByEmail(user.Email) {
		return -1, status.Errorf(codes.AlreadyExists, "User with such email already exists")
	}
	id, err := s.repo.Insert(user)
	if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Error")
	}
	return id, nil
}
