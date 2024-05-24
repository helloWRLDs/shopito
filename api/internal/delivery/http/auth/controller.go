package authdelivery

import (
	"database/sql"
	"fmt"
	"net/http"
	userdomain "shopito/api/internal/domain/user"
	authusecase "shopito/api/internal/usecase/auth"
	"shopito/api/pkg/types/errors"
	"shopito/api/pkg/types/response"
	jsonutil "shopito/api/pkg/util/json"
	"strconv"

	"github.com/go-chi/chi"
)

type (
	AuthDeliveryImpl struct {
		uc authusecase.AuthUseCase
	}
)

func New(db *sql.DB) *AuthDeliveryImpl {
	return &AuthDeliveryImpl{
		uc: authusecase.New(db),
	}
}

func (d *AuthDeliveryImpl) RegisterController(w http.ResponseWriter, r *http.Request) {
	user, err := jsonutil.DecodeJson[userdomain.User](r)
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("error parsing json"))
		return
	}
	id, error := d.uc.RegisterUser(r.Context(), &user)
	if error != nil {
		jsonutil.EncodeJson(w, error.Status(), error)
		return
	}
	jsonutil.EncodeJson(w, 201, response.NewJsonMessage(201, fmt.Sprintf("user registered with id=%v", id)))
}

func (d *AuthDeliveryImpl) LoginController(w http.ResponseWriter, r *http.Request) {
	user, err := jsonutil.DecodeJson[userdomain.User](r)
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrUnpocessableEntity.Status(), errors.ErrUnpocessableEntity.SetMessage("error parsing json"))
		return
	}
	fmt.Println(user.Password)
	token, error := d.uc.LoginUser(r.Context(), &user)
	if error != nil {
		jsonutil.EncodeJson(w, error.Status(), error)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, *token))
}

func (d *AuthDeliveryImpl) VerifyController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrBadRequest.Status(), errors.ErrBadRequest.SetMessage("wrong input"))
		return
	}
	token := chi.URLParam(r, "token")
	if err := d.uc.VerifyUser(r.Context(), id, token); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, fmt.Sprintf("verified user with id=%v", id)))
}

func (d *AuthDeliveryImpl) ResendController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		jsonutil.EncodeJson(w, errors.ErrBadRequest.Status(), errors.ErrBadRequest.SetMessage("wrong input"))
		return
	}
	if err := d.uc.RetryVerification(r.Context(), id); err != nil {
		jsonutil.EncodeJson(w, err.Status(), err)
		return
	}
	jsonutil.EncodeJson(w, 200, response.NewJsonMessage(200, "verification code was resend to email"))
}
