package usercontroller

import (
	"fmt"
	"net/http"
	"shopito/pkg/types/errors"
	grpcutil "shopito/pkg/util/grpc"
	jsonutil "shopito/pkg/util/json"
	userservice "shopito/services/api-gw/internal/service/users"
	"shopito/services/api-gw/protobuf"
	"strconv"

	"github.com/go-chi/chi"
)

type UserController struct {
	service userservice.Service
}

func New(service *userservice.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) DeleteUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("couldn't process the id"))
		return
	}
	err = c.service.DeleteUserService(id)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		jsonutil.EncodeJson(w, status, msg)
		return
	}
	jsonutil.EncodeJson(w, 200, fmt.Sprintf("deleted user with id=%v", id))
}

func (c *UserController) GetUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("couldn't process the id"))
		return
	}
	user, err := c.service.GetUserService(id)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		jsonutil.EncodeJson(w, status, msg)
		return
	}
	jsonutil.EncodeJson(w, 200, user)
}

func (c *UserController) ListUsersController(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.ListUsersService()
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		jsonutil.EncodeJson(w, status, msg)
		return
	}
	jsonutil.EncodeJson(w, 200, users.GetUsers())
}

func (c *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	user, err := jsonutil.DecodeJson[protobuf.CreateUserRequest](r)
	if err != nil {
		jsonutil.EncodeJson(w, 400, "Bad Gateway")
		return
	}

	id, err := c.service.CreateUserService(&user)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		jsonutil.EncodeJson(w, status, msg)
		return
	}
	jsonutil.EncodeJson(w, 201, fmt.Sprintf("Inserted user with id=%v", id))
}
