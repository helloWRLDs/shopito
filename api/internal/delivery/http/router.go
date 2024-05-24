package http

import (
	"database/sql"
	"net/http"
	authdelivery "shopito/api/internal/delivery/http/auth"
	"shopito/api/internal/delivery/http/middleware"
	productsdelivery "shopito/api/internal/delivery/http/products"
	userdelivery "shopito/api/internal/delivery/http/user"
	jsonutil "shopito/api/pkg/util/json"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func Router(db *sql.DB) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.SecureHeaders)
	router.Use(middleware.LogRequest)
	// router.Use(cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
	// 	AllowCredentials: true,
	// }).Handler)
	router.Use(cors.Default().Handler)
	router.Route("/api/v1", func(router chi.Router) {
		router.With(middleware.Authenticate, middleware.AuthenticateAdmin).
			Get("/ping", func(w http.ResponseWriter, r *http.Request) { jsonutil.EncodeJson(w, 200, "Pong") })
		router.Mount("/auth", authdelivery.New(db).Routes())
		router.Mount("/users", userdelivery.New(db).Routes())
		router.Mount("/products", productsdelivery.New(db).Routes())
	})
	return router
}
