package admincontroller

import (
	"fmt"
	"net/http"
	grpcutil "shopito/pkg/util/grpc"
	jsonutil "shopito/pkg/util/json"
	adminservice "shopito/services/api-gw/internal/service/admin"

	"github.com/go-chi/chi"
)

type AdminController struct {
	service adminservice.Service
}

func New(service *adminservice.AdminService) *AdminController {
	return &AdminController{
		service: service,
	}
}

func (c *AdminController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/promote/{id}", c.PromoteUserController)
	r.Get("/demote/{id}", c.DemoteUserController)
	r.Post("/", c.NotifyUsersController)

	return r
}

func (c *AdminController) PromoteUserController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := c.service.PromoteUserService(id); err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		jsonutil.EncodeJson(w, status, msg)
		return
	}
	jsonutil.EncodeJson(w, 200, fmt.Sprintf("promoted user with id = %v", id))
}

func (c *AdminController) DemoteUserController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := c.service.DemoteUserService(id); err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		jsonutil.EncodeJson(w, status, msg)
		return
	}
	jsonutil.EncodeJson(w, 200, fmt.Sprintf("demoted user with id = %v", id))
}

func (c *AdminController) NotifyUsersController(w http.ResponseWriter, r *http.Request) {

}
