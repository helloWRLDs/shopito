package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ADDR  string
	DB    *Db
	JWT   *Jwt
	EMAIL *Email
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	ADDR = os.Getenv("ADDR")
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
	EMAIL = &Email{
		HOST:     os.Getenv("EMAIL_HOST"),
		PORT:     os.Getenv("EMAIL_PORT"),
		USERNAME: os.Getenv("EMAIL_USERNAME"),
		PASSWORD: os.Getenv("EMAIL_PASSWORD"),
		FROM:     os.Getenv("EMAIL_FROM"),
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
	Email struct {
		HOST     string
		PORT     string
		USERNAME string
		PASSWORD string
		FROM     string
	}
)
