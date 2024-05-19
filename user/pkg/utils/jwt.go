package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strconv"
	"time"
)

var secretKey = []byte("secret-key")

func CreateToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": strconv.Itoa(int(userID)),
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to get subject")
	}
	log.Println(sub)
	return sub, nil
}
