package utils

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load .env file %v", err)
	}
}

// var SecretKey = "secretkey"
var SecretKey = os.Getenv("JWT_SECRET_KEY")

// create new token using user ID
func GenerateJwt(userId uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	return claims.SignedString([]byte(SecretKey))
}

// parse cookie to get token then user ID
func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err != nil || !token.Valid {
		return "", err
	}

	cliams := token.Claims.(*jwt.StandardClaims)

	return cliams.Issuer, nil
}
