package userdelivery

import (
	"shopito/api/internal/delivery/http/middleware"

	"github.com/go-chi/chi"
)

func (d *UserDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()

	r.With(middleware.Authenticate, middleware.AuthenticateAdmin).
		Get("/", d.GetUsersController)
	r.With(middleware.Authenticate, middleware.AuthenticateSelfOrAdmin).
		Get("/{userId}", d.GetUserController)
	r.With(middleware.Authenticate, middleware.AuthenticateSelfOrAdmin).
		Put("/{userId}", d.UpdateUserController)
	r.With(middleware.Authenticate, middleware.AuthenticateSelfOrAdmin).
		Delete("/{userId}", d.DeleteUserController)

	return r
}
