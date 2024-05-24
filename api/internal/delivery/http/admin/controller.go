package admindelivery

import (
	"database/sql"
	"fmt"
	"net/http"
	adminusecase "shopito/api/internal/usecase/admin"
	"shopito/api/pkg/types/errors"
	"shopito/api/pkg/types/response"
	jsonutil "shopito/api/pkg/util/json"
	"strconv"

	"github.com/go-chi/chi"
)

type AdminDeliveryImpl struct {
	uc adminusecase.AdminUseCase
}

func New(db *sql.DB) *AdminDeliveryImpl {
	return &AdminDeliveryImpl{
		uc: adminusecase.New(db),
	}
}

func (d *AdminDeliveryImpl) PromoteUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process user's id"))
		return
	}
	if err := d.uc.PromoteUser(r.Context(), id); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("promoted user with id=%v", id)))
}

func (d *AdminDeliveryImpl) NotifyUsersController(w http.ResponseWriter, r *http.Request) {
	message, err := jsonutil.DecodeJson[response.EmailMessage](r)
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process message"))
		return
	}
	if err := d.uc.NotifyUsers(r.Context(), message); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, "notification sent to users emails"))
}

func (d *AdminDeliveryImpl) DemoteUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process user's id"))
		return
	}
	if err := d.uc.DemoteUser(r.Context(), id); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("demoted user with id=%v", id)))
}
