package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	ADDR string
	DB   *Db
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Error("err parsing config")
		panic(err)
	}
	ADDR = os.Getenv("USERS_ADDR")
	DB = &Db{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		NAME:     os.Getenv("DB_NAME"),
	}
}

type Db struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	NAME     string
}
