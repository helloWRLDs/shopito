package jwtutil

import (
	"shopito/api/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id int, isAdmin, isVerified bool) (*string, error) {
	var secret []byte = []byte(config.JWT.SECRET)
	claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]string{
			"id":         strconv.FormatInt(int64(id), 10),
			"isAdmin":    strconv.FormatBool(isAdmin),
			"isVerified": strconv.FormatBool(isVerified),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}
