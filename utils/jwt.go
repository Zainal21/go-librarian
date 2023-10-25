package utils

import (
	"errors"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("inirahasia")

func GenerateToken(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	if tokenString == "" {
		return "", errors.New("Token is empty or invalid")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			log.Println("error algoritm")
			return nil, jwt.ErrInvalidKey
		}
		return secretKey, nil
	})

	if err != nil {
		log.Println(err.Error())
		log.Println("error parse jwt")
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, found := claims["username"].(string)
		log.Println(claims)
		if !found || username == "" {
			return "", errors.New("Username not found in claims")
		}
		return username, nil
	}
	// log.Println("error invalid key")
	return "", jwt.ErrInvalidKey
}
