package utils

import (
	"time"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/dgrijalva/jwt-go"
)

type GSPClaim struct {
	User    models.ProfiledUser `json:"user"`
	IsAdmin bool                `json:"is_admin"`
	jwt.StandardClaims
}

var JWTSecret = []byte("chester")

func GenerateJWT(u models.ProfiledUser) (string, error) {
	claims := &GSPClaim{
		User:    u,
		IsAdmin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(JWTSecret)
	if err != nil {
		return t, err
	}
	return t, nil
}
