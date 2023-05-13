package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultCost int = 10
)

func HashPassword(password string) (string, error) {
	if len(password) > 36 {
		ErrPasswordTooLong := errors.New("password length should not exceed 72 bytes")
		return "", ErrPasswordTooLong
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
