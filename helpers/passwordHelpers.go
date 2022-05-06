package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	valid := true
	msg := ""

	if err != nil {
		msg = "email or password is incorrect"
		valid = false
		fmt.Println(err)
	}

	return valid, msg
}