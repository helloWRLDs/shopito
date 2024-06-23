package authcontroller

import (
	"fmt"
	"net/http"
	protouser "shopito/pkg/protobuf/user"
	"shopito/pkg/types/errors"
	"shopito/pkg/types/response"
	grpcutil "shopito/pkg/util/grpc"
	jsonutil "shopito/pkg/util/json"
	authservice "shopito/services/api-gw/internal/service/auth"
)

type AuthController struct {
	service authservice.Service
}

func New(service *authservice.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// @Summary 	Login User
// @Tags 		Auth
// @Description Authenticate and Authorize User
// @Accept		json
// @Produce 	json
// @Param 		email	body	string  true  "email"
// @Param 		password	body	string  true  "password"
// @Success     200 {object} response.JsonMessage "OK"
// @Failure     422 {object} errors.HTTPError "Unprocessable entity"
// @Failure 	404 {object} errors.HTTPError "Not Found"
// @Failure     500 {object} errors.HTTPError "Internal server error"
// @Router /auth/login [post]
func (c *AuthController) LoginController(w http.ResponseWriter, r *http.Request) {

	user, err := jsonutil.DecodeJson[protouser.CreateUserRequest](r)
	if err != nil {
		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("Couldn't process the user data"))
		return
	}
	token, err := c.service.LoginUserService(&user)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		jsonutil.EncodeJson(w, status, msg)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, *token))
}

// @Summary 	Register User
// @Tags 		Auth
// @Description Register New User
// @Accept		json
// @Produce 	json
// @Param 		name	body	string  true  "name"
// @Param 		email	body	string  true  "email"
// @Param 		password	body	string  true  "password"
// @Success     200 {object} response.JsonMessage "OK"
// @Failure     422 {object} errors.HTTPError "Unprocessable entity"
// @Failure 	404 {object} errors.HTTPError "Not Found"
// @Failure     500 {object} errors.HTTPError "Internal server error"
// @Router /auth/register [post]
func (c *AuthController) RegisterController(w http.ResponseWriter, r *http.Request) {
	newUser, err := jsonutil.DecodeJson[protouser.CreateUserRequest](r)
	if err != nil {
		jsonutil.EncodeJson(w, 400, "Invalid Credentails")
		return
	}
	id, err := c.service.RegisterUserService(&newUser)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		jsonutil.EncodeJson(w, status, msg)
		return
	}
	jsonutil.EncodeJson(w, 201, response.NewJsonMessage(201, fmt.Sprintf("user registered with id=%v", id)))
}
