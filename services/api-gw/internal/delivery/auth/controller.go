package authcontroller

// import (
// 	"net/http"
// 	"shopito/pkg/types/errors"
// 	grpcutil "shopito/pkg/util/grpc"
// 	jsonutil "shopito/pkg/util/json"
// 	authservice "shopito/services/api-gw/internal/service/auth"
// 	"shopito/pkg/services/api-gw/protobuf"

// 	"github.com/go-chi/chi"
// )

// type AuthController struct {
// 	service authservice.Service
// }

// func New(service *authservice.AuthService) *AuthController {
// 	return &AuthController{
// 		service: service,
// 	}
// }

// func (c *AuthController) Routes() chi.Router {
// 	r := chi.NewRouter()

// 	r.Post("/login", c.LoginController)
// 	r.Post("/register", c.RegisterController)

// 	return r
// }

// func (c *AuthController) LoginController(w http.ResponseWriter, r *http.Request) {

// 	user, err := jsonutil.DecodeJson[protobuf.CreateUserRequest](r)
// 	if err != nil {
// 		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("Couldn't process the user data"))
// 		return
// 	}
// 	token, err := c.service.LoginUserService(&user)
// 	if err != nil {
// 		status, msg := grpcutil.GRPCToHTTPError(err)
// 		jsonutil.EncodeJson(w, status, msg)
// 		return
// 	}
// 	jsonutil.EncodeJson(w, 200, token)
// }

// func (c *AuthController) RegisterController(w http.ResponseWriter, r *http.Request) {

// }
