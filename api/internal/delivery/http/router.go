package http

import (
	"database/sql"
	"net/http"
	authdelivery "shopito/api/internal/delivery/http/auth"
	userdelivery "shopito/api/internal/delivery/http/user"
	jsonutil "shopito/api/pkg/util/json"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) chi.Router {
	router := chi.NewRouter()

	router.Use(SecureHeaders, Cors, LogRequest)
	router.Route("/api/v1", func(router chi.Router) {
		router.With(Authenticate, AuthenticateAdmin).Get("/ping", func(w http.ResponseWriter, r *http.Request) { jsonutil.EncodeJson(w, 200, "Pong") })
		router.Mount("/auth", authdelivery.New(db).Routes())
		router.Mount("/users", userdelivery.New(db).Routes())
	})
	return router
}
