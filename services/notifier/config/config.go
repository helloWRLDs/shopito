package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	ADDR       string
	EMAIL      *Email
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Error("err parsing config")
		panic(err)
	}
	ADDR = os.Getenv("NOTIFIER_ADDR")
	EMAIL = &Email{
		HOST:     os.Getenv("EMAIL_HOST"),
		PORT:     os.Getenv("EMAIL_PORT"),
		USERNAME: os.Getenv("EMAIL_USERNAME"),
		PASSWORD: os.Getenv("EMAIL_PASSWORD"),
		FROM:     os.Getenv("EMAIL_FROM"),
	}
}

type Email struct {
	HOST     string
	PORT     string
	USERNAME string
	PASSWORD string
	FROM     string
}
