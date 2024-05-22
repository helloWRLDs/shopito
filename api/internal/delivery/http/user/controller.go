package userdelivery

import (
	"fmt"
	"net/http"
	userdomain "shopito/api/internal/domain/user"
	"shopito/api/pkg/types/errors"
	"shopito/api/pkg/types/response"
	jsonutil "shopito/api/pkg/util/json"
	"strconv"

	"github.com/go-chi/chi"
)

func (d *UserDeliveryImpl) GetUsersController(w http.ResponseWriter, r *http.Request) {
	users, err := d.uc.GetUsers(r.Context())
	if err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, users)
}

func (d *UserDeliveryImpl) GetUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process user's id"))
		return
	}
	user, error := d.uc.GetUser(r.Context(), id)
	if error != nil {
		jsonutil.EncodeJson(w, error.Status(), error)
		return
	}
	jsonutil.EncodeJson(w, 200, user)
}

func (d *UserDeliveryImpl) UpdateUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process user's id"))
		return
	}
	newUser, err := jsonutil.DecodeJson[userdomain.User](r)
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process user's data"))
		return
	}
	if err := d.uc.UpdateUser(r.Context(), id, &newUser); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("updated user with id=%v", id)))
}

func (d *UserDeliveryImpl) DeleteUserController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("couldn't process user's id"))
		return
	}
	if err := d.uc.DeleteUser(r.Context(), id); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("deleted user with id=%v", id)))
}
