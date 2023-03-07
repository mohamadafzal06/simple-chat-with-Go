package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hash the password: %v\n", err)
	}

	return string(bPass), nil
}

func CheckHashPassword(hashPass string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
}
