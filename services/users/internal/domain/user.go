package domain

import (
	"errors"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password,omitempty"`
	IsAdmin    bool      `json:"is_admin"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func New(name, email, password string, isAdmin, isVerified bool) *User {
	return &User{
		Name:       name,
		Email:      email,
		Password:   password,
		IsAdmin:    isAdmin,
		IsVerified: isVerified,
	}
}

func (u *User) IsValid() error {
	if len(u.Name) <= 0 {
		return errors.New("name required")
	}
	if len(u.Email) <= 0 {
		return errors.New("email required")
	}
	if len(u.Password) < 5 {
		return errors.New("password is too short")
	}
	return nil
}
