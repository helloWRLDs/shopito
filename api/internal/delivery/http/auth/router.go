package authdelivery

import (
	"github.com/go-chi/chi"
)

func (d *AuthDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/register", d.RegisterController)
	r.Post("/login", d.LoginController)
	r.Post("/{id}/verify/{token}", d.VerifyController)
	r.Post("/{id}/resend", d.ResendController)
	return r
}
