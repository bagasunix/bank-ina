package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/bagasunix/bank-ina/pkg/envs"
	"github.com/bagasunix/bank-ina/pkg/errors"
)

func GenerateToken(jwtKey string, claims Claims) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return t.SignedString([]byte(jwtKey))
}

func ValidateToken(signedToken string) (claims *Claims, err error) {
	conf, _ := envs.LoadEnv()
	t, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err = errors.CustomError(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
			return nil, err
		}
		return []byte(conf.JWTKey), nil
	})

	if err != nil {
		return nil, errors.CustomError(err.Error())
	}

	claims, ok := t.Claims.(*Claims)
	if !ok {
		err = errors.CustomError("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		err = errors.CustomError("token expired")
		return nil, err
	}

	return claims, err
}
