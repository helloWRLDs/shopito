package usercontroller

import "github.com/go-chi/chi"

func (c *UserController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", c.GetUserController)
	r.Get("/", c.ListUsersController)
	r.Post("/", c.CreateUserController)
	r.Delete("/{id}", c.DeleteUserController)
	r.Put("/{id}", c.UpdateUserController)

	return r
}
