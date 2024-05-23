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
	r.Get("/{userId}", d.GetUserController)
	r.Put("/{userId}", d.UpdateUserController)
	r.Delete("/{userId}", d.DeleteUserController)

	return r
}
