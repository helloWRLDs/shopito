package admincontroller

import (
	"shopito/services/api-gw/internal/delivery/middleware"

	"github.com/go-chi/chi"
)

func (c *AdminController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Authenticate, middleware.AuthenticateAdmin)

	r.Post("/promote/{id}", c.PromoteUserController)
	r.Post("/demote/{id}", c.DemoteUserController)

	return r
}
