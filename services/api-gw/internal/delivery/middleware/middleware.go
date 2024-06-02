package middleware

import (
	"net/http"
	"shopito/pkg/types/errors"
	"shopito/services/api-gw/config"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")
		next.ServeHTTP(w, r)
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(
			logrus.Fields{
				"method": r.Method,
				"url":    r.URL.String(),
			},
		).Info("received request")
		next.ServeHTTP(w, r)
	})
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) < 8 {
			errors.SendErr(w, errors.ErrNotAuthorized.SetMessage("missing token"))
			return
		}
		claims := jwt.MapClaims{}
		var (
			key []byte = []byte(config.JWT.SECRET)
			// Slicee to separate Bearer tag
			token string = r.Header.Get("Authorization")[7:]
		)
		if token == "" {
			errors.SendErr(w, errors.ErrNotAuthorized.SetMessage("missing token"))
			return
		}
		_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			errors.SendErr(w, errors.ErrNotAuthorized.SetMessage("couldn't parse claims"))
			return
		}
		if err := claims.Valid(); err != nil {
			errors.SendErr(w, errors.ErrNotAuthorized.SetMessage("invalid token"))
			return
		}
		data, ok := claims["data"].(map[string]any)
		if !ok {
			errors.SendErr(w, errors.ErrNotAuthorized.SetMessage("couldn't process payload"))
			return
		}
		r.Header.Set("id", data["id"].(string))
		r.Header.Set("isAdmin", data["isAdmin"].(string))
		r.Header.Set("isVerified", data["isVerified"].(string))
		next.ServeHTTP(w, r)
	})
}

func AuthenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin, err := strconv.ParseBool(r.Header.Get("isAdmin"))
		if r.Header.Get("isAdmin") == "" || err != nil || !isAdmin {
			errors.SendErr(w, errors.ErrForbidden.SetMessage("no permission"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AuthenticateVerified(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isVerified, err := strconv.ParseBool(r.Header.Get("isVerified"))
		if r.Header.Get("isVerified") == "" || err != nil || !isVerified {
			errors.SendErr(w, errors.ErrForbidden.SetMessage("verified account required"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AuthenticateSelfOrAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin, _ := strconv.ParseBool(r.Header.Get("isAdmin"))
		if chi.URLParam(r, "userId") != r.Header.Get("id") && !isAdmin {
			errors.SendErr(w, errors.ErrForbidden.SetMessage("available for admin and current user only"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
