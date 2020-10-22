package utils

import (
	"fmt"
	"crypto/rand"
	"github.com/dgrijalva/jwt-go"
	"github.com/iamqiutian/uploadFile/g"
)

func CreateRandom() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

func CheckAuth(auth string) bool {
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		return g.MySigningKey, nil
	})

	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	return true
}