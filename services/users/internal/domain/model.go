package domain

import (
	protouser "shopito/pkg/protobuf/user"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
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

func (u *User) MapToProto() *protouser.User {
	return &protouser.User{
		Id:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		IsAdmin:    u.IsAdmin,
		IsVerified: u.IsVerified,
		CreatedAt:  &timestamppb.Timestamp{Seconds: int64(u.CreatedAt.Second()), Nanos: int32(u.CreatedAt.Nanosecond())},
		UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(u.UpdatedAt.Second()), Nanos: int32(u.UpdatedAt.Nanosecond())},
	}
}
