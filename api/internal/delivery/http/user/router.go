package userdelivery

import (
	"database/sql"
	"shopito/api/internal/delivery/http/middleware"
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
