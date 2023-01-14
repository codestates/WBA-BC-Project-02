package user

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(hashedPassword, userPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword)); err != nil {
		return err
	}
	return nil
}
