package admincontroller

import (
	"context"
	"fmt"
	"net/http"
	"shopito/pkg/types/errors"
	"shopito/pkg/types/response"
	grpcutil "shopito/pkg/util/grpc"
	jsonutil "shopito/pkg/util/json"
	adminservice "shopito/services/api-gw/internal/service/admin"
	"strconv"
	"time"

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

// @Summary Promote User
// @Tags Admin
// @Description Promote User By ID
// @Accept json
// @Produce json
// Param id path int true "User ID"
// @Success     200 {object} response.JsonMessage "OK"
// @Failure     422 {object} errors.HTTPError "Unprocessable entity"
// @Failure 	404 {object} errors.HTTPError "Not Found"
// @Failure     500 {object} errors.HTTPError "Internal server error"
// @Router /admin/promote/{id} [POST]
func (c *AdminController) PromoteUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("couldn't process the id"))
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	err = c.service.PromoteUserService(ctx, id)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		errors.SendErr(w, errors.GetHTTPErrorByCode(status, msg))
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("promoted user with id = %v", id)))
}

// @Summary Demote User
// @Tags Admin
// @Description Demote User By ID
// @Accept json
// @Produce json
// Param id path int true "User ID"
// @Success     200 {object} response.JsonMessage "OK"
// @Failure     422 {object} errors.HTTPError "Unprocessable entity"
// @Failure 	404 {object} errors.HTTPError "Not Found"
// @Failure     500 {object} errors.HTTPError "Internal server error"
// @Router /admin/demote/{id} [POST]
func (c *AdminController) DemoteUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("couldn't process the id"))
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	err = c.service.DemoteUserService(ctx, id)
	if err != nil {
		status, msg := grpcutil.GRPCToHTTPError(err)
		errors.SendErr(w, errors.GetHTTPErrorByCode(status, msg))
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("demoted user with id = %v", id)))
}
