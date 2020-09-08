package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaims struct {
	Id    int64  `json:"userId"`
	Name  string `json:"userName"`
	Token string `json:"token"`
	jwt.StandardClaims
}

func CreateJWTToken(id int64, username string) (string, error) {
	key, err := GetEnvVar("token_password")
	if err != nil {
		return "", err
	}

	var jwtKey = []byte(key)

	expirationTime := time.Now().Add(time.Hour * 7 * 24).Unix()
	claims := &TokenClaims{
		Id:   id,
		Name: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString(jwtKey)
	return tokenString, err
}

func VerifyJWTToken(tokenString string) (*TokenClaims, error) {
	claims := &TokenClaims{}
	jwtKey, err := GetEnvVar("token_password")
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UnExpected signing method")
		}

		return []byte(jwtKey), nil
	})

	if !token.Valid {
		return nil, err
	}

	return claims, err
}
