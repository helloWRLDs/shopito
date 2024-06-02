package delivery

import (
	"context"
	"shopito/services/users/internal/service"
	"shopito/pkg/protobuf/users"
)

type Delivery struct {
	userproto.UnimplementedUserServiceServer
	serv service.Service
}

func New(serv *service.UserService) *Delivery {
	return &Delivery{
		serv: serv,
	}
}

func (d *Delivery) CreateUser(ctx context.Context, request *userproto.CreateUserRequest) (*userproto.CreateUserResponse, error) {
	user := userproto.User{
		Name:     request.GetName(),
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}
	response := userproto.CreateUserResponse{}

	id, err := d.serv.InsertUserService(&user)
	if err != nil {
		response.Success = false
		return &response, err
	}
	response.Success = true
	response.Id = id
	return &response, nil
}

func (d *Delivery) DeleteUser(ctx context.Context, request *userproto.DeleteUserRequest) (*userproto.DeleteUserResponse, error) {
	response := &userproto.DeleteUserResponse{}
	if err := d.serv.DeleteUserService(request.GetId()); err != nil {
		response.Success = false
		return response, err
	}
	response.Success = true
	return response, nil
}

func (d *Delivery) GetUserByEmail(ctx context.Context, request *userproto.GetUserByEmailRequest) (*userproto.GetUserByEmailResponse, error) {
	user, err := d.serv.GetUserByEmailService(request.GetEmail())
	if err != nil {
		return nil, err
	}
	response := userproto.GetUserByEmailResponse{
		User: user,
	}
	return &response, nil
}

func (d *Delivery) GetUser(ctx context.Context, request *userproto.GetUserRequest) (*userproto.GetUserResponse, error) {
	user, err := d.serv.GetUserService(request.GetId())
	if err != nil {
		return nil, err
	}
	response := &userproto.GetUserResponse{
		User: user,
	}
	return response, nil
}

func (d *Delivery) GetUsers(ctx context.Context, request *userproto.GetUsersRequest) (*userproto.GetUsersResponse, error) {
	users, err := d.serv.GetUsersService()
	if err != nil {
		return nil, err
	}
	response := userproto.GetUsersResponse{
		Users: users,
	}

	return &response, nil
}

func (d *Delivery) UpdateUser(ctx context.Context, request *userproto.UpdateUserRequest) (*userproto.UpdateUserResponse, error) {
	return nil, nil
}
