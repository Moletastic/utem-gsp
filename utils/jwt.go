package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/dgrijalva/jwt-go"
)

type GSPClaim struct {
	User    models.User `json:"user"`
	IsAdmin bool        `json:"is_admin"`
	jwt.StandardClaims
}

func (c *GSPClaim) Bind(m jwt.Claims) {
	v, ok := m.(GSPClaim)
	if ok {
		fmt.Println("yei...")
		c = &v
	} else {
		var claim GSPClaim
		bytes, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = json.Unmarshal(bytes, &claim)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(Pretty(m))
		fmt.Println(Pretty(claim))
	}
}

var JWTSecret = []byte("chester")

func GenerateJWT(u models.User) (string, error) {
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
