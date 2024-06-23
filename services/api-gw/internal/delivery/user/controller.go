package usercontroller

import (
	"fmt"
	"net/http"
	protouser "shopito/pkg/protobuf/user"
	"shopito/pkg/types/errors"
	"shopito/pkg/types/response"
	grpcutil "shopito/pkg/util/grpc"
	jsonutil "shopito/pkg/util/json"
	userservice "shopito/services/api-gw/internal/service/users"
	"strconv"

	"github.com/go-chi/chi"
)

type UserController struct {
	service userservice.Service
}

func New(service *userservice.UserService) *UserController {
	return &UserController{service: service}
}

// @Summary 	Delete User
// @Tags 		Users
// @Description Delete user by id
// @Accept		json
// @Produce 	json
// @Param 		id	path	int  true  "User ID"
// @Success     200 {object} response.JsonMessage "OK"
// @Failure     422 {object} errors.HTTPError "Unprocessable entity"
// @Failure 	404 {object} errors.HTTPError "Not Found"
// @Failure     500 {object} errors.HTTPError "Internal server error"
// @Router /users/{id} [delete]
func (c *UserController) DeleteUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("couldn't process the id"))
		return
	}
	err = c.service.DeleteUserService(id)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		errors.SendErr(w, errors.GetHTTPErrorByCode(status, msg))
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("Deleted user with id=%v", id)))
}

// @Summary 		Get User
// @Tags Users
// @Description 	Get user by id
// @Accept json
// @Produce json
// @Param 			id				path		int  										true  		"User ID"
// @Success     	200 			{object} 	response.JsonMessage "OK"
// @Failure     	422 			{object} 	errors.HTTPError "Unprocessable entity"
// @Failure 		404 			{object} 	errors.HTTPError "Not Found"
// @Failure     	500 			{object} 	errors.HTTPError "Internal server error"
// @Router 			/users/{id} 	[get]
func (c *UserController) GetUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("couldn't process the id"))
		return
	}
	user, err := c.service.GetUserService(id)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		errors.SendErr(w, errors.GetHTTPErrorByCode(status, msg))
		return
	}
	jsonutil.EncodeJson(w, 200, user)
}

// @Summary 	List Users
// @Tags 		Users
// @Description List Users with parameters
// @Accept		json
// @Produce 	json
// @Success     200 							{array}		protouser.User "OK"
// @Failure     422 							{object} 	errors.HTTPError "Unprocessable entity"
// @Failure 	404 							{object} 	errors.HTTPError "Not Found"
// @Failure     500 							{object} 	errors.HTTPError "Internal server error"
// @Router 		/users 							[get]
func (c *UserController) ListUsersController(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.ListUsersService()
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		errors.SendErr(w, errors.GetHTTPErrorByCode(status, msg))
		return
	}
	jsonutil.EncodeJson(w, 200, users.GetUsers())
}

// @Summary 	Create User
// @Tags 		Users
// @Description Create New User
// @Accept		json
// @Produce 	json
// @Param 		New User	body	protouser.CreateUserRequest		true  		"New User Body"
// @Success     201 {object} response.JsonMessage "Created"
// @Failure     422 {object} errors.HTTPError "Unprocessable entity"
// @Failure 	404 {object} errors.HTTPError "Not Found"
// @Failure     500 {object} errors.HTTPError "Internal server error"
// @Router /users [post]
func (c *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	user, err := jsonutil.DecodeJson[protouser.CreateUserRequest](r)
	if err != nil {
		errors.SendErr(w, errors.ErrBadRequest.SetMessage("couldn't process request body"))
		return
	}

	id, err := c.service.CreateUserService(&user)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		errors.SendErr(w, errors.GetHTTPErrorByCode(status, msg))
		return
	}
	jsonutil.EncodeJson(w, 201, response.NewJsonMessage(201, fmt.Sprintf("Inserted user with id=%v", id)))
}

// @Summary 		Update User
// @Tags 			Users
// @Description 	Update user by id
// @Accept 			json
// @Produce 		json
// @Param 			id				path		int  										true  		"User ID"
// @Param 			User			body		protouser.User								true  		"Updated user body"
// @Success     	200 			{object} 	response.JsonMessage "OK"
// @Failure     	422 			{object} 	errors.HTTPError "Unprocessable entity"
// @Failure 		404 			{object} 	errors.HTTPError "Not Found"
// @Failure     	500 			{object} 	errors.HTTPError "Internal server error"
// @Router 			/users/{id} 	[put]
func (c *UserController) UpdateUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("couldn't process the id"))
		return
	}
	user, err := jsonutil.DecodeJson[protouser.User](r)
	if err != nil {
		errors.SendErr(w, errors.ErrBadRequest.SetMessage("couldn't process the body"))
		return
	}
	if err := c.service.UpdateUserService(id, &user); err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		errors.SendErr(w, errors.GetHTTPErrorByCode(status, msg))
		return
	}
	jsonutil.EncodeJson(w, 201, response.NewJsonMessage(201, fmt.Sprintf("updated user with id=%v", id)))
}
