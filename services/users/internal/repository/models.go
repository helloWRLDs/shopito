// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package repository

import (
	"time"
)

type User struct {
	ID         int64
	Name       string
	Email      string
	Password   string
	IsAdmin    bool
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
