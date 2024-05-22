package http

import (
	"fmt"
	"net/http"
	"shopito/api/config"
	"shopito/api/pkg/types/errors"
	jsonutil "shopito/api/pkg/util/json"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
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
		claims := jwt.MapClaims{}
		var (
			key []byte = []byte(config.JWT.SECRET)
			// Slicee to separate Bearer tag
			token string = r.Header.Get("Authorization")[7:]
		)
		if token == "" {
			jsonutil.EncodeJson(w, errors.ErrNotAuthorized.Status(), errors.ErrNotAuthorized.SetMessage("missing token"))
			return
		}
		_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			jsonutil.EncodeJson(w, errors.ErrNotAuthorized.Status(), errors.ErrNotAuthorized.SetMessage("error parsing claims"))
			return
		}
		if err := claims.Valid(); err != nil {
			jsonutil.EncodeJson(w, errors.ErrNotAuthorized.Status(), errors.ErrNotAuthorized.SetMessage("token is not valid"))
			return
		}
		data, ok := claims["data"].(map[string]any)
		if !ok {
			jsonutil.EncodeJson(w, errors.ErrNotAuthorized.Status(), errors.ErrNotAuthorized.SetMessage("error processing payload"))
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
			jsonutil.EncodeJson(w, errors.ErrForbidden.Status(), errors.ErrForbidden.SetMessage("not enough priveleges"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func VerifiedsOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isVerified, err := strconv.ParseBool(r.Header.Get("isVerified"))
		if r.Header.Get("isVerified") == "" || err != nil || !isVerified {
			jsonutil.EncodeJson(w, errors.ErrForbidden.Status(), errors.ErrForbidden.SetMessage("verified account required"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ExposeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Println("id", r.Header.Get("id"))
		fmt.Println("isAdmin", r.Header.Get("isAdmin"))
		fmt.Println("isVerified", r.Header.Get("isVerified"))
	})
}
