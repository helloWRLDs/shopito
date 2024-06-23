package delivery

import (
	"context"
	protouser "shopito/pkg/protobuf/user"
	"shopito/services/users/internal/service"
)

type Delivery struct {
	protouser.UnimplementedUserServiceServer
	serv service.Service
}

func New(serv *service.UserService) *Delivery {
	return &Delivery{serv: serv}
}

func (d *Delivery) CreateUser(ctx context.Context, request *protouser.CreateUserRequest) (*protouser.CreateUserResponse, error) {
	response := &protouser.CreateUserResponse{Success: true, Id: -1}
	id, err := d.serv.CreateUserService(ctx, request.GetName(), request.GetEmail(), request.GetPassword())
	if err != nil {
		response.Success = false
	}
	response.Id = id
	return response, err
}

func (d *Delivery) DeleteUser(ctx context.Context, request *protouser.DeleteUserRequest) (*protouser.DeleteUserResponse, error) {
	response := &protouser.DeleteUserResponse{Success: true}
	err := d.serv.DeleteUserService(ctx, request.GetId())
	if err != nil {
		response.Success = false
	}
	return response, err
}

func (d *Delivery) GetUserByEmail(ctx context.Context, request *protouser.GetUserByEmailRequest) (*protouser.GetUserByEmailResponse, error) {
	user, err := d.serv.GetUserByEmailService(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}
	return &protouser.GetUserByEmailResponse{User: user}, nil
}

func (d *Delivery) GetUserByID(ctx context.Context, request *protouser.GetUserByIDRequest) (*protouser.GetUserByIDResponse, error) {
	user, err := d.serv.GetUserByIDService(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &protouser.GetUserByIDResponse{User: user}, nil
}

func (d *Delivery) ListUsers(ctx context.Context, request *protouser.ListUsersRequest) (*protouser.ListUsersResponse, error) {
	users, err := d.serv.ListUserService(ctx)
	if err != nil {
		return nil, err
	}
	return &protouser.ListUsersResponse{Users: users}, nil
}

func (d *Delivery) UpdateUser(ctx context.Context, request *protouser.UpdateUserRequest) (*protouser.UpdateUserResponse, error) {
	response := &protouser.UpdateUserResponse{Success: true}
	err := d.serv.UpdateUserService(ctx, request.GetId(), request.GetUser())
	if err != nil {
		response.Success = false
		return response, err
	}
	return response, nil
}
