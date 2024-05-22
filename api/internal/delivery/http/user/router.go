package userdelivery

import (
	"database/sql"
	userusecase "shopito/api/internal/usecase/user"

	"github.com/go-chi/chi"
)

type UserDeliveryImpl struct {
	uc userusecase.UserUseCase
}

func New(db *sql.DB) *UserDeliveryImpl {
	return &UserDeliveryImpl{
		uc: userusecase.New(db),
	}
}

func (d *UserDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", d.GetUsersController)
	r.Get("/{id}", d.GetUserController)
	r.Put("/{id}", d.UpdateUserController)
	r.Delete("/{id}", d.DeleteUserController)

	return r
}
