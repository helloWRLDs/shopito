package service

import (
	"context"
	"shopito/services/notifier/config"
	"shopito/services/notifier/internal/adapter"
	"shopito/services/notifier/protobuf"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Service interface {
	SendEmailService(to, subject, body string) error
	SendAllEmailService(subject, body string) error
}

type NotifierService struct {
	adapter *adapter.EmailCredentials
}

func New(adapter *adapter.EmailCredentials) *NotifierService {
	return &NotifierService{
		adapter: adapter,
	}
}

func (s *NotifierService) SendEmailService(to, subject, body string) error {
	if err := s.adapter.SendEmail(to, subject, body); err != nil {
		return status.Errorf(codes.Unavailable, "Specified Email is not available")
	}
	return nil
}

func (s *NotifierService) SendAllEmailService(subject, body string) error {
	conn, err := grpc.NewClient(config.USERS_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return status.Errorf(codes.Unavailable, "Users service is unavailable")
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := protobuf.NewUserServiceClient(conn)
	r, err := c.GetUsers(ctx, &protobuf.GetUsersRequest{})
	if err != nil {
		return status.Errorf(codes.Unavailable, "Users service is unavailable")
	}
	users := r.Users
	for _, user := range users {
		if err := s.adapter.SendEmail(user.Email, subject, body); err != nil {
			logrus.WithField("email", user.Email).Warn("email is not valid(might be deleted from db)")
		}
	}
	return nil
}
