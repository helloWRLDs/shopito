package authcontroller

import "github.com/go-chi/chi"

func (c *AuthController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", c.LoginController)
	r.Post("/register", c.RegisterController)

	return r
}
