package admindelivery

import (
	"shopito/api/internal/delivery/http/middleware"

	"github.com/go-chi/chi"
)

func (d *AdminDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Authenticate, middleware.AuthenticateAdmin)

	r.Post("/promote/{id}", d.PromoteUserController)
	r.Post("/demote/{id}", d.DemoteUserController)
	r.Post("/notify", d.NotifyUsersController)

	return r
}
