package utils

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Username   string
	First_Name string
	Last_Name  string
	jwt.StandardClaims
}

func CheckAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Authorization") == "" {
		return false
	}
	return true
}

func GetTokens(firstname, lastname, username string) (string, string, error) {
	var secret = os.Getenv("SECRET")
	claims := &SignedDetails{
		Username:   username,
		First_Name: firstname,
		Last_Name:  lastname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		Username:   username,
		First_Name: firstname,
		Last_Name:  lastname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	return token, refreshToken, err
}

func ValidateToken(tokenString string) bool {
	var secret = os.Getenv("SECRET")
	claims := &SignedDetails{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	return true
}
