package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(email string) (string, error) {
	err := godotenv.Load(".env")

	secretKey := os.Getenv("SECRET_KEY")

	if err != nil {
		fmt.Println("Cannot get env variables")
		return "", err
	}

	var jwtKey = []byte(secretKey)

	expirationTime := time.Now().Add(5 * time.Hour)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}