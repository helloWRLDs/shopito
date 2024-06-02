package delivery

import (
	"context"
	"shopito/services/users/internal/service"
	"shopito/services/users/protobuf"
)

type Delivery struct {
	protobuf.UnimplementedUserServiceServer
	serv service.Service
}

func New(serv *service.UserService) *Delivery {
	return &Delivery{
		serv: serv,
	}
}

func (d *Delivery) CreateUser(ctx context.Context, request *protobuf.CreateUserRequest) (*protobuf.CreateUserResponse, error) {
	user := protobuf.User{
		Name:     request.GetName(),
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}
	response := protobuf.CreateUserResponse{}

	id, err := d.serv.InsertUserService(&user)
	if err != nil {
		response.Success = false
		return &response, err
	}
	response.Success = true
	response.Id = id
	return &response, nil
}

func (d *Delivery) DeleteUser(ctx context.Context, request *protobuf.DeleteUserRequest) (*protobuf.DeleteUserResponse, error) {
	response := &protobuf.DeleteUserResponse{}
	if err := d.serv.DeleteUserService(request.GetId()); err != nil {
		response.Success = false
		return response, err
	}
	response.Success = true
	return response, nil
}

func (d *Delivery) GetUserByEmail(ctx context.Context, request *protobuf.GetUserByEmailRequest) (*protobuf.GetUserByEmailResponse, error) {
	user, err := d.serv.GetUserByEmailService(request.GetEmail())
	if err != nil {
		return nil, err
	}
	response := protobuf.GetUserByEmailResponse{
		User: user,
	}
	return &response, nil
}

func (d *Delivery) GetUser(ctx context.Context, request *protobuf.GetUserRequest) (*protobuf.GetUserResponse, error) {
	user, err := d.serv.GetUserService(request.GetId())
	if err != nil {
		return nil, err
	}
	response := &protobuf.GetUserResponse{
		User: user,
	}
	return response, nil
}

func (d *Delivery) GetUsers(ctx context.Context, request *protobuf.GetUsersRequest) (*protobuf.GetUsersResponse, error) {
	users, err := d.serv.GetUsersService()
	if err != nil {
		return nil, err
	}
	response := protobuf.GetUsersResponse{
		Users: users,
	}

	return &response, nil
}

func (d *Delivery) UpdateUser(ctx context.Context, request *protobuf.UpdateUserRequest) (*protobuf.UpdateUserResponse, error) {
	return nil, nil
}
