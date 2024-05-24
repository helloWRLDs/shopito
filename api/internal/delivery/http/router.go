package http

import (
	"database/sql"
	"net/http"
	admindelivery "shopito/api/internal/delivery/http/admin"
	authdelivery "shopito/api/internal/delivery/http/auth"
	"shopito/api/internal/delivery/http/middleware"
	productsdelivery "shopito/api/internal/delivery/http/products"
	userdelivery "shopito/api/internal/delivery/http/user"
	jsonutil "shopito/api/pkg/util/json"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.SecureHeaders)
	router.Use(middleware.LogRequest)
	router.Use(middleware.Cors)
	router.Route("/api/v1", func(router chi.Router) {
		router.With(middleware.Authenticate, middleware.AuthenticateAdmin).
			Get("/ping", func(w http.ResponseWriter, r *http.Request) { jsonutil.EncodeJson(w, 200, "Pong") })
		router.Mount("/auth", authdelivery.New(db).Routes())
		router.Mount("/users", userdelivery.New(db).Routes())
		router.Mount("/products", productsdelivery.New(db).Routes())
		router.Mount("/admin", admindelivery.New(db).Routes())
	})
	return router
}
