package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	ADDR          string
	USERS_ADDR    string
	PRODUCTS_ADDR string
	DB            *Db
	JWT           *Jwt
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Panic("config load err")
	}
	ADDR = os.Getenv("GW_ADDR")
	// USERS_ADDR = os.Getenv("USERS_ADDR")
	USERS_ADDR = os.Getenv("USERS_ADDR")
	PRODUCTS_ADDR = os.Getenv("PRODUCTS_ADDR")
	JWT = &Jwt{
		SECRET: os.Getenv("JWT_SECRET"),
	}
	DB = &Db{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		NAME:     os.Getenv("DB_NAME"),
	}
}

type (
	Jwt struct {
		SECRET string
	}
	Db struct {
		HOST     string
		PORT     string
		USER     string
		PASSWORD string
		NAME     string
	}
)
