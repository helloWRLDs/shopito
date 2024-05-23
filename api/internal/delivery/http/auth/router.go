package authdelivery

import (
	"github.com/go-chi/chi"
)

func (d *AuthDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/register", d.RegisterController)
	r.Post("/login", d.LoginController)
	r.Post("/{userId}/verify/{token}", d.VerifyController)
	r.Post("/{userId}/resend", d.ResendController)
	return r
}
