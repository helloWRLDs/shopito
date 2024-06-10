package service

import (
	"shopito/services/notifier/internal/adapter"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	SendEmailService(to, subject, body string) error
	SendAllEmailService(subject, body string, emails []string) error
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

func (s *NotifierService) SendAllEmailService(subject, body string, emails []string) error {
	for _, email := range emails {
		if err := s.adapter.SendEmail(email, subject, body); err != nil {
			logrus.WithField("email", email).Warn("email is not valid(might be deleted from db)")
		}
	}
	return nil
}
