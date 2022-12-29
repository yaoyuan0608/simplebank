package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// return the bcrypt password
func HashPassword(password string) (string, error) {
	hahsedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hahsedPassword), nil
}

// check password is correct
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
