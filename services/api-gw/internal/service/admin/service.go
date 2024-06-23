package adminservice

import (
	"context"
	protouser "shopito/pkg/protobuf/user"
	userclient "shopito/services/api-gw/internal/client/users"

	"github.com/pkg/errors"
)

type Service interface {
	PromoteUserService(ctx context.Context, id int64) error
	DemoteUserService(ctx context.Context, id int64) error
}

type AdminService struct {
	userClient *userclient.UserClient
}

func New(userclient *userclient.UserClient) *AdminService {
	return &AdminService{
		userClient: userclient,
	}
}

func (s *AdminService) PromoteUserService(ctx context.Context, id int64) error {
	getResponse, err := s.userClient.GRPC.GetUserByID(ctx, &protouser.GetUserByIDRequest{Id: id})
	if err != nil {
		return err
	}
	updatedUser := getResponse.GetUser()
	if updatedUser.IsAdmin {
		return nil
	}
	updatedUser.IsAdmin = true
	putResponse, err := s.userClient.GRPC.UpdateUser(ctx, &protouser.UpdateUserRequest{Id: id, User: updatedUser})
	if err != nil {
		return err
	}
	if !putResponse.Success {
		return errors.Errorf("something went wrong")
	}
	return nil
}

func (s *AdminService) DemoteUserService(ctx context.Context, id int64) error {
	getResponse, err := s.userClient.GRPC.GetUserByID(ctx, &protouser.GetUserByIDRequest{Id: id})
	if err != nil {
		return err
	}
	updatedUser := getResponse.GetUser()
	if !updatedUser.IsAdmin {
		return nil
	}
	updatedUser.IsAdmin = false
	putResponse, err := s.userClient.GRPC.UpdateUser(ctx, &protouser.UpdateUserRequest{Id: id, User: updatedUser})
	if err != nil {
		return err
	}
	if !putResponse.Success {
		return errors.Errorf("something went wrong")
	}
	return nil
}
