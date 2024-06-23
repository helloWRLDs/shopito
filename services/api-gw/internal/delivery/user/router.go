package usercontroller

import (
	"shopito/services/api-gw/internal/delivery/middleware"

	"github.com/go-chi/chi"
)

func (c *UserController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", c.CreateUserController)
	r.With(middleware.Authenticate, middleware.AuthenticateAdmin).
		Get("/", c.ListUsersController)
	r.With(middleware.Authenticate, middleware.AuthenticateSelfOrAdmin).
		Get("/{id}", c.GetUserController)
	r.With(middleware.Authenticate, middleware.AuthenticateSelfOrAdmin).
		Delete("/{id}", c.DeleteUserController)
	r.With(middleware.Authenticate, middleware.AuthenticateSelfOrAdmin).
		Put("/{id}", c.UpdateUserController)

	return r
}
