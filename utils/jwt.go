package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	Hash   string `json:"hash"`
	Expire int64  `json:"expire"`
}

func CreateToken(userID int) (*Token, error) {
	t := new(Token)
	secret := Config("SECRET_JWT")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	expiresIn := time.Now().Add(time.Duration(60*60*24) * time.Second).Unix()
	tokenHash, err := token.SignedString([]byte(secret))
	if err != nil {
		return t, err
	}
	t.Hash = tokenHash
	t.Expire = expiresIn
	return t, nil
}

func ParseToken(tokenString string) (float64, error) {
	secret := Config("SECRET_JWT")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	err2 := claims.Valid()
	if err2 != nil {
		return 0, err2
	}

	return claims["id"].(float64), nil
}
