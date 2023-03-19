package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultCost int = 10
)

func HashPassword(password string) (string, error) {
	//TODO: does not accept password more than 72 bytes, check length
	//TODO: add salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
